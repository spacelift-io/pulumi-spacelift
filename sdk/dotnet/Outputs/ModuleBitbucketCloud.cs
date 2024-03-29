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
    public sealed class ModuleBitbucketCloud
    {
        /// <summary>
        /// The ID of the Bitbucket Cloud integration. If not specified, the default integration will be used.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Indicates whether this is the default Bitbucket Cloud integration
        /// </summary>
        public readonly bool? IsDefault;
        /// <summary>
        /// The Bitbucket project containing the repository
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private ModuleBitbucketCloud(
            string? id,

            bool? isDefault,

            string @namespace)
        {
            Id = id;
            IsDefault = isDefault;
            Namespace = @namespace;
        }
    }
}
