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

import (
	"errors"
	"log/slog"

	"github.com/thediveo/safe"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("slogging errors", func() {

	It("handles nil", func() {
		DeferCleanup(slog.SetDefault, slog.Default())
		var buf safe.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{})))

		slog.Error("ouch", Error(nil))
		Expect(buf.String()).To(MatchRegexp(`msg=ouch err=<nil>`))
	})

	It("reports the error text", func() {
		DeferCleanup(slog.SetDefault, slog.Default())
		var buf safe.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{})))

		slog.Error("ouch", Error(errors.New("JKL305")))
		Expect(buf.String()).To(MatchRegexp(`msg=ouch err=JKL305`))
	})

})
