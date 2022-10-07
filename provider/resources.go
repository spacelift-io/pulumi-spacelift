package spacelift

import (
	"fmt"

	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/spacelift-io/pulumi-spacelift/provider/pkg/version"
	"github.com/spacelift-io/terraform-provider-spacelift/spacelift"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "spacelift"
	// modules:
	mainMod = "index" // the spacelift module
)

const (
	providerCommit  = "dev"
	providerVersion = "dev"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
//
//nolint:lll
func Provider() tfbridge.ProviderInfo {

	// Instantiate the Terraform provider
	// TODO: use the actual commit and version of the provider here from ldflags
	p := shimv2.NewProvider(spacelift.Provider(providerCommit, providerVersion)())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "spacelift",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Spacelift",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "spacelift-io",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/spacelift-io/pulumi-spacelift/releases",
		Description:       "A Pulumi package for creating and managing Spacelift resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "spacelift", "category/cloud", "category/infrastructure"},
		License:    "Apache-2.0",
		Homepage:   "https://spacelift.io",
		Repository: "git://github.com/spacelift-io/pulumi-spacelift.git",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "spacelift-io",
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
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"spacelift_aws_integration":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AwsIntegration")},
			"spacelift_aws_integration_attachment":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AwsIntegrationAttachment")},
			"spacelift_aws_role":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AwsRole")},
			"spacelift_azure_integration":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureIntegration")},
			"spacelift_azure_integration_attachment": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AzureIntegrationAttachment")},
			"spacelift_context":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Context")},
			"spacelift_context_attachment":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContextAttachment")},
			"spacelift_drift_detection":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DriftDetection")},
			"spacelift_environment_variable":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariable")},
			"spacelift_gcp_service_account":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "GcpServiceAccount")},
			"spacelift_module":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Module")},
			"spacelift_mounted_file":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Mountedfile")},
			"spacelift_policy":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Policy")},
			"spacelift_policy_attachment":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PolicyAttachment")},
			"spacelift_run":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Run")},
			"spacelift_space":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Space")},
			"spacelift_stack": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Stack"), Fields: map[string]*tfbridge.SchemaInfo{
				"pulumi": {
					CSharpName: "CSHARPPULUMI",
				},
			}},
			"spacelift_stack_aws_role":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StackAwsRole")},
			"spacelift_stack_destructor":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StackDestructor")},
			"spacelift_stack_gcp_service_account": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StackGcpServiceAccount")},
			"spacelift_vcs_agent_pool":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VcsAgentPool")},
			"spacelift_webhook":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Webook")},
			"spacelift_worker_pool":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkerPool")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"spacelift_account":                                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAccount")},
			"spacelift_aws_integration":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAwsIntegration")},
			"spacelift_aws_integration_attachment":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAwsIntegrationAttachment")},
			"spacelift_aws_integration_attachment_external_id": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAwsIntegrationAttachmentExternalId")},
			"spacelift_aws_role":                               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAwsRole")},
			"spacelift_azure_devops_integration":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureDevopsIntegration")},
			"spacelift_azure_integration":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureIntegration")},
			"spacelift_azure_integration_attachment":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureIntegrationAttachment")},
			"spacelift_bitbucket_cloud_integration":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBitbucketCloudIntegration")},
			"spacelift_bitbucket_datacenter_integration":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBitbucketDatacenterIntegration")},
			"spacelift_context":                                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getContext")},
			"spacelift_context_attachment":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getContextAttachment")},
			"spacelift_current_stack":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCurrentStack")},
			"spacelift_drift_detection":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDriftDetection")},
			"spacelift_environment_variable":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironmentVariable")},
			"spacelift_gcp_service_account":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGcpServiceAccount")},
			"spacelift_github_enterprise_integration":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGithubEnterpriseIntegration")},
			"spacelift_gitlab_integration":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGitlabIntegration")},
			"spacelift_ips":                                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIPs")},
			"spacelift_module":                                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getModule")},
			"spacelift_mounted_file":                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMountedfile")},
			"spacelift_policies":                               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPolicies")},
			"spacelift_policy":                                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPolicy")},
			"spacelift_space":                                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSpace")},
			"spacelift_stack":                                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStack")},
			"spacelift_stack_aws_role":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStackAwsRole")},
			"spacelift_stack_gcp_service_account":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStackGcpServiceAccount")},
			"spacelift_vcs_agent_pool":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVcsAgentPool")},
			"spacelift_vcs_agent_pools":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVcsAgentPools")},
			"spacelift_webhook":                                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWebhook")},
			"spacelift_worker_pool":                            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWorkerPool")},
			"spacelift_worker_pools":                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWorkerPools")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@spacelift-io/pulumi-spacelift",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/spacelift-io/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "Spacelift",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
