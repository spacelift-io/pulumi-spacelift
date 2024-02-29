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
    public sealed class GetStackGitlabResult
    {
        /// <summary>
        /// ID of the Gitlab integration
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether this is the default Gitlab integration
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// GitLab namespace of the stack's repository
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private GetStackGitlabResult(
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
