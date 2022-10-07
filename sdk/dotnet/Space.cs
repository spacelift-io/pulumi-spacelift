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
    /// `spacelift.Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
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
    ///         var development = new Spacelift.Space("development", new Spacelift.SpaceArgs
    ///         {
    ///             ParentSpaceId = "root",
    ///             Description = "This a child of the root space. It contains all the resources common to the development infrastructure.",
    ///         });
    ///         var development_frontend = new Spacelift.Space("development-frontend", new Spacelift.SpaceArgs
    ///         {
    ///             ParentSpaceId = development.Id,
    ///             InheritEntities = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import spacelift:index/space:Space development $SPACE_ID
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/space:Space")]
    public partial class Space : Pulumi.CustomResource
    {
        /// <summary>
        /// free-form space description for users
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
        /// </summary>
        [Output("inheritEntities")]
        public Output<bool?> InheritEntities { get; private set; } = null!;

        /// <summary>
        /// name of the space
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// immutable ID (slug) of parent space. Defaults to `root`.
        /// </summary>
        [Output("parentSpaceId")]
        public Output<string?> ParentSpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Space resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Space(string name, SpaceArgs? args = null, CustomResourceOptions? options = null)
            : base("spacelift:index/space:Space", name, args ?? new SpaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Space(string name, Input<string> id, SpaceState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/space:Space", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Space resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Space Get(string name, Input<string> id, SpaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Space(name, id, state, options);
        }
    }

    public sealed class SpaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// free-form space description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
        /// </summary>
        [Input("inheritEntities")]
        public Input<bool>? InheritEntities { get; set; }

        /// <summary>
        /// name of the space
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// immutable ID (slug) of parent space. Defaults to `root`.
        /// </summary>
        [Input("parentSpaceId")]
        public Input<string>? ParentSpaceId { get; set; }

        public SpaceArgs()
        {
        }
    }

    public sealed class SpaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// free-form space description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
        /// </summary>
        [Input("inheritEntities")]
        public Input<bool>? InheritEntities { get; set; }

        /// <summary>
        /// name of the space
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// immutable ID (slug) of parent space. Defaults to `root`.
        /// </summary>
        [Input("parentSpaceId")]
        public Input<string>? ParentSpaceId { get; set; }

        public SpaceState()
        {
        }
    }
}
