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
    public sealed class ModuleAzureDevops
    {
        /// <summary>
        /// The name of the Azure DevOps project
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private ModuleAzureDevops(string project)
        {
            Project = project;
        }
    }
}
