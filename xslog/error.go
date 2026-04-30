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

package xslog

import "log/slog"

// Error returns a slog attribute named “err” with the specified error value.
// This is nothing more than just a wrapper for [slog.Any] to unify the
// attribute name and clearly convey semantics, in the spirit of [slog.String]
// and company.
func Error(err error) slog.Attr {
	return slog.Any("err", err)
}
