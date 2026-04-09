// Copyright 2024 Harald Albrecht.
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

	"github.com/thediveo/nonstd/xiter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("swapping sequence (k,v) pairs", func() {

	It("swaps each pair", func() {
		m := maps.Collect(xiter.Swap(slices.All([]string{"foo", "bar"})))
		Expect(m).To(HaveLen(2))
		Expect(m).To(HaveKeyWithValue("foo", 0))
		Expect(m).To(HaveKeyWithValue("bar", 1))
	})

	It("aborts", func() {
		count := 0
		for range xiter.Swap(slices.All([]string{"foo", "bar"})) {
			count++
			break
		}
		Expect(count).To(Equal(1))
	})

})
