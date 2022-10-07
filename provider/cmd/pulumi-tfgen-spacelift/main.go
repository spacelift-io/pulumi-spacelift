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
