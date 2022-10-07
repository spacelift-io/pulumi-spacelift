// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetContextAttachment
    {
        /// <summary>
        /// `spacelift.ContextAttachment` represents a Spacelift attachment of a single context to a single stack or module, with a predefined priority.
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
        ///         var apps_k8s_ie = Output.Create(Spacelift.GetContextAttachment.InvokeAsync(new Spacelift.GetContextAttachmentArgs
        ///         {
        ///             ContextId = "prod-k8s-ie",
        ///             StackId = "apps-cluster",
        ///         }));
        ///         var kafka_k8s_ie = Output.Create(Spacelift.GetContextAttachment.InvokeAsync(new Spacelift.GetContextAttachmentArgs
        ///         {
        ///             ContextId = "prod-k8s-ie",
        ///             ModuleId = "terraform-aws-kafka",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetContextAttachmentResult> InvokeAsync(GetContextAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContextAttachmentResult>("spacelift:index/getContextAttachment:getContextAttachment", args ?? new GetContextAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.ContextAttachment` represents a Spacelift attachment of a single context to a single stack or module, with a predefined priority.
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
        ///         var apps_k8s_ie = Output.Create(Spacelift.GetContextAttachment.InvokeAsync(new Spacelift.GetContextAttachmentArgs
        ///         {
        ///             ContextId = "prod-k8s-ie",
        ///             StackId = "apps-cluster",
        ///         }));
        ///         var kafka_k8s_ie = Output.Create(Spacelift.GetContextAttachment.InvokeAsync(new Spacelift.GetContextAttachmentArgs
        ///         {
        ///             ContextId = "prod-k8s-ie",
        ///             ModuleId = "terraform-aws-kafka",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetContextAttachmentResult> Invoke(GetContextAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetContextAttachmentResult>("spacelift:index/getContextAttachment:getContextAttachment", args ?? new GetContextAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContextAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("contextId", required: true)]
        public string ContextId { get; set; } = null!;

        [Input("moduleId")]
        public string? ModuleId { get; set; }

        [Input("stackId")]
        public string? StackId { get; set; }

        public GetContextAttachmentArgs()
        {
        }
    }

    public sealed class GetContextAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("contextId", required: true)]
        public Input<string> ContextId { get; set; } = null!;

        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public GetContextAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContextAttachmentResult
    {
        public readonly string ContextId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ModuleId;
        public readonly int Priority;
        public readonly string? StackId;

        [OutputConstructor]
        private GetContextAttachmentResult(
            string contextId,

            string id,

            string? moduleId,

            int priority,

            string? stackId)
        {
            ContextId = contextId;
            Id = id;
            ModuleId = moduleId;
            Priority = priority;
            StackId = stackId;
        }
    }
}
