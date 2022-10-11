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
        /// <summary>
        /// `spacelift.Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`spacelift.Stack`) or modules (`spacelift.Module`) using a context attachment (`spacelift.ContextAttachment`)`
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var prod_k8s_ie = Spacelift.GetContext.Invoke(new()
        ///     {
        ///         ContextId = "prod-k8s-ie",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetContextResult> InvokeAsync(GetContextArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContextResult>("spacelift:index/getContext:getContext", args ?? new GetContextArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`spacelift.Stack`) or modules (`spacelift.Module`) using a context attachment (`spacelift.ContextAttachment`)`
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var prod_k8s_ie = Spacelift.GetContext.Invoke(new()
        ///     {
        ///         ContextId = "prod-k8s-ie",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetContextResult> Invoke(GetContextInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContextResult>("spacelift:index/getContext:getContext", args ?? new GetContextInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContextArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// immutable ID (slug) of the context
        /// </summary>
        [Input("contextId", required: true)]
        public string ContextId { get; set; } = null!;

        public GetContextArgs()
        {
        }
        public static new GetContextArgs Empty => new GetContextArgs();
    }

    public sealed class GetContextInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// immutable ID (slug) of the context
        /// </summary>
        [Input("contextId", required: true)]
        public Input<string> ContextId { get; set; } = null!;

        public GetContextInvokeArgs()
        {
        }
        public static new GetContextInvokeArgs Empty => new GetContextInvokeArgs();
    }


    [OutputType]
    public sealed class GetContextResult
    {
        /// <summary>
        /// immutable ID (slug) of the context
        /// </summary>
        public readonly string ContextId;
        /// <summary>
        /// free-form context description for users
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// name of the context
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID (slug) of the space the context is in
        /// </summary>
        public readonly string SpaceId;

        [OutputConstructor]
        private GetContextResult(
            string contextId,

            string description,

            string id,

            ImmutableArray<string> labels,

            string name,

            string spaceId)
        {
            ContextId = contextId;
            Description = description;
            Id = id;
            Labels = labels;
            Name = name;
            SpaceId = spaceId;
        }
    }
}
