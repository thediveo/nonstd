// Copyright 2023 Harald Albrecht.
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

var _ = Describe("reversing slices", func() {

	It("reverses slices", func() {
		Expect(ReverseCopy([]rune{'H', 'E', 'L', 'O'})).To(Equal([]rune{'O', 'L', 'E', 'H'}))
		Expect(ReverseCopy([]rune{'H', 'E', 'L', 'l', 'O'})).To(Equal([]rune{'O', 'l', 'L', 'E', 'H'}))
	})

	It("works on a copy", func() {
		s := []rune{'F', 'U', 'B', 'A', 'R'}
		sorig := slices.Clone(s)
		Expect(ReverseCopy(s)).To(Equal([]rune{'R', 'A', 'B', 'U', 'F'}))
		Expect(s).To(Equal(sorig))
	})

})
