// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Outputs
{

    [OutputType]
    public sealed class GetStacksStackResult
    {
        public readonly ImmutableArray<string> AdditionalProjectGlobs;
        public readonly bool Administrative;
        public readonly ImmutableArray<string> AfterApplies;
        public readonly ImmutableArray<string> AfterDestroys;
        public readonly ImmutableArray<string> AfterInits;
        public readonly ImmutableArray<string> AfterPerforms;
        public readonly ImmutableArray<string> AfterPlans;
        public readonly ImmutableArray<string> AfterRuns;
        public readonly ImmutableArray<Outputs.GetStacksStackAnsibleResult> Ansibles;
        public readonly bool Autodeploy;
        public readonly bool Autoretry;
        public readonly string AwsAssumeRolePolicyStatement;
        public readonly ImmutableArray<Outputs.GetStacksStackAzureDevopResult> AzureDevops;
        public readonly ImmutableArray<string> BeforeApplies;
        public readonly ImmutableArray<string> BeforeDestroys;
        public readonly ImmutableArray<string> BeforeInits;
        public readonly ImmutableArray<string> BeforePerforms;
        public readonly ImmutableArray<string> BeforePlans;
        public readonly ImmutableArray<Outputs.GetStacksStackBitbucketCloudResult> BitbucketClouds;
        public readonly ImmutableArray<Outputs.GetStacksStackBitbucketDatacenterResult> BitbucketDatacenters;
        public readonly string Branch;
        public readonly ImmutableArray<Outputs.GetStacksStackCloudformationResult> Cloudformations;
        public readonly string Description;
        public readonly bool EnableLocalPreview;
        public readonly ImmutableArray<Outputs.GetStacksStackGithubEnterpriseResult> GithubEnterprises;
        public readonly ImmutableArray<Outputs.GetStacksStackGitlabResult> Gitlabs;
        public readonly ImmutableArray<Outputs.GetStacksStackKuberneteResult> Kubernetes;
        public readonly ImmutableArray<string> Labels;
        public readonly bool ManageState;
        public readonly string Name;
        public readonly string ProjectRoot;
        public readonly bool ProtectFromDeletion;
        public readonly ImmutableArray<Outputs.GetStacksStackPulumiResult> Pulumis;
        public readonly ImmutableArray<Outputs.GetStacksStackRawGitResult> RawGits;
        public readonly string Repository;
        public readonly string RunnerImage;
        public readonly ImmutableArray<Outputs.GetStacksStackShowcaseResult> Showcases;
        public readonly string SpaceId;
        public readonly string StackId;
        public readonly bool TerraformExternalStateAccess;
        public readonly bool TerraformSmartSanitization;
        public readonly string TerraformVersion;
        public readonly string TerraformWorkflowTool;
        public readonly string TerraformWorkspace;
        public readonly string WorkerPoolId;

        [OutputConstructor]
        private GetStacksStackResult(
            ImmutableArray<string> additionalProjectGlobs,

            bool administrative,

            ImmutableArray<string> afterApplies,

            ImmutableArray<string> afterDestroys,

            ImmutableArray<string> afterInits,

            ImmutableArray<string> afterPerforms,

            ImmutableArray<string> afterPlans,

            ImmutableArray<string> afterRuns,

            ImmutableArray<Outputs.GetStacksStackAnsibleResult> ansibles,

            bool autodeploy,

            bool autoretry,

            string awsAssumeRolePolicyStatement,

            ImmutableArray<Outputs.GetStacksStackAzureDevopResult> azureDevops,

            ImmutableArray<string> beforeApplies,

            ImmutableArray<string> beforeDestroys,

            ImmutableArray<string> beforeInits,

            ImmutableArray<string> beforePerforms,

            ImmutableArray<string> beforePlans,

            ImmutableArray<Outputs.GetStacksStackBitbucketCloudResult> bitbucketClouds,

            ImmutableArray<Outputs.GetStacksStackBitbucketDatacenterResult> bitbucketDatacenters,

            string branch,

            ImmutableArray<Outputs.GetStacksStackCloudformationResult> cloudformations,

            string description,

            bool enableLocalPreview,

            ImmutableArray<Outputs.GetStacksStackGithubEnterpriseResult> githubEnterprises,

            ImmutableArray<Outputs.GetStacksStackGitlabResult> gitlabs,

            ImmutableArray<Outputs.GetStacksStackKuberneteResult> kubernetes,

            ImmutableArray<string> labels,

            bool manageState,

            string name,

            string projectRoot,

            bool protectFromDeletion,

            ImmutableArray<Outputs.GetStacksStackPulumiResult> pulumis,

            ImmutableArray<Outputs.GetStacksStackRawGitResult> rawGits,

            string repository,

            string runnerImage,

            ImmutableArray<Outputs.GetStacksStackShowcaseResult> showcases,

            string spaceId,

            string stackId,

            bool terraformExternalStateAccess,

            bool terraformSmartSanitization,

            string terraformVersion,

            string terraformWorkflowTool,

            string terraformWorkspace,

            string workerPoolId)
        {
            AdditionalProjectGlobs = additionalProjectGlobs;
            Administrative = administrative;
            AfterApplies = afterApplies;
            AfterDestroys = afterDestroys;
            AfterInits = afterInits;
            AfterPerforms = afterPerforms;
            AfterPlans = afterPlans;
            AfterRuns = afterRuns;
            Ansibles = ansibles;
            Autodeploy = autodeploy;
            Autoretry = autoretry;
            AwsAssumeRolePolicyStatement = awsAssumeRolePolicyStatement;
            AzureDevops = azureDevops;
            BeforeApplies = beforeApplies;
            BeforeDestroys = beforeDestroys;
            BeforeInits = beforeInits;
            BeforePerforms = beforePerforms;
            BeforePlans = beforePlans;
            BitbucketClouds = bitbucketClouds;
            BitbucketDatacenters = bitbucketDatacenters;
            Branch = branch;
            Cloudformations = cloudformations;
            Description = description;
            EnableLocalPreview = enableLocalPreview;
            GithubEnterprises = githubEnterprises;
            Gitlabs = gitlabs;
            Kubernetes = kubernetes;
            Labels = labels;
            ManageState = manageState;
            Name = name;
            ProjectRoot = projectRoot;
            ProtectFromDeletion = protectFromDeletion;
            Pulumis = pulumis;
            RawGits = rawGits;
            Repository = repository;
            RunnerImage = runnerImage;
            Showcases = showcases;
            SpaceId = spaceId;
            StackId = stackId;
            TerraformExternalStateAccess = terraformExternalStateAccess;
            TerraformSmartSanitization = terraformSmartSanitization;
            TerraformVersion = terraformVersion;
            TerraformWorkflowTool = terraformWorkflowTool;
            TerraformWorkspace = terraformWorkspace;
            WorkerPoolId = workerPoolId;
        }
    }
}
