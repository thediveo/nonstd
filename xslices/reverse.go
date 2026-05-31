// Copyright 2023 Harald Albrecht.
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

import "slices"

// ReverseCopy returns a copy of the specified slice with the sequence of
// elements reversed. It differs from [slices.Reverse] in that it works on a
// copy and not in place.
func ReverseCopy[S ~[]E, E any](s S) S {
	s = slices.Clone(s)
	slices.Reverse(s)
	return s
}
