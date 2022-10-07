// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	spacelift "github.com/spacelift-io/pulumi-spacelift/provider"
	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/modify"
	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/version"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("spacelift", version.Version, spacelift.Provider())

	// read the language from the args passed in, once we've generated python we want to edit the file
	language := os.Args[1]
	if language == "python" {
		err := modify.ModifyPythonFiles()
		if err != nil {
			panic(err)
		}
	}
}
