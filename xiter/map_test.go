// Copyright 2026 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xiter

import (
	"slices"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("mapping sequences", func() {

	It("transforms a sequence", func() {
		wisdoms := []string{"foo", "bar"}
		Expect(slices.Collect(Map(slices.Values(wisdoms), strings.ToUpper))).To(
			Equal([]string{"FOO", "BAR"}))
	})

	It("correctly aborts a sequence", func() {
		var nexti int
		for i := range Map(func(yield func(int) bool) {
			for {
				if !yield(nexti) {
					return
				}
				nexti++
			}
		},
			func(e int) int { return e + 1 }) {
			Expect(i).To(Equal(1))
			break
		}
		Expect(nexti).To(Equal(0))
	})

})
