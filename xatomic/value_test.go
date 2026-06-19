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

package xatomic_test

import (
	"github.com/thediveo/nonstd/xatomic"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("spacing concatenated strings", func() {

	It("initially loads a zero value", func() {
		var v xatomic.Value[int]
		Expect(v.Load()).To(BeZero())
	})

	It("stores, loads, and swaps", func() {
		var v xatomic.Value[string]
		v.Store("foo")
		Expect(v.Load()).To(Equal("foo"))
		v.Store("bar")
		Expect(v.Load()).To(Equal("bar"))
		Expect(v.Swap("fool")).To(Equal("bar"))
	})

})
