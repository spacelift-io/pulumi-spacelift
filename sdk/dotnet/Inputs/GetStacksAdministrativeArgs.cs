// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Inputs
{

    public sealed class GetStacksAdministrativeInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("equals")]
        public Input<bool>? Equals { get; set; }

        public GetStacksAdministrativeInputArgs()
        {
        }
        public static new GetStacksAdministrativeInputArgs Empty => new GetStacksAdministrativeInputArgs();
    }
}
