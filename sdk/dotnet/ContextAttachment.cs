// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    [SpaceliftResourceType("spacelift:index/contextAttachment:ContextAttachment")]
    public partial class ContextAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the context to attach
        /// </summary>
        [Output("contextId")]
        public Output<string> ContextId { get; private set; } = null!;

        /// <summary>
        /// ID of the module to attach the context to
        /// </summary>
        [Output("moduleId")]
        public Output<string?> ModuleId { get; private set; } = null!;

        /// <summary>
        /// Priority of the context attachment, used in case of conflicts
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// ID of the stack to attach the context to
        /// </summary>
        [Output("stackId")]
        public Output<string?> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a ContextAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContextAttachment(string name, ContextAttachmentArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/contextAttachment:ContextAttachment", name, args ?? new ContextAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContextAttachment(string name, Input<string> id, ContextAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/contextAttachment:ContextAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContextAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContextAttachment Get(string name, Input<string> id, ContextAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ContextAttachment(name, id, state, options);
        }
    }

    public sealed class ContextAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the context to attach
        /// </summary>
        [Input("contextId", required: true)]
        public Input<string> ContextId { get; set; } = null!;

        /// <summary>
        /// ID of the module to attach the context to
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Priority of the context attachment, used in case of conflicts
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// ID of the stack to attach the context to
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public ContextAttachmentArgs()
        {
        }
    }

    public sealed class ContextAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the context to attach
        /// </summary>
        [Input("contextId")]
        public Input<string>? ContextId { get; set; }

        /// <summary>
        /// ID of the module to attach the context to
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Priority of the context attachment, used in case of conflicts
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// ID of the stack to attach the context to
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public ContextAttachmentState()
        {
        }
    }
}
