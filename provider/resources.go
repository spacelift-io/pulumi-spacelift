package spacelift

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"

	"github.com/spacelift-io/terraform-provider-spacelift/spacelift"

	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/version"
)

const (
	// packages:
	spaceliftPkg = "spacelift"
	// modules:
	spaceliftMod = "index"
)

var namespaceMap = map[string]string{
	spaceliftPkg: "Spacelift",
}

// makeMember manufactures a type token for the package and the given module and type.  It automatically
// uses the Spacelift package and names the file by simply lower casing the resource's first character.
func makeMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(spaceliftPkg + ":" + token + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.
func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
}

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(spacelift.Provider())
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "spacelift",
		Description: "A Pulumi package for creating and managing Spacelift resources.",
		Keywords:    []string{"pulumi", "spacelift"},
		License:     "Apache-2.0",
		Homepage:    "https://spacelift.io",
		GitHubOrg:   "spacelift-io",
		Repository:  "https://github.com/spacelift-io/pulumi-spacelift",
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key_endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SPACELIFT_API_KEY_ENDPOINT"},
				},
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.False(),
			},
			"api_key_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SPACELIFT_API_KEY_ID"},
				},
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.False(),
			},
			"api_key_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SPACELIFT_API_KEY_SECRET"},
				},
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.True(),
			},
			"api_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SPACELIFT_API_TOKEN"},
				},
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.True(),
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"spacelift_aws_role":                  {Tok: makeResource(spaceliftMod, "AwsRole")},
			"spacelift_context_attachment":        {Tok: makeResource(spaceliftMod, "ContextAttachment")},
			"spacelift_context":                   {Tok: makeResource(spaceliftMod, "Context")},
			"spacelift_environment_variable":      {Tok: makeResource(spaceliftMod, "EnvironmentVariable")},
			"spacelift_gcp_service_account":       {Tok: makeResource(spaceliftMod, "GcpServiceAccount")},
			"spacelift_module":                    {Tok: makeResource(spaceliftMod, "Module")},
			"spacelift_mounted_file":              {Tok: makeResource(spaceliftMod, "MountedFile")},
			"spacelift_policy_attachment":         {Tok: makeResource(spaceliftMod, "PolicyAttachment")},
			"spacelift_policy":                    {Tok: makeResource(spaceliftMod, "Policy")},
			"spacelift_stack":                     {Tok: makeResource(spaceliftMod, "Stack")},
			"spacelift_stack_aws_role":            {Tok: makeResource(spaceliftMod, "StackAwsRole")},
			"spacelift_stack_gcp_service_account": {Tok: makeResource(spaceliftMod, "StackGcpServiceAccount")},
			"spacelift_webhook":                   {Tok: makeResource(spaceliftMod, "Webhook")},
			"spacelift_worker_pool":               {Tok: makeResource(spaceliftMod, "WorkerPool")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"spacelift_aws_role":                  {Tok: makeDataSource(spaceliftMod, "getAwsRole")},
			"spacelift_context_attachment":        {Tok: makeDataSource(spaceliftMod, "getContextAttachment")},
			"spacelift_context":                   {Tok: makeDataSource(spaceliftMod, "getContext")},
			"spacelift_current_stack":             {Tok: makeDataSource(spaceliftMod, "getCurrentStack")},
			"spacelift_environment_variable":      {Tok: makeDataSource(spaceliftMod, "getEnvironmentVariable")},
			"spacelift_gcp_service_account":       {Tok: makeDataSource(spaceliftMod, "getGcpServiceAccount")},
			"spacelift_ips":                       {Tok: makeDataSource(spaceliftMod, "getIps")},
			"spacelift_module":                    {Tok: makeDataSource(spaceliftMod, "getModule")},
			"spacelift_mounted_file":              {Tok: makeDataSource(spaceliftMod, "getMountedFile")},
			"spacelift_policy":                    {Tok: makeDataSource(spaceliftMod, "getPolicy")},
			"spacelift_stack":                     {Tok: makeDataSource(spaceliftMod, "getStack")},
			"spacelift_webhook":                   {Tok: makeDataSource(spaceliftMod, "getWebhook")},
			"spacelift_stack_aws_role":            {Tok: makeDataSource(spaceliftMod, "getStackAwsRole")},
			"spacelift_stack_gcp_service_account": {Tok: makeDataSource(spaceliftMod, "getStackGcpServiceAccount")},
			"spacelift_worker_pool":               {Tok: makeDataSource(spaceliftMod, "getWorkerPool")},
			"spacelift_worker_pools":              {Tok: makeDataSource(spaceliftMod, "getWorkerPools")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@spacelift-io/pulumi-spacelift",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/spacelift-io/spacelift-%[1]s/sdk/", spaceliftPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				spaceliftPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}

	return prov
}
