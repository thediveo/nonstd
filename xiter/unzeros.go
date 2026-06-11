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

package xiter

import "iter"

// AllUnzeros iterates over all non-zero, or “unzero”, values in the passed
// sequence.
func AllUnzeros[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	var zero V
	return Filter(seq, func(v V) bool { return v != zero })
}

// WithoutNils is the pointer-restricted twin to AllUnzeros and exists solely
// for readability as an alias for AllUnzeros. No pun intended on people named
// Nils.
func WithoutNils[E any, P ~*E](seq iter.Seq[P]) iter.Seq[P] {
	return AllUnzeros(seq)
}
