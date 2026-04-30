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
	"fmt"
	"log/slog"

	"github.com/thediveo/safe"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lazy slog values", func() {

	It("evaluates lazily", func() {
		DeferCleanup(slog.SetDefault, slog.Default())
		var buf safe.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{})))

		count := 0
		fn := func() any { count++; return fmt.Sprintf("LAZY-%d", count) }

		slog.Debug("debug", Lazy("lazy", fn))
		slog.Info("info", Lazy("lazy", fn))
		slog.Error("error", Lazy("lazy", fn))

		Expect(buf.String()).To(MatchRegexp(`msg=info lazy=LAZY-1\n.*msg=error lazy=LAZY-2`))
	})

})
