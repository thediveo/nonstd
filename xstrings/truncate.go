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

package xstrings

// Truncate a string after at most maxrunes.
func Truncate(s string, maxrunes int) string {
	if maxrunes < 0 {
		panic("xstrings.Truncate: maximum length cannot be <0")
	}
	runes := 0
	for idx := range s {
		if runes == maxrunes {
			return s[:idx]
		}
		runes++
	}
	return s
}
