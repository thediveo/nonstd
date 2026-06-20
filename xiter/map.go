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

package xiter

import "iter"

// Map transforms the passed sequence into a different sequence using the
// specified function. It preserves nil-ness.
func Map[E, F any](seq iter.Seq[E], fn func(e E) F) iter.Seq[F] {
	return func(yield func(F) bool) {
		for e := range seq {
			if !yield(fn(e)) {
				break
			}
		}
	}
}

// Map2 transforms the passed sequence into a different sequence using the
// specified function. It preserves nil-ness.
func Map2[K, V, X, Y any](seq iter.Seq2[K, V], fn func(k K, v V) (X, Y)) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for k, v := range seq {
			if !yield(fn(k, v)) {
				break
			}
		}
	}
}
