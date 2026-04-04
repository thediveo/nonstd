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

package sets

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sets", func() {

	It("returns true or false, depending on an element in the set or not", func() {
		s := New(42)
		Expect(s.Contains(666)).To(BeFalse())
		Expect(s.Contains(42)).To(BeTrue())
	})

	It("returns the correct textual representation", func() {
		s := New(42)
		s.Add(666)
		Expect(s.String()).To(SatisfyAny(
			Equal("[42 666]"),
			Equal("[666 42]")))

		s.Delete(666)
		Expect(s.Elements()).To(ConsistOf(42))
	})

})
