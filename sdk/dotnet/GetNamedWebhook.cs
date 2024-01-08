// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetNamedWebhook
    {
        /// <summary>
        /// `spacelift.NamedWebhook` represents a named webhook endpoint used for creating webhookswhich are referred to in Notification policies to route messages.
        /// </summary>
        public static Task<GetNamedWebhookResult> InvokeAsync(GetNamedWebhookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNamedWebhookResult>("spacelift:index/getNamedWebhook:getNamedWebhook", args ?? new GetNamedWebhookArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.NamedWebhook` represents a named webhook endpoint used for creating webhookswhich are referred to in Notification policies to route messages.
        /// </summary>
        public static Output<GetNamedWebhookResult> Invoke(GetNamedWebhookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNamedWebhookResult>("spacelift:index/getNamedWebhook:getNamedWebhook", args ?? new GetNamedWebhookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamedWebhookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public string WebhookId { get; set; } = null!;

        public GetNamedWebhookArgs()
        {
        }
        public static new GetNamedWebhookArgs Empty => new GetNamedWebhookArgs();
    }

    public sealed class GetNamedWebhookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public Input<string> WebhookId { get; set; } = null!;

        public GetNamedWebhookInvokeArgs()
        {
        }
        public static new GetNamedWebhookInvokeArgs Empty => new GetNamedWebhookInvokeArgs();
    }


    [OutputType]
    public sealed class GetNamedWebhookResult
    {
        /// <summary>
        /// enables or disables sending webhooks.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// endpoint to send the requests to
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// labels for the webhook to use when referring in policies or filtering them
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// the name for the webhook which will also be used to generate the id
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// secret used to sign each request so you're able to verify that the request comes from us.
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// secret header keys which are currently set for this webhook
        /// </summary>
        public readonly ImmutableArray<string> SecretHeaderKeys;
        /// <summary>
        /// ID of the space the webhook is in
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// ID of the webhook
        /// </summary>
        public readonly string WebhookId;

        [OutputConstructor]
        private GetNamedWebhookResult(
            bool enabled,

            string endpoint,

            string id,

            ImmutableArray<string> labels,

            string name,

            string secret,

            ImmutableArray<string> secretHeaderKeys,

            string spaceId,

            string webhookId)
        {
            Enabled = enabled;
            Endpoint = endpoint;
            Id = id;
            Labels = labels;
            Name = name;
            Secret = secret;
            SecretHeaderKeys = secretHeaderKeys;
            SpaceId = spaceId;
            WebhookId = webhookId;
        }
    }
}