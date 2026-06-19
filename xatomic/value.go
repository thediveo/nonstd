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

package xatomic

import "sync/atomic"

// Value represents an atomic value that can be loaded, stored, and swapped.
//
// Note: because this type represents atomic values and not pointers to values,
// it is impossible to CompareAndSwap because a Value zero value has an internal
// zero value atomic.Pointer that never matches any “old” value. Also,
// atomic.Pointer.CompareAndSwap works on pointers, not values, but we would
// like to work on atomic value comparism and swapping ... which isn't
// supported. For these cases, go back to good old mutexes.
type Value[T any] struct {
	p atomic.Pointer[T]
}

// Load returns the stored value.
func (v *Value[T]) Load() T {
	p := v.p.Load()
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// Store atomically a new value.
func (v *Value[T]) Store(value T) {
	v.p.Store(&value)
}

// Swap atomically the old with a new value, returning the old value.
func (v *Value[T]) Swap(value T) (old T) {
	p := v.p.Swap(&value)
	if p == nil {
		return old // which actually is the zero value at this point
	}
	return *p
}
