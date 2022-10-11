//go:generate go run ./generate.go

package main

import (
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	spacelift "github.com/spacelift-io/pulumi-spacelift/provider"
	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/version"
)

//go:embed schema-embed.json
var pulumiSchema []byte

func main() {
	// Modify the path to point to the new provider
	tfbridge.Main("spacelift", version.Version, spacelift.Provider(), pulumiSchema)
}
