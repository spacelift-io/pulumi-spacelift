// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    [SpaceliftResourceType("spacelift:index/mountedFile:MountedFile")]
    public partial class MountedFile : Pulumi.CustomResource
    {
        /// <summary>
        /// SHA-256 checksum of the value
        /// </summary>
        [Output("checksum")]
        public Output<string> Checksum { get; private set; } = null!;

        /// <summary>
        /// Content of the mounted file encoded using Base-64
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// ID of the context on which the mounted file is defined
        /// </summary>
        [Output("contextId")]
        public Output<string?> ContextId { get; private set; } = null!;

        /// <summary>
        /// ID of the module on which the mounted file is defined
        /// </summary>
        [Output("moduleId")]
        public Output<string?> ModuleId { get; private set; } = null!;

        /// <summary>
        /// Relative path to the mounted file, without the /mnt/workspace/ prefix
        /// </summary>
        [Output("relativePath")]
        public Output<string> RelativePath { get; private set; } = null!;

        /// <summary>
        /// ID of the stack on which the mounted file is defined
        /// </summary>
        [Output("stackId")]
        public Output<string?> StackId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the content can be read back outside a Run
        /// </summary>
        [Output("writeOnly")]
        public Output<bool?> WriteOnly { get; private set; } = null!;


        /// <summary>
        /// Create a MountedFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MountedFile(string name, MountedFileArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/mountedFile:MountedFile", name, args ?? new MountedFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MountedFile(string name, Input<string> id, MountedFileState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/mountedFile:MountedFile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MountedFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MountedFile Get(string name, Input<string> id, MountedFileState? state = null, CustomResourceOptions? options = null)
        {
            return new MountedFile(name, id, state, options);
        }
    }

    public sealed class MountedFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content of the mounted file encoded using Base-64
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// ID of the context on which the mounted file is defined
        /// </summary>
        [Input("contextId")]
        public Input<string>? ContextId { get; set; }

        /// <summary>
        /// ID of the module on which the mounted file is defined
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Relative path to the mounted file, without the /mnt/workspace/ prefix
        /// </summary>
        [Input("relativePath", required: true)]
        public Input<string> RelativePath { get; set; } = null!;

        /// <summary>
        /// ID of the stack on which the mounted file is defined
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// Indicates whether the content can be read back outside a Run
        /// </summary>
        [Input("writeOnly")]
        public Input<bool>? WriteOnly { get; set; }

        public MountedFileArgs()
        {
        }
    }

    public sealed class MountedFileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SHA-256 checksum of the value
        /// </summary>
        [Input("checksum")]
        public Input<string>? Checksum { get; set; }

        /// <summary>
        /// Content of the mounted file encoded using Base-64
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// ID of the context on which the mounted file is defined
        /// </summary>
        [Input("contextId")]
        public Input<string>? ContextId { get; set; }

        /// <summary>
        /// ID of the module on which the mounted file is defined
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Relative path to the mounted file, without the /mnt/workspace/ prefix
        /// </summary>
        [Input("relativePath")]
        public Input<string>? RelativePath { get; set; }

        /// <summary>
        /// ID of the stack on which the mounted file is defined
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// Indicates whether the content can be read back outside a Run
        /// </summary>
        [Input("writeOnly")]
        public Input<bool>? WriteOnly { get; set; }

        public MountedFileState()
        {
        }
    }
}
