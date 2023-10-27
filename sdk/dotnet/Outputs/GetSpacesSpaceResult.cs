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
    public sealed class GetSpacesSpaceResult
    {
        public readonly string Description;
        public readonly bool InheritEntities;
        public readonly ImmutableArray<string> Labels;
        public readonly string Name;
        public readonly string ParentSpaceId;
        public readonly string SpaceId;

        [OutputConstructor]
        private GetSpacesSpaceResult(
            string description,

            bool inheritEntities,

            ImmutableArray<string> labels,

            string name,

            string parentSpaceId,

            string spaceId)
        {
            Description = description;
            InheritEntities = inheritEntities;
            Labels = labels;
            Name = name;
            ParentSpaceId = parentSpaceId;
            SpaceId = spaceId;
        }
    }
}
