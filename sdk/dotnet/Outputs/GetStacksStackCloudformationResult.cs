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
    public sealed class GetStacksStackCloudformationResult
    {
        public readonly string EntryTemplateFile;
        public readonly string Region;
        public readonly string StackName;
        public readonly string TemplateBucket;

        [OutputConstructor]
        private GetStacksStackCloudformationResult(
            string entryTemplateFile,

            string region,

            string stackName,

            string templateBucket)
        {
            EntryTemplateFile = entryTemplateFile;
            Region = region;
            StackName = stackName;
            TemplateBucket = templateBucket;
        }
    }
}
