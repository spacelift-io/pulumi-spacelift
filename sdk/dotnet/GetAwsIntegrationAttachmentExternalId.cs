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
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myStack = Spacelift.GetAwsIntegrationAttachmentExternalId.Invoke(new()
        ///     {
        ///         IntegrationId = spacelift_aws_integration.This.Id,
        ///         StackId = "my-stack-id",
        ///         Read = true,
        ///         Write = true,
        ///     });
        /// 
        ///     var myModule = Spacelift.GetAwsIntegrationAttachmentExternalId.Invoke(new()
        ///     {
        ///         IntegrationId = spacelift_aws_integration.This.Id,
        ///         ModuleId = "my-module-id",
        ///         Read = true,
        ///         Write = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAwsIntegrationAttachmentExternalIdResult> InvokeAsync(GetAwsIntegrationAttachmentExternalIdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAwsIntegrationAttachmentExternalIdResult>("spacelift:index/getAwsIntegrationAttachmentExternalId:getAwsIntegrationAttachmentExternalId", args ?? new GetAwsIntegrationAttachmentExternalIdArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.getAwsIntegrationAttachmentExternalId` is used to generate the external ID that would be used for role assumption when an AWS integration is attached to a stack or module.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myStack = Spacelift.GetAwsIntegrationAttachmentExternalId.Invoke(new()
        ///     {
        ///         IntegrationId = spacelift_aws_integration.This.Id,
        ///         StackId = "my-stack-id",
        ///         Read = true,
        ///         Write = true,
        ///     });
        /// 
        ///     var myModule = Spacelift.GetAwsIntegrationAttachmentExternalId.Invoke(new()
        ///     {
        ///         IntegrationId = spacelift_aws_integration.This.Id,
        ///         ModuleId = "my-module-id",
        ///         Read = true,
        ///         Write = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAwsIntegrationAttachmentExternalIdResult> Invoke(GetAwsIntegrationAttachmentExternalIdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAwsIntegrationAttachmentExternalIdResult>("spacelift:index/getAwsIntegrationAttachmentExternalId:getAwsIntegrationAttachmentExternalId", args ?? new GetAwsIntegrationAttachmentExternalIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsIntegrationAttachmentExternalIdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// immutable ID (slug) of the AWS integration
        /// </summary>
        [Input("integrationId", required: true)]
        public string IntegrationId { get; set; } = null!;

        /// <summary>
        /// immutable ID (slug) of the module
        /// </summary>
        [Input("moduleId")]
        public string? ModuleId { get; set; }

        /// <summary>
        /// whether the integration will be used for read operations
        /// </summary>
        [Input("read")]
        public bool? Read { get; set; }

        /// <summary>
        /// immutable ID (slug) of the stack
        /// </summary>
        [Input("stackId")]
        public string? StackId { get; set; }

        /// <summary>
        /// whether the integration will be used for write operations
        /// </summary>
        [Input("write")]
        public bool? Write { get; set; }

        public GetAwsIntegrationAttachmentExternalIdArgs()
        {
        }
        public static new GetAwsIntegrationAttachmentExternalIdArgs Empty => new GetAwsIntegrationAttachmentExternalIdArgs();
    }

    public sealed class GetAwsIntegrationAttachmentExternalIdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// immutable ID (slug) of the AWS integration
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// immutable ID (slug) of the module
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// whether the integration will be used for read operations
        /// </summary>
        [Input("read")]
        public Input<bool>? Read { get; set; }

        /// <summary>
        /// immutable ID (slug) of the stack
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// whether the integration will be used for write operations
        /// </summary>
        [Input("write")]
        public Input<bool>? Write { get; set; }

        public GetAwsIntegrationAttachmentExternalIdInvokeArgs()
        {
        }
        public static new GetAwsIntegrationAttachmentExternalIdInvokeArgs Empty => new GetAwsIntegrationAttachmentExternalIdInvokeArgs();
    }


    [OutputType]
    public sealed class GetAwsIntegrationAttachmentExternalIdResult
    {
        /// <summary>
        /// An assume role policy statement that can be attached to your role to allow Spacelift to assume it
        /// </summary>
        public readonly string AssumeRolePolicyStatement;
        /// <summary>
        /// The external ID that will be used during role assumption
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// immutable ID (slug) of the AWS integration
        /// </summary>
        public readonly string IntegrationId;
        /// <summary>
        /// immutable ID (slug) of the module
        /// </summary>
        public readonly string? ModuleId;
        /// <summary>
        /// whether the integration will be used for read operations
        /// </summary>
        public readonly bool? Read;
        /// <summary>
        /// immutable ID (slug) of the stack
        /// </summary>
        public readonly string? StackId;
        /// <summary>
        /// whether the integration will be used for write operations
        /// </summary>
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
