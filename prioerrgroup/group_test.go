// Copyright 2026 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package prioerrgroup_test

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/thediveo/testily/chans"
	"github.com/thediveo/testily/goroutines"

	"github.com/thediveo/nonstd/prioerrgroup"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thediveo/testily/chans"
)

type ErrorPriorities int16

const (
	fooTaskPrio ErrorPriorities = iota + 1
	barTaskPrio
)

var _ = Describe("prioritized error groups", func() {

	DescribeTable("keeping error priorities",
		func(prioerrs []*prioerrgroup.PrioritizedError[ErrorPriorities], expected string) {
			g, _ := prioerrgroup.WithContext[ErrorPriorities](context.Background())

			newG := func(
				fn func() *prioerrgroup.PrioritizedError[ErrorPriorities],
			) (func(), goroutines.Goroutine) {
				unblockch, done := chans.Make[Nothing]()
				gch := make(chan goroutines.Goroutine)
				g.Go(func() *prioerrgroup.PrioritizedError[ErrorPriorities] {
					closegch := sync.OnceFunc(func() { close(gch) })
					defer closegch()
					gch <- goroutines.Current()
					closegch()
					<-unblockch
					return fn()
				})
				gort := <-gch
				Expect(gort).NotTo(BeZero(), "could not determine go routine details")
				return done, gort
			}

			// Kick off a bunch of sub task go routines that eventually return
			// the passed error (or nil), but at first block so we can unblock
			// them in a controlled manner later, in the same sequence as
			// created.
			var unblocks []func()
			var gorts []goroutines.Goroutine
			for _, prioerr := range prioerrs {
				unblock, gort := newG(func() *prioerrgroup.PrioritizedError[ErrorPriorities] {
					return prioerr
				})
				defer unblock()
				unblocks = append(unblocks, unblock)
				gorts = append(gorts, gort)
			}

			// Now it's time to start the go routine that waits for the sub
			// tasks to complete or fail and then report back to us. We run this
			// on a separate go routine as we want to run the assertions on the
			// current go routine.
			errch, errcloser := Make[error](1)
			go func() {
				defer GinkgoRecover()
				defer errcloser()
				errch <- g.Wait()
			}()

			// Wait for all sub task go routines to have entered blocking state,
			// waiting for us to unblock them next.
			for _, gort := range gorts {
				Eventually(goroutines.ByID).WithArguments(gort.ID).
					Within(2*time.Second).ProbeEvery(10*time.Millisecond).
					Should(HaveField("State", goroutines.WaitChanReceive),
						"go routine didn't wait for <-chan")
			}

			// Unblock one sub task after another, making suring after each
			// unblock that the affected go routine in fact has terminated
			// before we proceed further.
			for idx := range gorts {
				unblocks[idx]()
				Eventually(goroutines.ByID).WithArguments(gorts[idx].ID).
					Within(2*time.Second).ProbeEvery(10*time.Millisecond).
					Should(HaveField("ID", BeZero()),
						"go routine won't terminate")
			}

			// Pick up any error or success reported by the prioritized error
			// group.
			var err error
			Eventually(errch).Within(2 * time.Second).ProbeEvery(10 * time.Millisecond).
				Should(Receive(&err))
			if expected == "" {
				Expect(err).NotTo(HaveOccurred())
			} else {
				Expect(err).To(MatchError(expected))
			}
		},
		Entry("higher prio last",
			[]*prioerrgroup.PrioritizedError[ErrorPriorities]{
				prioerrgroup.NewPrioritizedError(errors.New("foo error"), fooTaskPrio),
				prioerrgroup.NewPrioritizedError(errors.New("bar error"), barTaskPrio),
			},
			"bar error"),
		Entry("higher prio first",
			[]*prioerrgroup.PrioritizedError[ErrorPriorities]{
				prioerrgroup.NewPrioritizedError(errors.New("bar error"), barTaskPrio),
				prioerrgroup.NewPrioritizedError(errors.New("foo error"), fooTaskPrio),
			},
			"bar error"),
		Entry("best error last",
			[]*prioerrgroup.PrioritizedError[ErrorPriorities]{
				nil,
				prioerrgroup.NewPrioritizedError(errors.New("foo error"), fooTaskPrio),
			},
			"foo error"),
	)

})
