// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    [SpaceliftResourceType("spacelift:index/stackAwsRole:StackAwsRole")]
    public partial class StackAwsRole : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the module which assumes the AWS IAM role
        /// </summary>
        [Output("moduleId")]
        public Output<string?> ModuleId { get; private set; } = null!;

        /// <summary>
        /// ARN of the AWS IAM role to attach
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// ID of the stack which assumes the AWS IAM role
        /// </summary>
        [Output("stackId")]
        public Output<string?> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a StackAwsRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackAwsRole(string name, StackAwsRoleArgs args, CustomResourceOptions? options = null)
            : base("spacelift:index/stackAwsRole:StackAwsRole", name, args ?? new StackAwsRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StackAwsRole(string name, Input<string> id, StackAwsRoleState? state = null, CustomResourceOptions? options = null)
            : base("spacelift:index/stackAwsRole:StackAwsRole", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StackAwsRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackAwsRole Get(string name, Input<string> id, StackAwsRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new StackAwsRole(name, id, state, options);
        }
    }

    public sealed class StackAwsRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the module which assumes the AWS IAM role
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// ARN of the AWS IAM role to attach
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// ID of the stack which assumes the AWS IAM role
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public StackAwsRoleArgs()
        {
        }
    }

    public sealed class StackAwsRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the module which assumes the AWS IAM role
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// ARN of the AWS IAM role to attach
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// ID of the stack which assumes the AWS IAM role
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public StackAwsRoleState()
        {
        }
    }
}
