// Copyright 2024 Harald Albrecht.
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

package xiter_test

import (
	"slices"

	"github.com/thediveo/nonstd/xiter"
	"github.com/thediveo/nonstd/xslices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Make[T any](value T) *T { return &value }

var _ = Describe("iterating over non-zero sequence elements", func() {

	DescribeTable("produces non-zero elements",
		func(seq []string, expected []string) {
			Expect(slices.Collect(xiter.AllUnzeros(xslices.AllValues(seq)))).
				To(Equal(expected))
		},
		Entry("preserves nil-ness", nil, nil),
		Entry(nil, []string{"foo", "", "bar", ""}, []string{"foo", "bar"}),
	)

	It("produces non-nils", func() {
		Expect(slices.Collect(xiter.WithoutNils(
			xslices.AllValues([]*string{Make("foo"), nil, Make("bar"), nil})))).
			To(Equal([]*string{Make("foo"), Make("bar")}))
	})

})
