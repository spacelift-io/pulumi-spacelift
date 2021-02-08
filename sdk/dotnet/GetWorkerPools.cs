// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetWorkerPools
    {
        public static Task<GetWorkerPoolsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkerPoolsResult>("spacelift:index/getWorkerPools:getWorkerPools", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetWorkerPoolsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetWorkerPoolsWorkerPoolResult> WorkerPools;

        [OutputConstructor]
        private GetWorkerPoolsResult(
            string id,

            ImmutableArray<Outputs.GetWorkerPoolsWorkerPoolResult> workerPools)
        {
            Id = id;
            WorkerPools = workerPools;
        }
    }
}
