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

// Filter returns a new slice containing only elements for which test returns
// true.
func Filter[S ~[]E, E any](s S, test func(e E) bool) S {
	if s == nil {
		return nil
	}
	r := make(S, 0, len(s))
	for _, e := range s {
		if !test(e) {
			continue
		}
		r = append(r, e)
	}
	return r
}
