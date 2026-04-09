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

import (
	"unicode"
	"unicode/utf8"
)

// CutWhitespace slices s around the first (white) space, returning the text
// before and after it. The found result reports whether a white space appears
// in s. If the white space does not appear in s, the function returns (s, "",
// false).
//
// White space is determined using [utf8.IsSpace] and any character with
// Unicode's White Space property.
func CutWhitespace(s string) (before, after string, found bool) {
	for idx, r := range s {
		if !unicode.IsSpace(r) {
			continue
		}
		return s[:idx], s[idx+utf8.RuneLen(r):], true
	}
	return s, "", false
}
