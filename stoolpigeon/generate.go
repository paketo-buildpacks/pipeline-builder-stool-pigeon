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

package stoolpigeon

import (
	"fmt"
	"os"

	"github.com/buildpacks/libcnb"
	"github.com/buildpacks/libcnb/log"
)

type Generator struct {
	Logger log.Logger
}

func (Generator) Generate(context libcnb.GenerateContext) (libcnb.GenerateResult, error) {
	// here you can read the context.ApplicationPath folder
	// and create run.Dockerfile and build.Dockerfile in the context.OutputPath folder
	// and read metadata from the context.Extension struct

	// Just to use context to keep compiler happy =)
	fmt.Println(context.Extension.Info.ID)

	result := libcnb.NewGenerateResult()
	return result, nil
}

func ExampleGenerate() {
	generator := Generator{log.New(os.Stdout)}
	libcnb.ExtensionMain(nil, generator.Generate)
}
