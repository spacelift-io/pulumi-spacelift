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
    public sealed class GetModuleGithubEnterpriseResult
    {
        public readonly string Namespace;

        [OutputConstructor]
        private GetModuleGithubEnterpriseResult(string @namespace)
        {
            Namespace = @namespace;
        }
    }
}