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

package xslices

import (
	"slices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("all slice values", func() {

	It("correctly iterates over all slice values", func() {
		s := []string{"foo", "bar", "baz"}
		Expect(slices.Collect(AllValues(s))).To(Equal(s))
	})

	It("aborts iterating over all slice values", func() {
		num := 0
		for range AllValues([]string{"foo", "bar", "baz"}) {
			num++
			break
		}
		Expect(num).To(Equal(1))
	})

	It("handles nil slices", func() {
		Expect(slices.Collect(AllValues[[]string](nil))).To(BeEmpty())
	})

})
