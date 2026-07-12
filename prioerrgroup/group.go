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

package prioerrgroup

import (
	"cmp"
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Group is a collection of go routines working on sub tasks that are part of
// the same overall task. In case one or more sub tasks return with
// (prioritized) errors, only the highest priority error is returned.
//
// Use [WithContext] to create a Group, Group zero values are not usable.
//
// A Group has no limit on the number of active go routines.
type Group[P PrioritizedError[O], O cmp.Ordered] struct {
	m      sync.Mutex
	errgrp *errgroup.Group
	prerr  *PrioritizedError[O]
}

// PrioritizedError encapsulates an error value together with a priority.
// Priorities can be of any cmp.Ordered type. The higher a priority value, the
// higher its priority.
type PrioritizedError[O cmp.Ordered] struct {
	err  error
	prio O
}

// NewPrioritizedError returns a new prioritized error wrapping the specified
// error value with the given priority.
func NewPrioritizedError[O cmp.Ordered](err error, prio O) *PrioritizedError[O] {
	return &PrioritizedError[O]{err: err, prio: prio}
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to
// [Group.Go] returns a non-nil prioritized error or the first time [Group.Wait]
// returns, whichever occurs first.
func WithContext[O cmp.Ordered](ctx context.Context) (*Group[PrioritizedError[O], O], context.Context) {
	errgrp, ctx := errgroup.WithContext(ctx)
	return &Group[PrioritizedError[O], O]{
		errgrp: errgrp,
	}, ctx
}

// Go calls the given function in a new goroutine, expecting it to either return
// a nil priority error or otherwise a prioritized error. Use
// [NewPrioritizedError] to create a new prioritized error to return.
//
// As with [errgroup.Group.Go], the first call to Go must happen before a
// [Group.Wait].
//
// The first go routine in this group that returns a non-nil prioritized error
// will cancel the associated context.
func (g *Group[P, O]) Go(fn func() *PrioritizedError[O]) {
	g.m.Lock()
	errgrp := g.errgrp
	g.m.Unlock()
	errgrp.Go(func() error {
		err := fn()
		if err == nil {
			return nil
		}
		g.m.Lock()
		defer g.m.Unlock()
		if g.prerr == nil || err.prio > g.prerr.prio {
			g.prerr = err
		}
		return err.err
	})
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the highest (numerical) priority non-nil error (if any) from them.
func (g *Group[P, O]) Wait() error {
	// Due to racy sub tasks returning errors we ignore the underlying error
	// group's Wait error return value. Instead, we consult our highest priority
	// error value, if any.
	g.m.Lock()
	errgrp := g.errgrp
	g.m.Unlock()
	_ = errgrp.Wait()
	g.m.Lock()
	defer g.m.Unlock()
	if g.prerr == nil {
		return nil
	}
	return g.prerr.err
}
