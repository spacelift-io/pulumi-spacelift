package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"

	spacelift "github.com/spacelift-io/pulumi-spacelift/provider"
	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/version"
)

//go:generate go run ./generate.go

func main() {
	tfbridge.Main("spacelift", version.Version, spacelift.Provider(), pulumiSchema)
}
