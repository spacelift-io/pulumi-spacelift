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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// using Spacelift = Pulumi.Spacelift;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var k8s_coreStack = new Spacelift.Stack("k8s-coreStack", new()
    ///     {
    ///         Branch = "master",
    ///         Repository = "core-infra",
    ///     });
    /// 
    ///     var k8s_coreStackGcpServiceAccount = new Spacelift.StackGcpServiceAccount("k8s-coreStackGcpServiceAccount", new()
    ///     {
    ///         StackId = k8s_coreStack.Id,
    ///         TokenScopes = new[]
    ///         {
    ///             "https://www.googleapis.com/auth/compute",
    ///             "https://www.googleapis.com/auth/cloud-platform",
    ///             "https://www.googleapis.com/auth/devstorage.full_control",
    ///         },
    ///     });
    /// 
    ///     var k8s_coreProject = new Gcp.Organizations.Project("k8s-coreProject", new()
    ///     {
    ///         ProjectId = "unicorn-k8s-core",
    ///         OrgId = @var.Gcp_organization_id,
    ///     });
    /// 
    ///     var k8s_coreIAMMember = new Gcp.Projects.IAMMember("k8s-coreIAMMember", new()
    ///     {
    ///         Project = k8s_coreProject.Id,
    ///         Role = "roles/owner",
    ///         Member = k8s_coreStackGcpServiceAccount.ServiceAccountEmail.Apply(serviceAccountEmail =&gt; $"serviceAccount:{serviceAccountEmail}"),
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount")]
    public partial class StackGcpServiceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the module which uses GCP service account credentials
        /// </summary>
        [Output("moduleId")]
        public Output<string?> ModuleId { get; private set; } = null!;

        /// <summary>
        /// Email address of the GCP service account dedicated for this stack
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// ID of the stack which uses GCP service account credentials
        /// </summary>
        [Output("stackId")]
        public Output<string?> StackId { get; private set; } = null!;

        /// <summary>
        /// List of scopes that will be requested when generating temporary GCP service account credentials
        /// </summary>
        [Output("tokenScopes")]
        public Output<ImmutableArray<string>> TokenScopes { get; private set; } = null!;


        /// <summary>
        /// Create a StackGcpServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackGcpServiceAccount(string name, StackGcpServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount", name, args ?? new StackGcpServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StackGcpServiceAccount(string name, Input<string> id, StackGcpServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://downloads.spacelift.io/pulumi-plugins",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StackGcpServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackGcpServiceAccount Get(string name, Input<string> id, StackGcpServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new StackGcpServiceAccount(name, id, state, options);
        }
    }

    public sealed class StackGcpServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the module which uses GCP service account credentials
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// ID of the stack which uses GCP service account credentials
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("tokenScopes", required: true)]
        private InputList<string>? _tokenScopes;

        /// <summary>
        /// List of scopes that will be requested when generating temporary GCP service account credentials
        /// </summary>
        public InputList<string> TokenScopes
        {
            get => _tokenScopes ?? (_tokenScopes = new InputList<string>());
            set => _tokenScopes = value;
        }

        public StackGcpServiceAccountArgs()
        {
        }
        public static new StackGcpServiceAccountArgs Empty => new StackGcpServiceAccountArgs();
    }

    public sealed class StackGcpServiceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the module which uses GCP service account credentials
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Email address of the GCP service account dedicated for this stack
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// ID of the stack which uses GCP service account credentials
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("tokenScopes")]
        private InputList<string>? _tokenScopes;

        /// <summary>
        /// List of scopes that will be requested when generating temporary GCP service account credentials
        /// </summary>
        public InputList<string> TokenScopes
        {
            get => _tokenScopes ?? (_tokenScopes = new InputList<string>());
            set => _tokenScopes = value;
        }

        public StackGcpServiceAccountState()
        {
        }
        public static new StackGcpServiceAccountState Empty => new StackGcpServiceAccountState();
    }
}
