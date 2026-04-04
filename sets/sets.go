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

package sets

import "fmt"

// Set of elements of type E, which must be comparable.
type Set[E comparable] map[E]struct{}

// New returns a new set for elements of type E, optionally initialized with the
// passed elements.
func New[E comparable](es ...E) Set[E] {
	s := Set[E]{}
	for _, e := range es {
		s[e] = struct{}{}
	}
	return s
}

// Add an element to the set. It is no error to add the same element multiple
// times.
func (s Set[E]) Add(e E) {
	s[e] = struct{}{}
}

// Delete an element from the set.
func (s Set[E]) Delete(e E) {
	delete(s, e)
}

// Contains returns true if the set contains the specified element, otherwise
// false.
func (s Set[E]) Contains(e E) bool {
	_, ok := s[e]
	return ok
}

// Elements returns the elements of the set, in undefined order.
func (s Set[E]) Elements() []E {
	els := make([]E, 0, len(s))
	for e := range s {
		els = append(els, e)
	}
	return els
}

// String returns a textual representation of the set elements, in form of the
// standard Go slice notation. Please note that the order of elements is
// undefined.
func (s Set[E]) String() string {
	return fmt.Sprint(s.Elements())
}
