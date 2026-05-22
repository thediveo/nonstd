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

package xslices_test

import (
	"github.com/thediveo/nonstd/xslices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("returning the first matching element from a slice", func() {

	DescribeTable("returning a match or zero",
		func(s []string, expected string) {
			Expect(xslices.FindFunc(s, func(e string) bool { return e == "bar" })).To(Equal(expected))
		},
		Entry(nil, nil, ""),
		Entry(nil, []string{"foo", "baz"}, ""),
		Entry(nil, []string{"foo", "bar", "baz"}, "bar"),
	)

	DescribeTable("returning a match with a match indication",
		func(s []string, expected string, match bool) {
			v, ok := xslices.FindFuncOk(s, func(e string) bool { return e == "bar" })
			Expect(ok).To(Equal(match))
			Expect(v).To(Equal(expected))
		},
		Entry(nil, nil, "", false),
		Entry(nil, []string{"foo", "baz"}, "", false),
		Entry(nil, []string{"foo", "bar", "baz"}, "bar", true),
	)

})
