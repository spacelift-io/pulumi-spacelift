// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetIPs
    {
        /// <summary>
        /// `spacelift.getIPs` returns the list of Spacelift's outgoing IP addresses, which you can use to whitelist connections coming from the Spacelift's "mothership". **NOTE:** this does not include the IP addresses of the workers in Spacelift's public worker pool. If you need to ensure that requests made during runs originate from a known set of IP addresses, please consider setting up a [private worker pool](https://docs.spacelift.io/concepts/worker-pools).
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
        ///     var ips = Spacelift.GetIPs.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIPsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIPsResult>("spacelift:index/getIPs:getIPs", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// `spacelift.getIPs` returns the list of Spacelift's outgoing IP addresses, which you can use to whitelist connections coming from the Spacelift's "mothership". **NOTE:** this does not include the IP addresses of the workers in Spacelift's public worker pool. If you need to ensure that requests made during runs originate from a known set of IP addresses, please consider setting up a [private worker pool](https://docs.spacelift.io/concepts/worker-pools).
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
        ///     var ips = Spacelift.GetIPs.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIPsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIPsResult>("spacelift:index/getIPs:getIPs", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIPsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the list of spacelift.io outgoing IP addresses
        /// </summary>
        public readonly ImmutableArray<string> Ips;

        [OutputConstructor]
        private GetIPsResult(
            string id,

            ImmutableArray<string> ips)
        {
            Id = id;
            Ips = ips;
        }
    }
}
