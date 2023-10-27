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
    public sealed class GetAwsIntegrationsIntegrationResult
    {
        public readonly int DurationSeconds;
        public readonly string ExternalId;
        public readonly bool GenerateCredentialsInWorker;
        public readonly string IntegrationId;
        public readonly ImmutableArray<string> Labels;
        public readonly string Name;
        public readonly string RoleArn;
        public readonly string SpaceId;

        [OutputConstructor]
        private GetAwsIntegrationsIntegrationResult(
            int durationSeconds,

            string externalId,

            bool generateCredentialsInWorker,

            string integrationId,

            ImmutableArray<string> labels,

            string name,

            string roleArn,

            string spaceId)
        {
            DurationSeconds = durationSeconds;
            ExternalId = externalId;
            GenerateCredentialsInWorker = generateCredentialsInWorker;
            IntegrationId = integrationId;
            Labels = labels;
            Name = name;
            RoleArn = roleArn;
            SpaceId = spaceId;
        }
    }
}
