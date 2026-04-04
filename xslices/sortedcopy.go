// Copyright 2022, 2025 Harald Albrecht.
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
	"cmp"
	"slices"
)

// SortedCopy returns a sorted copy of a slice of any odered type in ascending
// order. It preserves nil-ness.
func SortedCopy[S ~[]E, E cmp.Ordered](s S) S {
	s = slices.Clone(s)
	slices.Sort(s)
	return s
}

// SortedCopyFunc returns a sorted copy of a slice, in an order determined by
// the cmp function. This sort is not guaranteed to be stable. It preserves
// nil-ness.
func SortedCopyFunc[S ~[]E, E any](s S, cmp func(a, b E) int) S {
	s = slices.Clone(s)
	slices.SortFunc(s, cmp)
	return s
}

// StableSortedCopyFunc returns a sorted copy of a slice, in an order determined
// by the cmp function while keeping the original order of equal elements. It
// preserves nil-ness.
func StableSortedCopyFunc[S ~[]E, E any](s S, cmp func(a, b E) int) S {
	s = slices.Clone(s)
	slices.SortStableFunc(s, cmp)
	return s
}
