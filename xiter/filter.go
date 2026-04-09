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

// Filter iterates of the values of seq for which test returns true.
func Filter[V any](seq iter.Seq[V], test func(v V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !test(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Filter2 iterates of the (k,v) pairs of seq for which test returns true.
func Filter2[K, V any](seq iter.Seq2[K, V], test func(k K, v V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !test(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
