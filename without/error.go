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

package without

// Error throws away the final passed error value, returning only the first
// passed function result value.
func Error[R any](r R, _ error) R { return r }

// Error2 throws away the final passed error value, returning only the first two
// passed function result values.
func Error2[R1, R2 any](r1 R1, r2 R2, _ error) (R1, R2) { return r1, r2 }

// Error3 throws away the final passed error value, returning only the first three
// passed function result values.
func Error3[R1, R2, R3 any](r1 R1, r2 R2, r3 R3, _ error) (R1, R2, R3) { return r1, r2, r3 }
