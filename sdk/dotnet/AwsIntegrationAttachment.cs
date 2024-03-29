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
    /// `spacelift.AwsIntegrationAttachment` represents the attachment between a reusable AWS integration and a single stack or module.
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
    ///     // For a stack
    ///     var thisAwsIntegrationAttachment = new Spacelift.AwsIntegrationAttachment("thisAwsIntegrationAttachment", new()
    ///     {
    ///         IntegrationId = spacelift_aws_integration.This.Id,
    ///         StackId = "my-stack-id",
    ///         Read = true,
    ///         Write = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_iam_role.This,
    ///         },
    ///     });
    /// 
    ///     // For a module
    ///     var thisIndex_awsIntegrationAttachmentAwsIntegrationAttachment = new Spacelift.AwsIntegrationAttachment("thisIndex/awsIntegrationAttachmentAwsIntegrationAttachment", new()
    ///     {
    ///         IntegrationId = spacelift_aws_integration.This.Id,
    ///         ModuleId = "my-module-id",
    ///         Read = true,
    ///         Write = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_iam_role.This,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment read_write_my_stack $INTEGRATION_ID/$PROJECT_ID
    /// ```
    /// </summary>
    [SpaceliftResourceType("spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment")]
    public partial class AwsIntegrationAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Internal ID of the attachment entity
        /// </summary>
        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// ID of the integration to attach
        /// </summary>
        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// ID of the module to attach the integration to
        /// </summary>
        [Output("moduleId")]
        public Output<string?> ModuleId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this attachment is used for read operations. Defaults to `true`.
        /// </summary>
        [Output("read")]
        public Output<bool?> Read { get; private set; } = null!;

        /// <summary>
        /// ID of the stack to attach the integration to
        /// </summary>
        [Output("stackId")]
        public Output<string?> StackId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this attachment is used for write operations. Defaults to `true`.
        /// </summary>
        [Output("write")]
        public Output<bool?> Write { get; private set; } = null!;


        /// <summary>
        /// Create a AwsIntegrationAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsIntegrationAttachment(string name, AwsIntegrationAttachmentArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment", name, args ?? new AwsIntegrationAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsIntegrationAttachment(string name, Input<string> id, AwsIntegrationAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AwsIntegrationAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsIntegrationAttachment Get(string name, Input<string> id, AwsIntegrationAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsIntegrationAttachment(name, id, state, options);
        }
    }

    public sealed class AwsIntegrationAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the integration to attach
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// ID of the module to attach the integration to
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Indicates whether this attachment is used for read operations. Defaults to `true`.
        /// </summary>
        [Input("read")]
        public Input<bool>? Read { get; set; }

        /// <summary>
        /// ID of the stack to attach the integration to
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// Indicates whether this attachment is used for write operations. Defaults to `true`.
        /// </summary>
        [Input("write")]
        public Input<bool>? Write { get; set; }

        public AwsIntegrationAttachmentArgs()
        {
        }
        public static new AwsIntegrationAttachmentArgs Empty => new AwsIntegrationAttachmentArgs();
    }

    public sealed class AwsIntegrationAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internal ID of the attachment entity
        /// </summary>
        [Input("attachmentId")]
        public Input<string>? AttachmentId { get; set; }

        /// <summary>
        /// ID of the integration to attach
        /// </summary>
        [Input("integrationId")]
        public Input<string>? IntegrationId { get; set; }

        /// <summary>
        /// ID of the module to attach the integration to
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// Indicates whether this attachment is used for read operations. Defaults to `true`.
        /// </summary>
        [Input("read")]
        public Input<bool>? Read { get; set; }

        /// <summary>
        /// ID of the stack to attach the integration to
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// Indicates whether this attachment is used for write operations. Defaults to `true`.
        /// </summary>
        [Input("write")]
        public Input<bool>? Write { get; set; }

        public AwsIntegrationAttachmentState()
        {
        }
        public static new AwsIntegrationAttachmentState Empty => new AwsIntegrationAttachmentState();
    }
}
