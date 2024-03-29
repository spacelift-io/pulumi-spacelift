// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetContexts
    {
        /// <summary>
        /// `spacelift.getContexts` represents all the contexts in the Spacelift account visible to the API user.
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
        ///     var contexts = Spacelift.GetContexts.Invoke(new()
        ///     {
        ///         Labels = new[]
        ///         {
        ///             new Spacelift.Inputs.GetContextsLabelInputArgs
        ///             {
        ///                 AnyOfs = new[]
        ///                 {
        ///                     "foo",
        ///                     "bar",
        ///                 },
        ///             },
        ///             new Spacelift.Inputs.GetContextsLabelInputArgs
        ///             {
        ///                 AnyOfs = new[]
        ///                 {
        ///                     "baz",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetContextsResult> InvokeAsync(GetContextsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContextsResult>("spacelift:index/getContexts:getContexts", args ?? new GetContextsArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.getContexts` represents all the contexts in the Spacelift account visible to the API user.
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
        ///     var contexts = Spacelift.GetContexts.Invoke(new()
        ///     {
        ///         Labels = new[]
        ///         {
        ///             new Spacelift.Inputs.GetContextsLabelInputArgs
        ///             {
        ///                 AnyOfs = new[]
        ///                 {
        ///                     "foo",
        ///                     "bar",
        ///                 },
        ///             },
        ///             new Spacelift.Inputs.GetContextsLabelInputArgs
        ///             {
        ///                 AnyOfs = new[]
        ///                 {
        ///                     "baz",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetContextsResult> Invoke(GetContextsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContextsResult>("spacelift:index/getContexts:getContexts", args ?? new GetContextsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContextsArgs : global::Pulumi.InvokeArgs
    {
        [Input("labels")]
        private List<Inputs.GetContextsLabelArgs>? _labels;
        public List<Inputs.GetContextsLabelArgs> Labels
        {
            get => _labels ?? (_labels = new List<Inputs.GetContextsLabelArgs>());
            set => _labels = value;
        }

        public GetContextsArgs()
        {
        }
        public static new GetContextsArgs Empty => new GetContextsArgs();
    }

    public sealed class GetContextsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("labels")]
        private InputList<Inputs.GetContextsLabelInputArgs>? _labels;
        public InputList<Inputs.GetContextsLabelInputArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.GetContextsLabelInputArgs>());
            set => _labels = value;
        }

        public GetContextsInvokeArgs()
        {
        }
        public static new GetContextsInvokeArgs Empty => new GetContextsInvokeArgs();
    }


    [OutputType]
    public sealed class GetContextsResult
    {
        public readonly ImmutableArray<Outputs.GetContextsContextResult> Contexts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Require contexts to have one of the labels
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContextsLabelResult> Labels;

        [OutputConstructor]
        private GetContextsResult(
            ImmutableArray<Outputs.GetContextsContextResult> contexts,

            string id,

            ImmutableArray<Outputs.GetContextsLabelResult> labels)
        {
            Contexts = contexts;
            Id = id;
            Labels = labels;
        }
    }
}
