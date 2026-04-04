// Copyright 2023, 2025 Harald Albrecht.
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
	"strings"

	"slices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sorting", func() {

	DescribeTable("sorting a copy",
		func(slice []string, expected []string) {
			original := slices.Clone(slice)
			actual := SortedCopy(slice)
			Expect(actual).To(ConsistOf(expected), "incorrectly sorted")
			Expect(slice).To(ConsistOf(original), "modified original")
		},
		Entry(nil, []string{"bar", "foo"}, []string{"bar", "foo"}),
		Entry(nil, []string{"foo", "bar"}, []string{"bar", "foo"}),
	)

	DescribeTable("sorting a copy using a compare function",
		func(slice []string, expected []string) {
			original := slices.Clone(slice)
			actual := SortedCopyFunc(slice, strings.Compare)
			Expect(actual).To(ConsistOf(expected), "incorrectly sorted")
			Expect(slice).To(ConsistOf(original), "modified original")
		},
		Entry(nil, []string{"bar", "foo"}, []string{"bar", "foo"}),
		Entry(nil, []string{"foo", "bar"}, []string{"bar", "foo"}),
	)

	DescribeTable("stable sorting a copy using a compare function",
		func(slice []string, expected []string) {
			original := slices.Clone(slice)
			actual := StableSortedCopyFunc(slice, func(a, b string) int {
				return strings.Compare(a[:3], b[:3])
			})
			Expect(actual).To(ConsistOf(expected), "incorrectly sorted")
			Expect(slice).To(ConsistOf(original), "modified original")
		},
		Entry(nil, []string{"bar", "foo"}, []string{"bar", "foo"}),
		Entry(nil, []string{"foo2", "foo1", "bar"}, []string{"bar", "foo2", "foo1"}),
		Entry(nil, []string{"foo1", "foo2", "bar"}, []string{"bar", "foo1", "foo2"}),
	)

})
