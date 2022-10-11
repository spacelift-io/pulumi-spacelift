// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Inputs
{

    public sealed class StackPulumiGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// State backend to log into on Run initialize.
        /// </summary>
        [Input("loginUrl", required: true)]
        public Input<string> LoginUrl { get; set; } = null!;

        /// <summary>
        /// Pulumi stack name to use with the state backend.
        /// </summary>
        [Input("stackName", required: true)]
        public Input<string> StackName { get; set; } = null!;

        public StackPulumiGetArgs()
        {
        }
        public static new StackPulumiGetArgs Empty => new StackPulumiGetArgs();
    }
}
