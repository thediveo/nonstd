// Copyright 2026 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xslices

// Map transforms the elements of the passed slice using the specified function,
// returning a new slice of the transformed elements. It preserves nil-ness.
func Map[S ~[]E, E, F any](s S, fn func(e E) F) []F {
	if s == nil {
		return nil
	}
	r := make([]F, len(s))
	for idx, e := range s {
		r[idx] = fn(e)
	}
	return r
}
