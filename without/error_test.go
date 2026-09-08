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

package without

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("without errors", func() {

	It("returns the single result value", func() {
		x1 := 42
		r1 := Error(x1, errors.New("Oh Globbits!"))
		Expect(r1).To(Equal(x1))
	})

	It("returns the pair of result values", func() {
		x1 := 42
		x2 := "foo!"
		r1, r2 := Error2(x1, x2, errors.New("Oh Globbits!"))
		Expect(r1).To(Equal(x1))
		Expect(r2).To(Equal(x2))
	})

	It("returns the triple of result values", func() {
		x1 := 42
		x2 := "foo!"
		x3 := 6.66
		r1, r2, r3 := Error3(x1, x2, x3, errors.New("Oh Globbits!"))
		Expect(r1).To(Equal(x1))
		Expect(r2).To(Equal(x2))
		Expect(r3).To(Equal(x3))
	})

})
