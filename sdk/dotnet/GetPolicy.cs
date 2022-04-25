// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetPolicy
    {
        /// <summary>
        /// `spacelift.Policy` represents a Spacelift **policy** - a collection of customer-defined rules that are applied by Spacelift at one of the decision points within the application.
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
        ///         var policy = Output.Create(Spacelift.GetPolicy.InvokeAsync(new Spacelift.GetPolicyArgs
        ///         {
        ///             PolicyId = spacelift_policy.Policy.Id,
        ///         }));
        ///         this.PolicyBody = policy.Apply(policy =&gt; policy.Body);
        ///     }
        /// 
        ///     [Output("policyBody")]
        ///     public Output&lt;string&gt; PolicyBody { get; set; }
        /// }
        /// ```
        /// 
        /// &lt;!-- schema generated by tfplugindocs --&gt;
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Schema
        /// 
        /// ### Required
        /// 
        /// - **policy_id** (String) immutable ID (slug) of the policy
        /// 
        /// ### Optional
        /// 
        /// - **id** (String) The ID of this resource.
        /// 
        /// ### Read-Only
        /// 
        /// - **body** (String) body of the policy
        /// - **labels** (Set of String)
        /// - **name** (String) name of the policy
        /// - **type** (String) type of the policy
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("spacelift:index/getPolicy:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("policyId", required: true)]
        public string PolicyId { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        public readonly string Body;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Labels;
        public readonly string Name;
        public readonly string PolicyId;
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyResult(
            string body,

            string id,

            ImmutableArray<string> labels,

            string name,

            string policyId,

            string type)
        {
            Body = body;
            Id = id;
            Labels = labels;
            Name = name;
            PolicyId = policyId;
            Type = type;
        }
    }
}
