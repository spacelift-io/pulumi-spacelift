// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetAwsIntegrationAttachmentExternalId
    {
        /// <summary>
        /// `spacelift.getAwsIntegrationAttachmentExternalId` is used to generate the external ID that would be used for role assumption when an AWS integration is attached to a stack or module.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myStack = Output.Create(Spacelift.GetAwsIntegrationAttachmentExternalId.InvokeAsync(new Spacelift.GetAwsIntegrationAttachmentExternalIdArgs
        ///         {
        ///             IntegrationId = spacelift_aws_integration.This.Id,
        ///             StackId = "my-stack-id",
        ///             Read = true,
        ///             Write = true,
        ///         }));
        ///         var myModule = Output.Create(Spacelift.GetAwsIntegrationAttachmentExternalId.InvokeAsync(new Spacelift.GetAwsIntegrationAttachmentExternalIdArgs
        ///         {
        ///             IntegrationId = spacelift_aws_integration.This.Id,
        ///             ModuleId = "my-module-id",
        ///             Read = true,
        ///             Write = true,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAwsIntegrationAttachmentExternalIdResult> InvokeAsync(GetAwsIntegrationAttachmentExternalIdArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAwsIntegrationAttachmentExternalIdResult>("spacelift:index/getAwsIntegrationAttachmentExternalId:getAwsIntegrationAttachmentExternalId", args ?? new GetAwsIntegrationAttachmentExternalIdArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.getAwsIntegrationAttachmentExternalId` is used to generate the external ID that would be used for role assumption when an AWS integration is attached to a stack or module.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myStack = Output.Create(Spacelift.GetAwsIntegrationAttachmentExternalId.InvokeAsync(new Spacelift.GetAwsIntegrationAttachmentExternalIdArgs
        ///         {
        ///             IntegrationId = spacelift_aws_integration.This.Id,
        ///             StackId = "my-stack-id",
        ///             Read = true,
        ///             Write = true,
        ///         }));
        ///         var myModule = Output.Create(Spacelift.GetAwsIntegrationAttachmentExternalId.InvokeAsync(new Spacelift.GetAwsIntegrationAttachmentExternalIdArgs
        ///         {
        ///             IntegrationId = spacelift_aws_integration.This.Id,
        ///             ModuleId = "my-module-id",
        ///             Read = true,
        ///             Write = true,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAwsIntegrationAttachmentExternalIdResult> Invoke(GetAwsIntegrationAttachmentExternalIdInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAwsIntegrationAttachmentExternalIdResult>("spacelift:index/getAwsIntegrationAttachmentExternalId:getAwsIntegrationAttachmentExternalId", args ?? new GetAwsIntegrationAttachmentExternalIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsIntegrationAttachmentExternalIdArgs : Pulumi.InvokeArgs
    {
        [Input("integrationId", required: true)]
        public string IntegrationId { get; set; } = null!;

        [Input("moduleId")]
        public string? ModuleId { get; set; }

        [Input("read")]
        public bool? Read { get; set; }

        [Input("stackId")]
        public string? StackId { get; set; }

        [Input("write")]
        public bool? Write { get; set; }

        public GetAwsIntegrationAttachmentExternalIdArgs()
        {
        }
    }

    public sealed class GetAwsIntegrationAttachmentExternalIdInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        [Input("read")]
        public Input<bool>? Read { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("write")]
        public Input<bool>? Write { get; set; }

        public GetAwsIntegrationAttachmentExternalIdInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAwsIntegrationAttachmentExternalIdResult
    {
        public readonly string AssumeRolePolicyStatement;
        public readonly string ExternalId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IntegrationId;
        public readonly string? ModuleId;
        public readonly bool? Read;
        public readonly string? StackId;
        public readonly bool? Write;

        [OutputConstructor]
        private GetAwsIntegrationAttachmentExternalIdResult(
            string assumeRolePolicyStatement,

            string externalId,

            string id,

            string integrationId,

            string? moduleId,

            bool? read,

            string? stackId,

            bool? write)
        {
            AssumeRolePolicyStatement = assumeRolePolicyStatement;
            ExternalId = externalId;
            Id = id;
            IntegrationId = integrationId;
            ModuleId = moduleId;
            Read = read;
            StackId = stackId;
            Write = write;
        }
    }
}