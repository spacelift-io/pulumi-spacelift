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
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var space = Output.Create(Spacelift.GetSpace.InvokeAsync(new Spacelift.GetSpaceArgs
        ///         {
        ///             SpaceId = spacelift_space.Space.Id,
        ///         }));
        ///         this.SpaceDescription = space.Apply(space =&gt; space.Description);
        ///     }
        /// 
        ///     [Output("spaceDescription")]
        ///     public Output&lt;string&gt; SpaceDescription { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSpaceResult> InvokeAsync(GetSpaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSpaceResult>("spacelift:index/getSpace:getSpace", args ?? new GetSpaceArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
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
        ///         var space = Output.Create(Spacelift.GetSpace.InvokeAsync(new Spacelift.GetSpaceArgs
        ///         {
        ///             SpaceId = spacelift_space.Space.Id,
        ///         }));
        ///         this.SpaceDescription = space.Apply(space =&gt; space.Description);
        ///     }
        /// 
        ///     [Output("spaceDescription")]
        ///     public Output&lt;string&gt; SpaceDescription { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSpaceResult> Invoke(GetSpaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSpaceResult>("spacelift:index/getSpace:getSpace", args ?? new GetSpaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSpaceArgs : Pulumi.InvokeArgs
    {
        [Input("spaceId", required: true)]
        public string SpaceId { get; set; } = null!;

        public GetSpaceArgs()
        {
        }
    }

    public sealed class GetSpaceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        public GetSpaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSpaceResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool InheritEntities;
        public readonly string Name;
        public readonly string ParentSpaceId;
        public readonly string SpaceId;

        [OutputConstructor]
        private GetSpaceResult(
            string description,

            string id,

            bool inheritEntities,

            string name,

            string parentSpaceId,

            string spaceId)
        {
            Description = description;
            Id = id;
            InheritEntities = inheritEntities;
            Name = name;
            ParentSpaceId = parentSpaceId;
            SpaceId = spaceId;
        }
    }
}
