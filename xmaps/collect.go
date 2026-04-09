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

package xmaps

import "iter"

// Collect key-value pairs from seq into a new map and returning it, aggregating
// values for the same key (as opposed to replacing keys with new values).
//
// Note: use [CollectT] in case you want to use a specific slice type other than
// []V.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K][]V {
	m := map[K][]V{}
	for k, v := range seq {
		m[k] = append(m[k], v)
	}
	return m
}

// Collect key-value paris from a seq into a new map with the specified value
// (slice) type T, aggregating values for the same key in slice type T. In
// contrast to [Collect], CollectT allows explicitly specifying the slice type T
// the map should use for aggregating multiple values for the same key, as Go
// has no way to infer this type.
func CollectT[T ~[]V, K comparable, V any](seq iter.Seq2[K, V]) map[K]T {
	m := map[K]T{}
	for k, v := range seq {
		m[k] = append(m[k], v)
	}
	return m
}
