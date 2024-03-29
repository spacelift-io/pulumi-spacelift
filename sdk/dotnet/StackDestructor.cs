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
    /// `spacelift.StackDestructor` is used to destroy the resources of a Stack before deleting it. `depends_on` should be used to make sure that all necessary resources (environment variables, roles, integrations, etc.) are still in place when the destruction run is executed. **Note:** Destroying this resource will delete the resources in the stack. If this resource needs to be deleted and the resources in the stacks are to be preserved, ensure that the `deactivated` attribute is set to `true`.
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
    ///     var k8s_coreStack = new Spacelift.Stack("k8s-coreStack");
    /// 
    ///     // ...
    ///     var credentials = new Spacelift.EnvironmentVariable("credentials");
    /// 
    ///     // ...
    ///     var k8s_coreStackDestructor = new Spacelift.StackDestructor("k8s-coreStackDestructor", new()
    ///     {
    ///         StackId = k8s_coreStack.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             credentials,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/stackDestructor:StackDestructor")]
    public partial class StackDestructor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set to true, destruction won't delete the stack
        /// </summary>
        [Output("deactivated")]
        public Output<bool?> Deactivated { get; private set; } = null!;

        /// <summary>
        /// ID of the stack to delete and destroy on destruction
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a StackDestructor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackDestructor(string name, StackDestructorArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/stackDestructor:StackDestructor", name, args ?? new StackDestructorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StackDestructor(string name, Input<string> id, StackDestructorState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/stackDestructor:StackDestructor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StackDestructor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackDestructor Get(string name, Input<string> id, StackDestructorState? state = null, CustomResourceOptions? options = null)
        {
            return new StackDestructor(name, id, state, options);
        }
    }

    public sealed class StackDestructorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, destruction won't delete the stack
        /// </summary>
        [Input("deactivated")]
        public Input<bool>? Deactivated { get; set; }

        /// <summary>
        /// ID of the stack to delete and destroy on destruction
        /// </summary>
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public StackDestructorArgs()
        {
        }
        public static new StackDestructorArgs Empty => new StackDestructorArgs();
    }

    public sealed class StackDestructorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, destruction won't delete the stack
        /// </summary>
        [Input("deactivated")]
        public Input<bool>? Deactivated { get; set; }

        /// <summary>
        /// ID of the stack to delete and destroy on destruction
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public StackDestructorState()
        {
        }
        public static new StackDestructorState Empty => new StackDestructorState();
    }
}
