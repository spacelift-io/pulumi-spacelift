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
    public sealed class GetVcsAgentPoolsVcsAgentPoolResult
    {
        public readonly string Description;
        public readonly string Name;
        public readonly string VcsAgentPoolId;

        [OutputConstructor]
        private GetVcsAgentPoolsVcsAgentPoolResult(
            string description,

            string name,

            string vcsAgentPoolId)
        {
            Description = description;
            Name = name;
            VcsAgentPoolId = vcsAgentPoolId;
        }
    }
}