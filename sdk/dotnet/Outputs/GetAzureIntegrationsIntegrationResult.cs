// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Outputs
{

    [OutputType]
    public sealed class GetAzureIntegrationsIntegrationResult
    {
        public readonly bool AdminConsentProvided;
        public readonly string AdminConsentUrl;
        public readonly string ApplicationId;
        public readonly string DefaultSubscriptionId;
        public readonly string DisplayName;
        public readonly string IntegrationId;
        public readonly ImmutableArray<string> Labels;
        public readonly string Name;
        public readonly string SpaceId;
        public readonly string TenantId;

        [OutputConstructor]
        private GetAzureIntegrationsIntegrationResult(
            bool adminConsentProvided,

            string adminConsentUrl,

            string applicationId,

            string defaultSubscriptionId,

            string displayName,

            string integrationId,

            ImmutableArray<string> labels,

            string name,

            string spaceId,

            string tenantId)
        {
            AdminConsentProvided = adminConsentProvided;
            AdminConsentUrl = adminConsentUrl;
            ApplicationId = applicationId;
            DefaultSubscriptionId = defaultSubscriptionId;
            DisplayName = displayName;
            IntegrationId = integrationId;
            Labels = labels;
            Name = name;
            SpaceId = spaceId;
            TenantId = tenantId;
        }
    }
}
