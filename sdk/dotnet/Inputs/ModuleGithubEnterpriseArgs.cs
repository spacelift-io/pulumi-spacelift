// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Inputs
{

    public sealed class ModuleGithubEnterpriseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GitHub organization / user the repository belongs to
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public ModuleGithubEnterpriseArgs()
        {
        }
        public static new ModuleGithubEnterpriseArgs Empty => new ModuleGithubEnterpriseArgs();
    }
}
