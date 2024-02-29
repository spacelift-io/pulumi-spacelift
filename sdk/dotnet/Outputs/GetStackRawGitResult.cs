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
    public sealed class GetStackRawGitResult
    {
        /// <summary>
        /// User-friendly namespace for the repository, this is for cosmetic purposes only
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// HTTPS URL of the Git repository
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetStackRawGitResult(
            string @namespace,

            string url)
        {
            Namespace = @namespace;
            Url = url;
        }
    }
}
