// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetAzureDevopsIntegration
    {
        /// <summary>
        /// `spacelift.getAzureDevopsIntegration` returns details about Azure DevOps integration
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
        ///     var azureDevopsIntegration = Spacelift.GetAzureDevopsIntegration.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAzureDevopsIntegrationResult> InvokeAsync(GetAzureDevopsIntegrationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAzureDevopsIntegrationResult>("spacelift:index/getAzureDevopsIntegration:getAzureDevopsIntegration", args ?? new GetAzureDevopsIntegrationArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.getAzureDevopsIntegration` returns details about Azure DevOps integration
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
        ///     var azureDevopsIntegration = Spacelift.GetAzureDevopsIntegration.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAzureDevopsIntegrationResult> Invoke(GetAzureDevopsIntegrationInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureDevopsIntegrationResult>("spacelift:index/getAzureDevopsIntegration:getAzureDevopsIntegration", args ?? new GetAzureDevopsIntegrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAzureDevopsIntegrationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Azure DevOps integration id. If not provided, the default integration will be returned
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        public GetAzureDevopsIntegrationArgs()
        {
        }
        public static new GetAzureDevopsIntegrationArgs Empty => new GetAzureDevopsIntegrationArgs();
    }

    public sealed class GetAzureDevopsIntegrationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Azure DevOps integration id. If not provided, the default integration will be returned
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetAzureDevopsIntegrationInvokeArgs()
        {
        }
        public static new GetAzureDevopsIntegrationInvokeArgs Empty => new GetAzureDevopsIntegrationInvokeArgs();
    }


    [OutputType]
    public sealed class GetAzureDevopsIntegrationResult
    {
        /// <summary>
        /// Azure DevOps integration description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Azure DevOps integration id. If not provided, the default integration will be returned
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Azure DevOps integration is default
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// Azure DevOps integration labels
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// Azure DevOps integration name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure DevOps integration organization url
        /// </summary>
        public readonly string OrganizationUrl;
        /// <summary>
        /// Azure DevOps integration space id
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// Azure DevOps integration webhook password
        /// </summary>
        public readonly string WebhookPassword;
        /// <summary>
        /// Azure DevOps integration webhook url
        /// </summary>
        public readonly string WebhookUrl;

        [OutputConstructor]
        private GetAzureDevopsIntegrationResult(
            string description,

            string? id,

            bool isDefault,

            ImmutableArray<string> labels,

            string name,

            string organizationUrl,

            string spaceId,

            string webhookPassword,

            string webhookUrl)
        {
            Description = description;
            Id = id;
            IsDefault = isDefault;
            Labels = labels;
            Name = name;
            OrganizationUrl = organizationUrl;
            SpaceId = spaceId;
            WebhookPassword = webhookPassword;
            WebhookUrl = webhookUrl;
        }
    }
}
