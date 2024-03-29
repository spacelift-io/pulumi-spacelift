// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetBitbucketDatacenterIntegration
    {
        /// <summary>
        /// `spacelift.getBitbucketDatacenterIntegration` returns details about Bitbucket Datacenter integration
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
        ///     var bitbucketDatacenterIntegration = Spacelift.GetBitbucketDatacenterIntegration.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBitbucketDatacenterIntegrationResult> InvokeAsync(GetBitbucketDatacenterIntegrationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBitbucketDatacenterIntegrationResult>("spacelift:index/getBitbucketDatacenterIntegration:getBitbucketDatacenterIntegration", args ?? new GetBitbucketDatacenterIntegrationArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.getBitbucketDatacenterIntegration` returns details about Bitbucket Datacenter integration
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
        ///     var bitbucketDatacenterIntegration = Spacelift.GetBitbucketDatacenterIntegration.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBitbucketDatacenterIntegrationResult> Invoke(GetBitbucketDatacenterIntegrationInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBitbucketDatacenterIntegrationResult>("spacelift:index/getBitbucketDatacenterIntegration:getBitbucketDatacenterIntegration", args ?? new GetBitbucketDatacenterIntegrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBitbucketDatacenterIntegrationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bitbucket Datacenter integration id. If not provided, the default integration will be returned
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        public GetBitbucketDatacenterIntegrationArgs()
        {
        }
        public static new GetBitbucketDatacenterIntegrationArgs Empty => new GetBitbucketDatacenterIntegrationArgs();
    }

    public sealed class GetBitbucketDatacenterIntegrationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bitbucket Datacenter integration id. If not provided, the default integration will be returned
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetBitbucketDatacenterIntegrationInvokeArgs()
        {
        }
        public static new GetBitbucketDatacenterIntegrationInvokeArgs Empty => new GetBitbucketDatacenterIntegrationInvokeArgs();
    }


    [OutputType]
    public sealed class GetBitbucketDatacenterIntegrationResult
    {
        /// <summary>
        /// Bitbucket Datacenter integration api host
        /// </summary>
        public readonly string ApiHost;
        /// <summary>
        /// Bitbucket Datacenter integration description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Bitbucket Datacenter integration id. If not provided, the default integration will be returned
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Bitbucket Datacenter integration is default
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// Bitbucket Datacenter integration labels
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// Bitbucket Datacenter integration name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Bitbucket Datacenter integration space id
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// Bitbucket Datacenter integration user facing host
        /// </summary>
        public readonly string UserFacingHost;
        /// <summary>
        /// Bitbucket Datacenter username
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// Bitbucket Datacenter integration webhook secret
        /// </summary>
        public readonly string WebhookSecret;
        /// <summary>
        /// Bitbucket Datacenter integration webhook URL
        /// </summary>
        public readonly string WebhookUrl;

        [OutputConstructor]
        private GetBitbucketDatacenterIntegrationResult(
            string apiHost,

            string description,

            string? id,

            bool isDefault,

            ImmutableArray<string> labels,

            string name,

            string spaceId,

            string userFacingHost,

            string username,

            string webhookSecret,

            string webhookUrl)
        {
            ApiHost = apiHost;
            Description = description;
            Id = id;
            IsDefault = isDefault;
            Labels = labels;
            Name = name;
            SpaceId = spaceId;
            UserFacingHost = userFacingHost;
            Username = username;
            WebhookSecret = webhookSecret;
            WebhookUrl = webhookUrl;
        }
    }
}
