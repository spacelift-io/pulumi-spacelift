// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    /// <summary>
    /// `spacelift.Stack` combines source code and configuration to create a runtime environment where resources are managed. In this way it's similar to a stack in AWS CloudFormation, or a project on generic CI/CD platforms.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Spacelift = Pulumi.Spacelift;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Terraform stack using github.com as VCS
    ///         var k8s_cluster = new Spacelift.Stack("k8s-cluster", new Spacelift.StackArgs
    ///         {
    ///             Administrative = true,
    ///             Autodeploy = true,
    ///             Branch = "master",
    ///             Description = "Provisions a Kubernetes cluster",
    ///             ProjectRoot = "cluster",
    ///             Repository = "core-infra",
    ///             TerraformVersion = "0.12.6",
    ///         });
    ///         // Terraform stack using Bitbucket Cloud as VCS
    ///         var k8s_cluster_bitbucket_cloud = new Spacelift.Stack("k8s-cluster-bitbucket-cloud", new Spacelift.StackArgs
    ///         {
    ///             Administrative = true,
    ///             Autodeploy = true,
    ///             BitbucketCloud = new Spacelift.Inputs.StackBitbucketCloudArgs
    ///             {
    ///                 Namespace = "SPACELIFT",
    ///             },
    ///             Branch = "master",
    ///             Description = "Provisions a Kubernetes cluster",
    ///             ProjectRoot = "cluster",
    ///             Repository = "core-infra",
    ///             TerraformVersion = "0.12.6",
    ///         });
    ///         // Terraform stack using Bitbucket Data Center as VCS
    ///         var k8s_cluster_bitbucket_datacenter = new Spacelift.Stack("k8s-cluster-bitbucket-datacenter", new Spacelift.StackArgs
    ///         {
    ///             Administrative = true,
    ///             Autodeploy = true,
    ///             BitbucketDatacenter = new Spacelift.Inputs.StackBitbucketDatacenterArgs
    ///             {
    ///                 Namespace = "SPACELIFT",
    ///             },
    ///             Branch = "master",
    ///             Description = "Provisions a Kubernetes cluster",
    ///             ProjectRoot = "cluster",
    ///             Repository = "core-infra",
    ///             TerraformVersion = "0.12.6",
    ///         });
    ///         // Terraform stack using GitHub Enterprise as VCS
    ///         var k8s_cluster_github_enterprise = new Spacelift.Stack("k8s-cluster-github-enterprise", new Spacelift.StackArgs
    ///         {
    ///             Administrative = true,
    ///             Autodeploy = true,
    ///             Branch = "master",
    ///             Description = "Provisions a Kubernetes cluster",
    ///             GithubEnterprise = new Spacelift.Inputs.StackGithubEnterpriseArgs
    ///             {
    ///                 Namespace = "spacelift",
    ///             },
    ///             ProjectRoot = "cluster",
    ///             Repository = "core-infra",
    ///             TerraformVersion = "0.12.6",
    ///         });
    ///         // Terraform stack using GitLab as VCS
    ///         var k8s_cluster_gitlab = new Spacelift.Stack("k8s-cluster-gitlab", new Spacelift.StackArgs
    ///         {
    ///             Administrative = true,
    ///             Autodeploy = true,
    ///             Branch = "master",
    ///             Description = "Provisions a Kubernetes cluster",
    ///             Gitlab = new Spacelift.Inputs.StackGitlabArgs
    ///             {
    ///                 Namespace = "spacelift",
    ///             },
    ///             ProjectRoot = "cluster",
    ///             Repository = "core-infra",
    ///             TerraformVersion = "0.12.6",
    ///         });
    ///         // CloudFormation stack using github.com as VCS
    ///         var k8s_cluster_cloudformation = new Spacelift.Stack("k8s-cluster-cloudformation", new Spacelift.StackArgs
    ///         {
    ///             Autodeploy = true,
    ///             Branch = "master",
    ///             Cloudformation = new Spacelift.Inputs.StackCloudformationArgs
    ///             {
    ///                 EntryTemplateFile = "main.yaml",
    ///                 Region = "eu-central-1",
    ///                 StackName = "k8s-cluster",
    ///                 TemplateBucket = "s3://bucket",
    ///             },
    ///             Description = "Provisions a Kubernetes cluster",
    ///             ProjectRoot = "cluster",
    ///             Repository = "core-infra",
    ///         });
    ///         // Pulumi stack using github.com as VCS
    ///         var k8s_cluster_pulumi = new Spacelift.Stack("k8s-cluster-pulumi", new Spacelift.StackArgs
    ///         {
    ///             Autodeploy = true,
    ///             Branch = "master",
    ///             Description = "Provisions a Kubernetes cluster",
    ///             ProjectRoot = "cluster",
    ///             CSHARPPULUMI = new Spacelift.Inputs.StackPulumiArgs
    ///             {
    ///                 LoginUrl = "s3://pulumi-state-bucket",
    ///                 StackName = "kubernetes-core-services",
    ///             },
    ///             Repository = "core-infra",
    ///             RunnerImage = "public.ecr.aws/t0p9w2l5/runner-pulumi-javascript:latest",
    ///         });
    ///         // Kubernetes stack using github.com as VCS
    ///         var k8s_core_kubernetes = new Spacelift.Stack("k8s-core-kubernetes", new Spacelift.StackArgs
    ///         {
    ///             Autodeploy = true,
    ///             BeforeInits = 
    ///             {
    ///                 "aws eks update-kubeconfig --region us-east-2 --name k8s-cluster",
    ///             },
    ///             Branch = "master",
    ///             Description = "Shared cluster services (Datadog, Istio etc.)",
    ///             Kubernetes = new Spacelift.Inputs.StackKubernetesArgs
    ///             {
    ///                 Namespace = "core",
    ///             },
    ///             ProjectRoot = "core-services",
    ///             Repository = "core-infra",
    ///         });
    ///         // Ansible stack using github.com as VCS
    ///         var ansible_stack = new Spacelift.Stack("ansible-stack", new Spacelift.StackArgs
    ///         {
    ///             Ansible = new Spacelift.Inputs.StackAnsibleArgs
    ///             {
    ///                 Playbook = "main.yml",
    ///             },
    ///             Autodeploy = true,
    ///             Branch = "master",
    ///             Description = "Provisioning EC2 machines",
    ///             Repository = "ansible-playbooks",
    ///             RunnerImage = "public.ecr.aws/spacelift/runner-ansible:latest",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import spacelift:index/stack:Stack k8s_core $STACK_ID
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/stack:Stack")]
    public partial class Stack : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether this stack can manage others. Defaults to `false`.
        /// </summary>
        [Output("administrative")]
        public Output<bool?> Administrative { get; private set; } = null!;

        /// <summary>
        /// List of after-apply scripts
        /// </summary>
        [Output("afterApplies")]
        public Output<ImmutableArray<string>> AfterApplies { get; private set; } = null!;

        /// <summary>
        /// List of after-destroy scripts
        /// </summary>
        [Output("afterDestroys")]
        public Output<ImmutableArray<string>> AfterDestroys { get; private set; } = null!;

        /// <summary>
        /// List of after-init scripts
        /// </summary>
        [Output("afterInits")]
        public Output<ImmutableArray<string>> AfterInits { get; private set; } = null!;

        /// <summary>
        /// List of after-perform scripts
        /// </summary>
        [Output("afterPerforms")]
        public Output<ImmutableArray<string>> AfterPerforms { get; private set; } = null!;

        /// <summary>
        /// List of after-plan scripts
        /// </summary>
        [Output("afterPlans")]
        public Output<ImmutableArray<string>> AfterPlans { get; private set; } = null!;

        /// <summary>
        /// Ansible-specific configuration. Presence means this Stack is an Ansible Stack.
        /// </summary>
        [Output("ansible")]
        public Output<Outputs.StackAnsible?> Ansible { get; private set; } = null!;

        /// <summary>
        /// Indicates whether changes to this stack can be automatically deployed. Defaults to `false`.
        /// </summary>
        [Output("autodeploy")]
        public Output<bool?> Autodeploy { get; private set; } = null!;

        /// <summary>
        /// Indicates whether obsolete proposed changes should automatically be retried. Defaults to `false`.
        /// </summary>
        [Output("autoretry")]
        public Output<bool?> Autoretry { get; private set; } = null!;

        /// <summary>
        /// AWS IAM assume role policy statement setting up trust relationship
        /// </summary>
        [Output("awsAssumeRolePolicyStatement")]
        public Output<string> AwsAssumeRolePolicyStatement { get; private set; } = null!;

        /// <summary>
        /// Azure DevOps VCS settings
        /// </summary>
        [Output("azureDevops")]
        public Output<Outputs.StackAzureDevops?> AzureDevops { get; private set; } = null!;

        /// <summary>
        /// List of before-apply scripts
        /// </summary>
        [Output("beforeApplies")]
        public Output<ImmutableArray<string>> BeforeApplies { get; private set; } = null!;

        /// <summary>
        /// List of before-destroy scripts
        /// </summary>
        [Output("beforeDestroys")]
        public Output<ImmutableArray<string>> BeforeDestroys { get; private set; } = null!;

        /// <summary>
        /// List of before-init scripts
        /// </summary>
        [Output("beforeInits")]
        public Output<ImmutableArray<string>> BeforeInits { get; private set; } = null!;

        /// <summary>
        /// List of before-perform scripts
        /// </summary>
        [Output("beforePerforms")]
        public Output<ImmutableArray<string>> BeforePerforms { get; private set; } = null!;

        /// <summary>
        /// List of before-plan scripts
        /// </summary>
        [Output("beforePlans")]
        public Output<ImmutableArray<string>> BeforePlans { get; private set; } = null!;

        /// <summary>
        /// Bitbucket Cloud VCS settings
        /// </summary>
        [Output("bitbucketCloud")]
        public Output<Outputs.StackBitbucketCloud?> BitbucketCloud { get; private set; } = null!;

        /// <summary>
        /// Bitbucket Datacenter VCS settings
        /// </summary>
        [Output("bitbucketDatacenter")]
        public Output<Outputs.StackBitbucketDatacenter?> BitbucketDatacenter { get; private set; } = null!;

        /// <summary>
        /// GitHub branch to apply changes to
        /// </summary>
        [Output("branch")]
        public Output<string> Branch { get; private set; } = null!;

        /// <summary>
        /// CloudFormation-specific configuration. Presence means this Stack is a CloudFormation Stack.
        /// </summary>
        [Output("cloudformation")]
        public Output<Outputs.StackCloudformation?> Cloudformation { get; private set; } = null!;

        /// <summary>
        /// Free-form stack description for users
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether local preview runs can be triggered on this Stack. Defaults to `false`.
        /// </summary>
        [Output("enableLocalPreview")]
        public Output<bool?> EnableLocalPreview { get; private set; } = null!;

        /// <summary>
        /// Indicates whether GitHub users can deploy from the Checks API. Defaults to `true`.
        /// </summary>
        [Output("githubActionDeploy")]
        public Output<bool?> GithubActionDeploy { get; private set; } = null!;

        /// <summary>
        /// GitHub Enterprise (self-hosted) VCS settings
        /// </summary>
        [Output("githubEnterprise")]
        public Output<Outputs.StackGithubEnterprise?> GithubEnterprise { get; private set; } = null!;

        /// <summary>
        /// GitLab VCS settings
        /// </summary>
        [Output("gitlab")]
        public Output<Outputs.StackGitlab?> Gitlab { get; private set; } = null!;

        /// <summary>
        /// State file to upload when creating a new stack
        /// </summary>
        [Output("importState")]
        public Output<string?> ImportState { get; private set; } = null!;

        /// <summary>
        /// Path to the state file to upload when creating a new stack
        /// </summary>
        [Output("importStateFile")]
        public Output<string?> ImportStateFile { get; private set; } = null!;

        /// <summary>
        /// Kubernetes-specific configuration. Presence means this Stack is a Kubernetes Stack.
        /// </summary>
        [Output("kubernetes")]
        public Output<Outputs.StackKubernetes?> Kubernetes { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Determines if Spacelift should manage state for this stack. Defaults to `true`.
        /// </summary>
        [Output("manageState")]
        public Output<bool?> ManageState { get; private set; } = null!;

        /// <summary>
        /// Name of the stack - should be unique in one account
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Project root is the optional directory relative to the workspace root containing the entrypoint to the Stack.
        /// </summary>
        [Output("projectRoot")]
        public Output<string?> ProjectRoot { get; private set; } = null!;

        /// <summary>
        /// Protect this stack from accidental deletion. If set, attempts to delete this stack will fail. Defaults to `false`.
        /// </summary>
        [Output("protectFromDeletion")]
        public Output<bool?> ProtectFromDeletion { get; private set; } = null!;

        /// <summary>
        /// Pulumi-specific configuration. Presence means this Stack is a Pulumi Stack.
        /// </summary>
        [Output("pulumi")]
        public Output<Outputs.StackPulumi?> CSHARPPULUMI { get; private set; } = null!;

        /// <summary>
        /// Name of the repository, without the owner part
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// Name of the Docker image used to process Runs
        /// </summary>
        [Output("runnerImage")]
        public Output<string?> RunnerImage { get; private set; } = null!;

        [Output("showcase")]
        public Output<Outputs.StackShowcase?> Showcase { get; private set; } = null!;

        /// <summary>
        /// Allows setting the custom ID (slug) for the stack
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// ID (slug) of the space the stack is in
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// Terraform version to use
        /// </summary>
        [Output("terraformVersion")]
        public Output<string?> TerraformVersion { get; private set; } = null!;

        /// <summary>
        /// Terraform workspace to select
        /// </summary>
        [Output("terraformWorkspace")]
        public Output<string?> TerraformWorkspace { get; private set; } = null!;

        /// <summary>
        /// ID of the worker pool to use
        /// </summary>
        [Output("workerPoolId")]
        public Output<string?> WorkerPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a Stack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stack(string name, StackArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/stack:Stack", name, args ?? new StackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stack(string name, Input<string> id, StackState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/stack:Stack", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/spacelift-io/pulumi-spacelift/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Stack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stack Get(string name, Input<string> id, StackState? state = null, CustomResourceOptions? options = null)
        {
            return new Stack(name, id, state, options);
        }
    }

    public sealed class StackArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether this stack can manage others. Defaults to `false`.
        /// </summary>
        [Input("administrative")]
        public Input<bool>? Administrative { get; set; }

        [Input("afterApplies")]
        private InputList<string>? _afterApplies;

        /// <summary>
        /// List of after-apply scripts
        /// </summary>
        public InputList<string> AfterApplies
        {
            get => _afterApplies ?? (_afterApplies = new InputList<string>());
            set => _afterApplies = value;
        }

        [Input("afterDestroys")]
        private InputList<string>? _afterDestroys;

        /// <summary>
        /// List of after-destroy scripts
        /// </summary>
        public InputList<string> AfterDestroys
        {
            get => _afterDestroys ?? (_afterDestroys = new InputList<string>());
            set => _afterDestroys = value;
        }

        [Input("afterInits")]
        private InputList<string>? _afterInits;

        /// <summary>
        /// List of after-init scripts
        /// </summary>
        public InputList<string> AfterInits
        {
            get => _afterInits ?? (_afterInits = new InputList<string>());
            set => _afterInits = value;
        }

        [Input("afterPerforms")]
        private InputList<string>? _afterPerforms;

        /// <summary>
        /// List of after-perform scripts
        /// </summary>
        public InputList<string> AfterPerforms
        {
            get => _afterPerforms ?? (_afterPerforms = new InputList<string>());
            set => _afterPerforms = value;
        }

        [Input("afterPlans")]
        private InputList<string>? _afterPlans;

        /// <summary>
        /// List of after-plan scripts
        /// </summary>
        public InputList<string> AfterPlans
        {
            get => _afterPlans ?? (_afterPlans = new InputList<string>());
            set => _afterPlans = value;
        }

        /// <summary>
        /// Ansible-specific configuration. Presence means this Stack is an Ansible Stack.
        /// </summary>
        [Input("ansible")]
        public Input<Inputs.StackAnsibleArgs>? Ansible { get; set; }

        /// <summary>
        /// Indicates whether changes to this stack can be automatically deployed. Defaults to `false`.
        /// </summary>
        [Input("autodeploy")]
        public Input<bool>? Autodeploy { get; set; }

        /// <summary>
        /// Indicates whether obsolete proposed changes should automatically be retried. Defaults to `false`.
        /// </summary>
        [Input("autoretry")]
        public Input<bool>? Autoretry { get; set; }

        /// <summary>
        /// Azure DevOps VCS settings
        /// </summary>
        [Input("azureDevops")]
        public Input<Inputs.StackAzureDevopsArgs>? AzureDevops { get; set; }

        [Input("beforeApplies")]
        private InputList<string>? _beforeApplies;

        /// <summary>
        /// List of before-apply scripts
        /// </summary>
        public InputList<string> BeforeApplies
        {
            get => _beforeApplies ?? (_beforeApplies = new InputList<string>());
            set => _beforeApplies = value;
        }

        [Input("beforeDestroys")]
        private InputList<string>? _beforeDestroys;

        /// <summary>
        /// List of before-destroy scripts
        /// </summary>
        public InputList<string> BeforeDestroys
        {
            get => _beforeDestroys ?? (_beforeDestroys = new InputList<string>());
            set => _beforeDestroys = value;
        }

        [Input("beforeInits")]
        private InputList<string>? _beforeInits;

        /// <summary>
        /// List of before-init scripts
        /// </summary>
        public InputList<string> BeforeInits
        {
            get => _beforeInits ?? (_beforeInits = new InputList<string>());
            set => _beforeInits = value;
        }

        [Input("beforePerforms")]
        private InputList<string>? _beforePerforms;

        /// <summary>
        /// List of before-perform scripts
        /// </summary>
        public InputList<string> BeforePerforms
        {
            get => _beforePerforms ?? (_beforePerforms = new InputList<string>());
            set => _beforePerforms = value;
        }

        [Input("beforePlans")]
        private InputList<string>? _beforePlans;

        /// <summary>
        /// List of before-plan scripts
        /// </summary>
        public InputList<string> BeforePlans
        {
            get => _beforePlans ?? (_beforePlans = new InputList<string>());
            set => _beforePlans = value;
        }

        /// <summary>
        /// Bitbucket Cloud VCS settings
        /// </summary>
        [Input("bitbucketCloud")]
        public Input<Inputs.StackBitbucketCloudArgs>? BitbucketCloud { get; set; }

        /// <summary>
        /// Bitbucket Datacenter VCS settings
        /// </summary>
        [Input("bitbucketDatacenter")]
        public Input<Inputs.StackBitbucketDatacenterArgs>? BitbucketDatacenter { get; set; }

        /// <summary>
        /// GitHub branch to apply changes to
        /// </summary>
        [Input("branch", required: true)]
        public Input<string> Branch { get; set; } = null!;

        /// <summary>
        /// CloudFormation-specific configuration. Presence means this Stack is a CloudFormation Stack.
        /// </summary>
        [Input("cloudformation")]
        public Input<Inputs.StackCloudformationArgs>? Cloudformation { get; set; }

        /// <summary>
        /// Free-form stack description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether local preview runs can be triggered on this Stack. Defaults to `false`.
        /// </summary>
        [Input("enableLocalPreview")]
        public Input<bool>? EnableLocalPreview { get; set; }

        /// <summary>
        /// Indicates whether GitHub users can deploy from the Checks API. Defaults to `true`.
        /// </summary>
        [Input("githubActionDeploy")]
        public Input<bool>? GithubActionDeploy { get; set; }

        /// <summary>
        /// GitHub Enterprise (self-hosted) VCS settings
        /// </summary>
        [Input("githubEnterprise")]
        public Input<Inputs.StackGithubEnterpriseArgs>? GithubEnterprise { get; set; }

        /// <summary>
        /// GitLab VCS settings
        /// </summary>
        [Input("gitlab")]
        public Input<Inputs.StackGitlabArgs>? Gitlab { get; set; }

        /// <summary>
        /// State file to upload when creating a new stack
        /// </summary>
        [Input("importState")]
        public Input<string>? ImportState { get; set; }

        /// <summary>
        /// Path to the state file to upload when creating a new stack
        /// </summary>
        [Input("importStateFile")]
        public Input<string>? ImportStateFile { get; set; }

        /// <summary>
        /// Kubernetes-specific configuration. Presence means this Stack is a Kubernetes Stack.
        /// </summary>
        [Input("kubernetes")]
        public Input<Inputs.StackKubernetesArgs>? Kubernetes { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Determines if Spacelift should manage state for this stack. Defaults to `true`.
        /// </summary>
        [Input("manageState")]
        public Input<bool>? ManageState { get; set; }

        /// <summary>
        /// Name of the stack - should be unique in one account
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project root is the optional directory relative to the workspace root containing the entrypoint to the Stack.
        /// </summary>
        [Input("projectRoot")]
        public Input<string>? ProjectRoot { get; set; }

        /// <summary>
        /// Protect this stack from accidental deletion. If set, attempts to delete this stack will fail. Defaults to `false`.
        /// </summary>
        [Input("protectFromDeletion")]
        public Input<bool>? ProtectFromDeletion { get; set; }

        /// <summary>
        /// Pulumi-specific configuration. Presence means this Stack is a Pulumi Stack.
        /// </summary>
        [Input("pulumi")]
        public Input<Inputs.StackPulumiArgs>? CSHARPPULUMI { get; set; }

        /// <summary>
        /// Name of the repository, without the owner part
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// Name of the Docker image used to process Runs
        /// </summary>
        [Input("runnerImage")]
        public Input<string>? RunnerImage { get; set; }

        [Input("showcase")]
        public Input<Inputs.StackShowcaseArgs>? Showcase { get; set; }

        /// <summary>
        /// Allows setting the custom ID (slug) for the stack
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// ID (slug) of the space the stack is in
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// Terraform version to use
        /// </summary>
        [Input("terraformVersion")]
        public Input<string>? TerraformVersion { get; set; }

        /// <summary>
        /// Terraform workspace to select
        /// </summary>
        [Input("terraformWorkspace")]
        public Input<string>? TerraformWorkspace { get; set; }

        /// <summary>
        /// ID of the worker pool to use
        /// </summary>
        [Input("workerPoolId")]
        public Input<string>? WorkerPoolId { get; set; }

        public StackArgs()
        {
        }
    }

    public sealed class StackState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether this stack can manage others. Defaults to `false`.
        /// </summary>
        [Input("administrative")]
        public Input<bool>? Administrative { get; set; }

        [Input("afterApplies")]
        private InputList<string>? _afterApplies;

        /// <summary>
        /// List of after-apply scripts
        /// </summary>
        public InputList<string> AfterApplies
        {
            get => _afterApplies ?? (_afterApplies = new InputList<string>());
            set => _afterApplies = value;
        }

        [Input("afterDestroys")]
        private InputList<string>? _afterDestroys;

        /// <summary>
        /// List of after-destroy scripts
        /// </summary>
        public InputList<string> AfterDestroys
        {
            get => _afterDestroys ?? (_afterDestroys = new InputList<string>());
            set => _afterDestroys = value;
        }

        [Input("afterInits")]
        private InputList<string>? _afterInits;

        /// <summary>
        /// List of after-init scripts
        /// </summary>
        public InputList<string> AfterInits
        {
            get => _afterInits ?? (_afterInits = new InputList<string>());
            set => _afterInits = value;
        }

        [Input("afterPerforms")]
        private InputList<string>? _afterPerforms;

        /// <summary>
        /// List of after-perform scripts
        /// </summary>
        public InputList<string> AfterPerforms
        {
            get => _afterPerforms ?? (_afterPerforms = new InputList<string>());
            set => _afterPerforms = value;
        }

        [Input("afterPlans")]
        private InputList<string>? _afterPlans;

        /// <summary>
        /// List of after-plan scripts
        /// </summary>
        public InputList<string> AfterPlans
        {
            get => _afterPlans ?? (_afterPlans = new InputList<string>());
            set => _afterPlans = value;
        }

        /// <summary>
        /// Ansible-specific configuration. Presence means this Stack is an Ansible Stack.
        /// </summary>
        [Input("ansible")]
        public Input<Inputs.StackAnsibleGetArgs>? Ansible { get; set; }

        /// <summary>
        /// Indicates whether changes to this stack can be automatically deployed. Defaults to `false`.
        /// </summary>
        [Input("autodeploy")]
        public Input<bool>? Autodeploy { get; set; }

        /// <summary>
        /// Indicates whether obsolete proposed changes should automatically be retried. Defaults to `false`.
        /// </summary>
        [Input("autoretry")]
        public Input<bool>? Autoretry { get; set; }

        /// <summary>
        /// AWS IAM assume role policy statement setting up trust relationship
        /// </summary>
        [Input("awsAssumeRolePolicyStatement")]
        public Input<string>? AwsAssumeRolePolicyStatement { get; set; }

        /// <summary>
        /// Azure DevOps VCS settings
        /// </summary>
        [Input("azureDevops")]
        public Input<Inputs.StackAzureDevopsGetArgs>? AzureDevops { get; set; }

        [Input("beforeApplies")]
        private InputList<string>? _beforeApplies;

        /// <summary>
        /// List of before-apply scripts
        /// </summary>
        public InputList<string> BeforeApplies
        {
            get => _beforeApplies ?? (_beforeApplies = new InputList<string>());
            set => _beforeApplies = value;
        }

        [Input("beforeDestroys")]
        private InputList<string>? _beforeDestroys;

        /// <summary>
        /// List of before-destroy scripts
        /// </summary>
        public InputList<string> BeforeDestroys
        {
            get => _beforeDestroys ?? (_beforeDestroys = new InputList<string>());
            set => _beforeDestroys = value;
        }

        [Input("beforeInits")]
        private InputList<string>? _beforeInits;

        /// <summary>
        /// List of before-init scripts
        /// </summary>
        public InputList<string> BeforeInits
        {
            get => _beforeInits ?? (_beforeInits = new InputList<string>());
            set => _beforeInits = value;
        }

        [Input("beforePerforms")]
        private InputList<string>? _beforePerforms;

        /// <summary>
        /// List of before-perform scripts
        /// </summary>
        public InputList<string> BeforePerforms
        {
            get => _beforePerforms ?? (_beforePerforms = new InputList<string>());
            set => _beforePerforms = value;
        }

        [Input("beforePlans")]
        private InputList<string>? _beforePlans;

        /// <summary>
        /// List of before-plan scripts
        /// </summary>
        public InputList<string> BeforePlans
        {
            get => _beforePlans ?? (_beforePlans = new InputList<string>());
            set => _beforePlans = value;
        }

        /// <summary>
        /// Bitbucket Cloud VCS settings
        /// </summary>
        [Input("bitbucketCloud")]
        public Input<Inputs.StackBitbucketCloudGetArgs>? BitbucketCloud { get; set; }

        /// <summary>
        /// Bitbucket Datacenter VCS settings
        /// </summary>
        [Input("bitbucketDatacenter")]
        public Input<Inputs.StackBitbucketDatacenterGetArgs>? BitbucketDatacenter { get; set; }

        /// <summary>
        /// GitHub branch to apply changes to
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// CloudFormation-specific configuration. Presence means this Stack is a CloudFormation Stack.
        /// </summary>
        [Input("cloudformation")]
        public Input<Inputs.StackCloudformationGetArgs>? Cloudformation { get; set; }

        /// <summary>
        /// Free-form stack description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether local preview runs can be triggered on this Stack. Defaults to `false`.
        /// </summary>
        [Input("enableLocalPreview")]
        public Input<bool>? EnableLocalPreview { get; set; }

        /// <summary>
        /// Indicates whether GitHub users can deploy from the Checks API. Defaults to `true`.
        /// </summary>
        [Input("githubActionDeploy")]
        public Input<bool>? GithubActionDeploy { get; set; }

        /// <summary>
        /// GitHub Enterprise (self-hosted) VCS settings
        /// </summary>
        [Input("githubEnterprise")]
        public Input<Inputs.StackGithubEnterpriseGetArgs>? GithubEnterprise { get; set; }

        /// <summary>
        /// GitLab VCS settings
        /// </summary>
        [Input("gitlab")]
        public Input<Inputs.StackGitlabGetArgs>? Gitlab { get; set; }

        /// <summary>
        /// State file to upload when creating a new stack
        /// </summary>
        [Input("importState")]
        public Input<string>? ImportState { get; set; }

        /// <summary>
        /// Path to the state file to upload when creating a new stack
        /// </summary>
        [Input("importStateFile")]
        public Input<string>? ImportStateFile { get; set; }

        /// <summary>
        /// Kubernetes-specific configuration. Presence means this Stack is a Kubernetes Stack.
        /// </summary>
        [Input("kubernetes")]
        public Input<Inputs.StackKubernetesGetArgs>? Kubernetes { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Determines if Spacelift should manage state for this stack. Defaults to `true`.
        /// </summary>
        [Input("manageState")]
        public Input<bool>? ManageState { get; set; }

        /// <summary>
        /// Name of the stack - should be unique in one account
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project root is the optional directory relative to the workspace root containing the entrypoint to the Stack.
        /// </summary>
        [Input("projectRoot")]
        public Input<string>? ProjectRoot { get; set; }

        /// <summary>
        /// Protect this stack from accidental deletion. If set, attempts to delete this stack will fail. Defaults to `false`.
        /// </summary>
        [Input("protectFromDeletion")]
        public Input<bool>? ProtectFromDeletion { get; set; }

        /// <summary>
        /// Pulumi-specific configuration. Presence means this Stack is a Pulumi Stack.
        /// </summary>
        [Input("pulumi")]
        public Input<Inputs.StackPulumiGetArgs>? CSHARPPULUMI { get; set; }

        /// <summary>
        /// Name of the repository, without the owner part
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// Name of the Docker image used to process Runs
        /// </summary>
        [Input("runnerImage")]
        public Input<string>? RunnerImage { get; set; }

        [Input("showcase")]
        public Input<Inputs.StackShowcaseGetArgs>? Showcase { get; set; }

        /// <summary>
        /// Allows setting the custom ID (slug) for the stack
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// ID (slug) of the space the stack is in
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// Terraform version to use
        /// </summary>
        [Input("terraformVersion")]
        public Input<string>? TerraformVersion { get; set; }

        /// <summary>
        /// Terraform workspace to select
        /// </summary>
        [Input("terraformWorkspace")]
        public Input<string>? TerraformWorkspace { get; set; }

        /// <summary>
        /// ID of the worker pool to use
        /// </summary>
        [Input("workerPoolId")]
        public Input<string>? WorkerPoolId { get; set; }

        public StackState()
        {
        }
    }
}
