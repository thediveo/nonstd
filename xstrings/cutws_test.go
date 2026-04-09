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

package xstrings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thediveo/nonstd/xstrings"
)

var _ = Describe("cutting at white spaces", func() {

	DescribeTable("CutWhitespace",
		func(s, before, after string, found bool) {
			b, a, f := xstrings.CutWhitespace(s)
			Expect(b).To(Equal(before))
			Expect(a).To(Equal(after))
			Expect(f).To(Equal(found))
		},
		Entry(nil, "", "", "", false),
		Entry(nil, "\t", "", "", true),
		Entry(nil, "abc-def", "abc-def", "", false),
		Entry(nil, "abc def", "abc", "def", true),
		Entry(nil, "abc  def", "abc", " def", true),
		Entry(nil, "abc\tdef", "abc", "def", true),
		Entry(nil, "abc\t\tdef", "abc", "\tdef", true),
		Entry(nil, "\tabc", "", "abc", true),
	)

})
