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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("collecting key-values", func() {

	It("collects into separate slices", func() {
		seq := func(yield func(int, string) bool) {
			for k, v := range map[int]string{
				42:  "42",
				666: "666",
			} {
				if !yield(k, v) {
					return
				}
			}
		}
		ks, vs := Collect2(seq)
		Expect(ks).To(Equal([]int{42, 666}))
		Expect(vs).To(Equal([]string{"42", "666"}))
	})

})
