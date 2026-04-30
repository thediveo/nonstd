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

// LazyValuer produces a potentially computationally expensive slog.Value only
// on demand when LogValue is called. Please note that LazyValuer does not cache
// values.
type LazyValuer[T any] func() T

// LogValue calls the lazy evaluation function, returning its result wrapped in
// a slog.Value.
func (v LazyValuer[T]) LogValue() slog.Value { return slog.AnyValue(v()) }

var _ slog.LogValuer = (LazyValuer[int])(nil)

// Lazy returns a slog.Attr that will evaluate its attribute value only when
// actually needed.
func Lazy[T any](key string, fn LazyValuer[T]) slog.Attr {
	return slog.Attr{Key: key, Value: slog.AnyValue(fn)}
}
