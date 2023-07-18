/*
 * Copyright 2023 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stoolpigeon_test

import (
	"os"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/pipeline-builder-stool-pigeon/stoolpigeon"
	"github.com/sclevine/spec"
)

func testGenerate(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx      libcnb.GenerateContext
		generate stoolpigeon.Generator
	)

	it.Before(func() {
		var err error
		ctx.ApplicationPath, err = os.MkdirTemp("", "stoolpigeon")
		Expect(err).NotTo(HaveOccurred())
	})

	it.After(func() {
		Expect(os.RemoveAll(ctx.ApplicationPath)).To(Succeed())
	})

	it("passes always", func() {
		Expect(generate.Generate(ctx)).To(Equal(libcnb.GenerateResult{}))
	})
}
