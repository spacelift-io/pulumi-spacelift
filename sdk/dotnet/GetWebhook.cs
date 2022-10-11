// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetWebhook
    {
        /// <summary>
        /// `spacelift.Webook` represents a webhook endpoint to which Spacelift sends the POST request about run state changes.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var webhook = Spacelift.GetWebhook.Invoke(new()
        ///     {
        ///         WebhookId = spacelift_webhook.Webhook.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWebhookResult> InvokeAsync(GetWebhookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebhookResult>("spacelift:index/getWebhook:getWebhook", args ?? new GetWebhookArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.Webook` represents a webhook endpoint to which Spacelift sends the POST request about run state changes.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var webhook = Spacelift.GetWebhook.Invoke(new()
        ///     {
        ///         WebhookId = spacelift_webhook.Webhook.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWebhookResult> Invoke(GetWebhookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhookResult>("spacelift:index/getWebhook:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the stack which triggers the webhooks
        /// </summary>
        [Input("moduleId")]
        public string? ModuleId { get; set; }

        /// <summary>
        /// ID of the stack which triggers the webhooks
        /// </summary>
        [Input("stackId")]
        public string? StackId { get; set; }

        /// <summary>
        /// ID of the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public string WebhookId { get; set; } = null!;

        public GetWebhookArgs()
        {
        }
        public static new GetWebhookArgs Empty => new GetWebhookArgs();
    }

    public sealed class GetWebhookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the stack which triggers the webhooks
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// ID of the stack which triggers the webhooks
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// ID of the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public Input<string> WebhookId { get; set; } = null!;

        public GetWebhookInvokeArgs()
        {
        }
        public static new GetWebhookInvokeArgs Empty => new GetWebhookInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebhookResult
    {
        /// <summary>
        /// enables or disables sending webhooks
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// endpoint to send the POST request to
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the stack which triggers the webhooks
        /// </summary>
        public readonly string? ModuleId;
        /// <summary>
        /// secret used to sign each POST request so you're able to verify that the request comes from us
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// ID of the stack which triggers the webhooks
        /// </summary>
        public readonly string? StackId;
        /// <summary>
        /// ID of the webhook
        /// </summary>
        public readonly string WebhookId;

        [OutputConstructor]
        private GetWebhookResult(
            bool enabled,

            string endpoint,

            string id,

            string? moduleId,

            string secret,

            string? stackId,

            string webhookId)
        {
            Enabled = enabled;
            Endpoint = endpoint;
            Id = id;
            ModuleId = moduleId;
            Secret = secret;
            StackId = stackId;
            WebhookId = webhookId;
        }
    }
}
