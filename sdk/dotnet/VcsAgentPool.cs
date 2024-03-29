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
    /// `spacelift.VcsAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Spacelift = Pulumi.Spacelift;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ghe = new Spacelift.VcsAgentPool("ghe", new()
    ///     {
    ///         Description = "VCS agent pool for our internal GitHub Enterprise",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import spacelift:index/vcsAgentPool:VcsAgentPool ghe $VCS_AGENT_POOL_ID
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/vcsAgentPool:VcsAgentPool")]
    public partial class VcsAgentPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// VCS agent pool configuration, encoded using base64
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// Free-form VCS agent pool description for users
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the VCS agent pool, must be unique within an account
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a VcsAgentPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VcsAgentPool(string name, VcsAgentPoolArgs? args = null, CustomResourceOptions? options = null)
            : base("spacelift:index/vcsAgentPool:VcsAgentPool", name, args ?? new VcsAgentPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VcsAgentPool(string name, Input<string> id, VcsAgentPoolState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/vcsAgentPool:VcsAgentPool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://downloads.spacelift.io/pulumi-plugins",
                AdditionalSecretOutputs =
                {
                    "config",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VcsAgentPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VcsAgentPool Get(string name, Input<string> id, VcsAgentPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new VcsAgentPool(name, id, state, options);
        }
    }

    public sealed class VcsAgentPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Free-form VCS agent pool description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the VCS agent pool, must be unique within an account
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VcsAgentPoolArgs()
        {
        }
        public static new VcsAgentPoolArgs Empty => new VcsAgentPoolArgs();
    }

    public sealed class VcsAgentPoolState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private Input<string>? _config;

        /// <summary>
        /// VCS agent pool configuration, encoded using base64
        /// </summary>
        public Input<string>? Config
        {
            get => _config;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _config = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Free-form VCS agent pool description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the VCS agent pool, must be unique within an account
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VcsAgentPoolState()
        {
        }
        public static new VcsAgentPoolState Empty => new VcsAgentPoolState();
    }
}
