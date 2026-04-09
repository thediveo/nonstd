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

package xmaps_test

import (
	"slices"

	"github.com/thediveo/nonstd/xiter"
	"github.com/thediveo/nonstd/xmaps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("collecting sequences into maps", func() {

	It("aggregates keys with multiple values", func() {
		seq := xiter.Swap(slices.All([]string{
			"foo",
			"bar",
			"foo",
		}))
		m := xmaps.Collect(seq)
		Expect(m).To(HaveLen(2))
		Expect(m).To(HaveKeyWithValue("foo", []int{0, 2}))
		Expect(m).To(HaveKeyWithValue("bar", []int{1}))
	})

	It("aggregates keys with multiple values into a value slice of type T", func() {
		type Indices []int
		seq := xiter.Swap(slices.All([]string{
			"foo",
			"bar",
			"foo",
		}))
		m := xmaps.CollectT[Indices](seq)
		Expect(m).To(HaveLen(2))
		Expect(m).To(HaveKeyWithValue("foo", Indices{0, 2}))
		Expect(m).To(HaveKeyWithValue("bar", Indices{1}))
	})

})
