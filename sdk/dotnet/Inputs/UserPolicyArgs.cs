// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift.Inputs
{

    public sealed class UserPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of access to the space. Possible values are: READ, WRITE, ADMIN
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// ID (slug) of the space the user has access to
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        public UserPolicyArgs()
        {
        }
        public static new UserPolicyArgs Empty => new UserPolicyArgs();
    }
}
