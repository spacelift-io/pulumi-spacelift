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
    public sealed class GetModuleBitbucketDatacenterResult
    {
        /// <summary>
        /// ID of the Bitbucket Datacenter integration
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether this is the default Bitbucket Datacenter integration
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// Bitbucket Datacenter namespace of the stack's repository
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private GetModuleBitbucketDatacenterResult(
            string id,

            bool isDefault,

            string @namespace)
        {
            Id = id;
            IsDefault = isDefault;
            Namespace = @namespace;
        }
    }
}
