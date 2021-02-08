package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfgen"

	spacelift "github.com/spacelift-io/pulumi-spacelift/provider"
	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/version"
)

func main() {
	tfgen.Main("spacelift", version.Version, spacelift.Provider())
}
