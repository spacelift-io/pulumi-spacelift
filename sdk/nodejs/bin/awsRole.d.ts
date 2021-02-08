import * as pulumi from "@pulumi/pulumi";
export declare class AwsRole extends pulumi.CustomResource {
    /**
     * Get an existing AwsRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsRoleState, opts?: pulumi.CustomResourceOptions): AwsRole;
    /**
     * Returns true if the given object is an instance of AwsRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is AwsRole;
    /**
     * ID of the module which assumes the AWS IAM role
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * ARN of the AWS IAM role to attach
     */
    readonly roleArn: pulumi.Output<string>;
    /**
     * ID of the stack which assumes the AWS IAM role
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * Create a AwsRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsRoleArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering AwsRole resources.
 */
export interface AwsRoleState {
    /**
     * ID of the module which assumes the AWS IAM role
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ARN of the AWS IAM role to attach
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * ID of the stack which assumes the AWS IAM role
     */
    readonly stackId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a AwsRole resource.
 */
export interface AwsRoleArgs {
    /**
     * ID of the module which assumes the AWS IAM role
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ARN of the AWS IAM role to attach
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * ID of the stack which assumes the AWS IAM role
     */
    readonly stackId?: pulumi.Input<string>;
}
