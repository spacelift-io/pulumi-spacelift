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
    /// `spacelift.NamedWebhook` represents a named webhook endpoint used for creating webhookswhich are referred to in Notification policies to route messages.
    /// </summary>
    [SpaceliftResourceType("spacelift:index/namedWebhook:NamedWebhook")]
    public partial class NamedWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// enables or disables sending webhooks.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// endpoint to send the requests to
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// labels for the webhook to use when referring in policies or filtering them
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// the name for the webhook which will also be used to generate the id
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
        /// </summary>
        [Output("secret")]
        public Output<string?> Secret { get; private set; } = null!;

        /// <summary>
        /// ID of the space the webhook is in
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a NamedWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NamedWebhook(string name, NamedWebhookArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/namedWebhook:NamedWebhook", name, args ?? new NamedWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NamedWebhook(string name, Input<string> id, NamedWebhookState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/namedWebhook:NamedWebhook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://downloads.spacelift.io/pulumi-plugins",
                AdditionalSecretOutputs =
                {
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NamedWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NamedWebhook Get(string name, Input<string> id, NamedWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new NamedWebhook(name, id, state, options);
        }
    }

    public sealed class NamedWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// enables or disables sending webhooks.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// endpoint to send the requests to
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// labels for the webhook to use when referring in policies or filtering them
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// the name for the webhook which will also be used to generate the id
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// ID of the space the webhook is in
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        public NamedWebhookArgs()
        {
        }
        public static new NamedWebhookArgs Empty => new NamedWebhookArgs();
    }

    public sealed class NamedWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// enables or disables sending webhooks.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// endpoint to send the requests to
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// labels for the webhook to use when referring in policies or filtering them
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// the name for the webhook which will also be used to generate the id
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// ID of the space the webhook is in
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        public NamedWebhookState()
        {
        }
        public static new NamedWebhookState Empty => new NamedWebhookState();
    }
}