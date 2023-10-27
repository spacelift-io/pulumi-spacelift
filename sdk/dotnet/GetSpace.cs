// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetSpace
    {
        /// <summary>
        /// `spacelift.Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
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
        ///     var space = Spacelift.GetSpace.Invoke(new()
        ///     {
        ///         SpaceId = spacelift_space.Space.Id,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["spaceDescription"] = space.Apply(getSpaceResult =&gt; getSpaceResult.Description),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSpaceResult> InvokeAsync(GetSpaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSpaceResult>("spacelift:index/getSpace:getSpace", args ?? new GetSpaceArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
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
        ///     var space = Spacelift.GetSpace.Invoke(new()
        ///     {
        ///         SpaceId = spacelift_space.Space.Id,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["spaceDescription"] = space.Apply(getSpaceResult =&gt; getSpaceResult.Description),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSpaceResult> Invoke(GetSpaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSpaceResult>("spacelift:index/getSpace:getSpace", args ?? new GetSpaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSpaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// immutable ID (slug) of the space
        /// </summary>
        [Input("spaceId", required: true)]
        public string SpaceId { get; set; } = null!;

        public GetSpaceArgs()
        {
        }
        public static new GetSpaceArgs Empty => new GetSpaceArgs();
    }

    public sealed class GetSpaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// immutable ID (slug) of the space
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        public GetSpaceInvokeArgs()
        {
        }
        public static new GetSpaceInvokeArgs Empty => new GetSpaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSpaceResult
    {
        /// <summary>
        /// free-form space description for users
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// indication whether access to this space inherits read access to entities from the parent space
        /// </summary>
        public readonly bool InheritEntities;
        /// <summary>
        /// list of labels describing a space
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// name of the space
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// immutable ID (slug) of parent space
        /// </summary>
        public readonly string ParentSpaceId;
        /// <summary>
        /// immutable ID (slug) of the space
        /// </summary>
        public readonly string SpaceId;

        [OutputConstructor]
        private GetSpaceResult(
            string description,

            string id,

            bool inheritEntities,

            ImmutableArray<string> labels,

            string name,

            string parentSpaceId,

            string spaceId)
        {
            Description = description;
            Id = id;
            InheritEntities = inheritEntities;
            Labels = labels;
            Name = name;
            ParentSpaceId = parentSpaceId;
            SpaceId = spaceId;
        }
    }
}
