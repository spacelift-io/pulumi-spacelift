// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetAwsIntegrationAttachment
    {
        /// <summary>
        /// `spacelift.AwsIntegrationAttachment` represents the attachment between a reusable AWS integration and a single stack or module.
        /// </summary>
        public static Task<GetAwsIntegrationAttachmentResult> InvokeAsync(GetAwsIntegrationAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAwsIntegrationAttachmentResult>("spacelift:index/getAwsIntegrationAttachment:getAwsIntegrationAttachment", args ?? new GetAwsIntegrationAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.AwsIntegrationAttachment` represents the attachment between a reusable AWS integration and a single stack or module.
        /// </summary>
        public static Output<GetAwsIntegrationAttachmentResult> Invoke(GetAwsIntegrationAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAwsIntegrationAttachmentResult>("spacelift:index/getAwsIntegrationAttachment:getAwsIntegrationAttachment", args ?? new GetAwsIntegrationAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsIntegrationAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("integrationId", required: true)]
        public string IntegrationId { get; set; } = null!;

        [Input("moduleId")]
        public string? ModuleId { get; set; }

        [Input("stackId")]
        public string? StackId { get; set; }

        public GetAwsIntegrationAttachmentArgs()
        {
        }
    }

    public sealed class GetAwsIntegrationAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public GetAwsIntegrationAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAwsIntegrationAttachmentResult
    {
        public readonly string AttachmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IntegrationId;
        public readonly string? ModuleId;
        public readonly bool Read;
        public readonly string? StackId;
        public readonly bool Write;

        [OutputConstructor]
        private GetAwsIntegrationAttachmentResult(
            string attachmentId,

            string id,

            string integrationId,

            string? moduleId,

            bool read,

            string? stackId,

            bool write)
        {
            AttachmentId = attachmentId;
            Id = id;
            IntegrationId = integrationId;
            ModuleId = moduleId;
            Read = read;
            StackId = stackId;
            Write = write;
        }
    }
}