//go:build ignore

package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func main() {
	version, found := os.LookupEnv("VERSION")
	if !found {
		log.Fatal("version not found")
	}

	schemaContents, err := os.ReadFile("./schema.json")
	if err != nil {
		log.Fatal(err)
	}

	var packageSpec schema.PackageSpec
	err = json.Unmarshal(schemaContents, &packageSpec)
	if err != nil {
		log.Fatalf("cannot deserialize schema: %v", err)
	}

	packageSpec.Version = version
	versionedContents, err := json.Marshal(packageSpec)
	if err != nil {
		log.Fatalf("cannot reserialize schema: %v", err)
	}

	// Clean up schema.go as it may be present & gitignored and tolerate an error if the file isn't present.
	err = os.Remove("./schema.go")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err)
	}

	err = os.WriteFile("./schema-embed.json", versionedContents, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
