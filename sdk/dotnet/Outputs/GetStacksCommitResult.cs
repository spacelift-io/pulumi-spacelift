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
    public sealed class GetStacksCommitResult
    {
        public readonly ImmutableArray<string> AnyOfs;

        [OutputConstructor]
        private GetStacksCommitResult(ImmutableArray<string> anyOfs)
        {
            AnyOfs = anyOfs;
        }
    }
}
