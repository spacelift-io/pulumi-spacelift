// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Inputs
{

    public sealed class StackAnsibleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The playbook Ansible should run.
        /// </summary>
        [Input("playbook", required: true)]
        public Input<string> Playbook { get; set; } = null!;

        public StackAnsibleArgs()
        {
        }
        public static new StackAnsibleArgs Empty => new StackAnsibleArgs();
    }
}
