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

package xslices

import "slices"

// FindFuncOk returns the first value satisfying f(s[i]) and true; otherwise the
// type E zero value and false, if none do.
func FindFuncOk[S ~[]E, E any](s S, f func(E) bool) (E, bool) {
	idx := slices.IndexFunc(s, f)
	if idx < 0 {
		var zero E
		return zero, false
	}
	return s[idx], true
}

// FindFunc returns the first value satisfying f(s[i]); otherwise the type E
// zero value, if none do.
func FindFunc[S ~[]E, E any](s S, f func(E) bool) E {
	e, _ := FindFuncOk(s, f)
	return e
}
