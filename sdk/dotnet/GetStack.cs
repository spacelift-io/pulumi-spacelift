// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetStack
    {
        public static Task<GetStackResult> InvokeAsync(GetStackArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStackResult>("spacelift:index/getStack:getStack", args ?? new GetStackArgs(), options.WithVersion());
    }


    public sealed class GetStackArgs : Pulumi.InvokeArgs
    {
        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetStackArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStackResult
    {
        public readonly bool Administrative;
        public readonly bool Autodeploy;
        public readonly bool Autoretry;
        public readonly string AwsAssumeRolePolicyStatement;
        public readonly ImmutableArray<string> BeforeInits;
        public readonly string Branch;
        public readonly ImmutableArray<Outputs.GetStackCloudformationResult> Cloudformations;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetStackGitlabResult> Gitlabs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Labels;
        public readonly bool ManageState;
        public readonly string Name;
        public readonly string ProjectRoot;
        public readonly ImmutableArray<Outputs.GetStackPulumiResult> Pulumis;
        public readonly string Repository;
        public readonly string RunnerImage;
        public readonly string StackId;
        public readonly string TerraformVersion;
        public readonly string TerraformWorkspace;
        public readonly string WorkerPoolId;

        [OutputConstructor]
        private GetStackResult(
            bool administrative,

            bool autodeploy,

            bool autoretry,

            string awsAssumeRolePolicyStatement,

            ImmutableArray<string> beforeInits,

            string branch,

            ImmutableArray<Outputs.GetStackCloudformationResult> cloudformations,

            string description,

            ImmutableArray<Outputs.GetStackGitlabResult> gitlabs,

            string id,

            ImmutableArray<string> labels,

            bool manageState,

            string name,

            string projectRoot,

            ImmutableArray<Outputs.GetStackPulumiResult> pulumis,

            string repository,

            string runnerImage,

            string stackId,

            string terraformVersion,

            string terraformWorkspace,

            string workerPoolId)
        {
            Administrative = administrative;
            Autodeploy = autodeploy;
            Autoretry = autoretry;
            AwsAssumeRolePolicyStatement = awsAssumeRolePolicyStatement;
            BeforeInits = beforeInits;
            Branch = branch;
            Cloudformations = cloudformations;
            Description = description;
            Gitlabs = gitlabs;
            Id = id;
            Labels = labels;
            ManageState = manageState;
            Name = name;
            ProjectRoot = projectRoot;
            Pulumis = pulumis;
            Repository = repository;
            RunnerImage = runnerImage;
            StackId = stackId;
            TerraformVersion = terraformVersion;
            TerraformWorkspace = terraformWorkspace;
            WorkerPoolId = workerPoolId;
        }
    }
}
