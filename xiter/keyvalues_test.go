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
	"maps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("iterating over keys or values only from key-value sequences", func() {

	var m = map[string]int{
		"foo": 42,
		"bar": 666,
	}

	Context("only keys", func() {

		It("yields all keys", func() {
			Expect(Keys(maps.All(m))).To(ConsistOf("foo", "bar"))
		})

		It("aborts the key sequence", func() {
			count := 0
			for range Keys(maps.All(m)) {
				count++
				break
			}
			Expect(count).To(Equal(1))
		})

	})

	Context("only values", func() {

		It("yields all values", func() {
			Expect(Values(maps.All(m))).To(ConsistOf(42, 666))
		})

		It("aborts the value sequence", func() {
			count := 0
			for range Values(maps.All(m)) {
				count++
				break
			}
			Expect(count).To(Equal(1))
		})

	})

})
