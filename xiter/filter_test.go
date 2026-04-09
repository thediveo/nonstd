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

package xiter_test

import (
	"maps"
	"slices"
	"strings"

	"github.com/thediveo/nonstd/xiter"
	"github.com/thediveo/nonstd/xslices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("filtering sequences", func() {

	Context("value sequences", func() {

		It("filters values", func() {
			s := slices.Collect(
				xiter.Filter(xslices.AllValues([]string{"micro", "microslop", "slopslop", "foo"}),
					func(v string) bool { return !strings.Contains(v, "slop") }))
			Expect(s).To(Equal([]string{"micro", "foo"}))
		})

		It("aborts", func() {
			count := 0
			for range xiter.Filter(xslices.AllValues([]string{"micro", "microslop", "slopslop", "foo"}),
				func(string) bool { return true }) {
				count++
				break
			}
			Expect(count).To(Equal(1))
		})

	})

	Context("key-value sequences", func() {

		It("filters key-values", func() {
			m := maps.Collect(
				xiter.Filter2(slices.All([]string{"micro", "microslop", "slopslop", "foo"}),
					func(k int, v string) bool {
						return k != 0 && !strings.Contains(v, "slop")
					}))
			Expect(m).To(Equal(map[int]string{3: "foo"}))
		})

		It("aborts", func() {
			count := 0
			for range xiter.Filter2(slices.All([]string{"micro", "microslop", "slopslop", "foo"}),
				func(int, string) bool { return true }) {
				count++
				break
			}
			Expect(count).To(Equal(1))
		})

	})

})
