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
    /// `spacelift.Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`spacelift.Stack`) or modules (`spacelift.Module`) using a context attachment (`spacelift.ContextAttachment`)`
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
    ///         var prod_k8s_ie = new Spacelift.Context("prod-k8s-ie", new Spacelift.ContextArgs
    ///         {
    ///             Description = "Configuration details for the compute cluster in 🇮🇪",
    ///             Name = "Production cluster (Ireland)",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// &lt;!-- schema generated by tfplugindocs --&gt;
    /// ## Schema
    /// 
    /// ### Required
    /// 
    /// - **name** (String) Name of the context - should be unique in one account
    /// 
    /// ### Optional
    /// 
    /// - **description** (String) Free-form context description for users
    /// - **id** (String) The ID of this resource.
    /// - **labels** (Set of String)
    /// 
    /// ## Import
    /// 
    /// Import is supported using the following syntax
    /// 
    /// ```sh
    ///  $ pulumi import spacelift:index/context:Context prod-k8s-ie $CONTEXT_ID
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/context:Context")]
    public partial class Context : Pulumi.CustomResource
    {
        /// <summary>
        /// Free-form context description for users
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the context - should be unique in one account
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Context resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Context(string name, ContextArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/context:Context", name, args ?? new ContextArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Context(string name, Input<string> id, ContextState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/context:Context", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Context resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Context Get(string name, Input<string> id, ContextState? state = null, CustomResourceOptions? options = null)
        {
            return new Context(name, id, state, options);
        }
    }

    public sealed class ContextArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Free-form context description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the context - should be unique in one account
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ContextArgs()
        {
        }
    }

    public sealed class ContextState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Free-form context description for users
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the context - should be unique in one account
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ContextState()
        {
        }
    }
}
