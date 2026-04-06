// Copyright 2023 Harald Albrecht.
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

// FirstOk returns the first element produced by the it iterator as well as
// true; otherwise, if the iterator produced no values, it returns the zero
// value and false.
func FirstOk[V any](it iter.Seq[V]) (V, bool) {
	for v := range it {
		return v, true
	}
	var zero V
	return zero, false
}

// First2Ok returns the first (K, V) element produced by the it iterator as well
// as true; otherwise, if the iterator produced no values, it returns zero K and
// V, as well as false.
func First2Ok[K, V any](it iter.Seq2[K, V]) (K, V, bool) {
	for k, v := range it {
		return k, v, true
	}
	var zerok K
	var zerov V
	return zerok, zerov, false
}
