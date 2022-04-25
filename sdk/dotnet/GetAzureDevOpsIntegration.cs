// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetAzureDevOpsIntegration
    {
        /// <summary>
        /// `spacelift.getAzureDevOpsIntegration` returns details about Azure DevOps integration
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
        ///         var azureDevopsIntegration = Output.Create(Spacelift.GetAzureDevOpsIntegration.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// &lt;!-- schema generated by tfplugindocs --&gt;
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Schema
        /// 
        /// ### Optional
        /// 
        /// - **id** (String) The ID of this resource.
        /// 
        /// ### Read-Only
        /// 
        /// - **organization_url** (String) Azure DevOps integration organization url
        /// - **webhook_password** (String) Azure DevOps integration webhook password
        /// </summary>
        public static Task<GetAzureDevOpsIntegrationResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAzureDevOpsIntegrationResult>("spacelift:index/getAzureDevOpsIntegration:getAzureDevOpsIntegration", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetAzureDevOpsIntegrationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationUrl;
        public readonly string WebhookPassword;

        [OutputConstructor]
        private GetAzureDevOpsIntegrationResult(
            string id,

            string organizationUrl,

            string webhookPassword)
        {
            Id = id;
            OrganizationUrl = organizationUrl;
            WebhookPassword = webhookPassword;
        }
    }
}