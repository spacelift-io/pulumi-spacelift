// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetVcsAgentPool
    {
        /// <summary>
        /// `spacelift.VcsAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations
        /// </summary>
        public static Task<GetVcsAgentPoolResult> InvokeAsync(GetVcsAgentPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVcsAgentPoolResult>("spacelift:index/getVcsAgentPool:getVcsAgentPool", args ?? new GetVcsAgentPoolArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.VcsAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations
        /// </summary>
        public static Output<GetVcsAgentPoolResult> Invoke(GetVcsAgentPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVcsAgentPoolResult>("spacelift:index/getVcsAgentPool:getVcsAgentPool", args ?? new GetVcsAgentPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVcsAgentPoolArgs : Pulumi.InvokeArgs
    {
        [Input("vcsAgentPoolId", required: true)]
        public string VcsAgentPoolId { get; set; } = null!;

        public GetVcsAgentPoolArgs()
        {
        }
    }

    public sealed class GetVcsAgentPoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("vcsAgentPoolId", required: true)]
        public Input<string> VcsAgentPoolId { get; set; } = null!;

        public GetVcsAgentPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVcsAgentPoolResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string VcsAgentPoolId;

        [OutputConstructor]
        private GetVcsAgentPoolResult(
            string description,

            string id,

            string name,

            string vcsAgentPoolId)
        {
            Description = description;
            Id = id;
            Name = name;
            VcsAgentPoolId = vcsAgentPoolId;
        }
    }
}
