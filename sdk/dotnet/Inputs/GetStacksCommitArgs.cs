// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Inputs
{

    public sealed class GetStacksCommitInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("anyOfs", required: true)]
        private InputList<string>? _anyOfs;
        public InputList<string> AnyOfs
        {
            get => _anyOfs ?? (_anyOfs = new InputList<string>());
            set => _anyOfs = value;
        }

        public GetStacksCommitInputArgs()
        {
        }
        public static new GetStacksCommitInputArgs Empty => new GetStacksCommitInputArgs();
    }
}
