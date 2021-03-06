// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetContext
    {
        public static Task<GetContextResult> InvokeAsync(GetContextArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContextResult>("spacelift:index/getContext:getContext", args ?? new GetContextArgs(), options.WithVersion());
    }


    public sealed class GetContextArgs : Pulumi.InvokeArgs
    {
        [Input("contextId", required: true)]
        public string ContextId { get; set; } = null!;

        public GetContextArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContextResult
    {
        public readonly string ContextId;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetContextResult(
            string contextId,

            string description,

            string id,

            string name)
        {
            ContextId = contextId;
            Description = description;
            Id = id;
            Name = name;
        }
    }
}
