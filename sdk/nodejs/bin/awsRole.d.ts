import * as pulumi from "@pulumi/pulumi";
/**
 * `spacelift.AwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.
 *
 * If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).
 *
 * Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and Run ID as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).
 *
 * ## Schema
 *
 * ### Required
 *
 * - **role_arn** (String) ARN of the AWS IAM role to attach
 *
 * ### Optional
 *
 * - **external_id** (String) Custom external ID (works only for private workers).
 * - **generate_credentials_in_worker** (Boolean) Generate AWS credentials in the private worker
 * - **id** (String) The ID of this resource.
 * - **module_id** (String) ID of the module which assumes the AWS IAM role
 * - **stack_id** (String) ID of the stack which assumes the AWS IAM role
 *
 * ## Import
 *
 * Import is supported using the following syntax
 *
 * ```sh
 *  $ pulumi import spacelift:index/awsRole:AwsRole k8s-core stack/$STACK_ID
 * ```
 *
 * ```sh
 *  $ pulumi import spacelift:index/awsRole:AwsRole k8s-core module/$MODULE_ID
 * ```
 */
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
     * Custom external ID (works only for private workers).
     */
    readonly externalId: pulumi.Output<string | undefined>;
    /**
     * Generate AWS credentials in the private worker
     */
    readonly generateCredentialsInWorker: pulumi.Output<boolean | undefined>;
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
     * Custom external ID (works only for private workers).
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * Generate AWS credentials in the private worker
     */
    readonly generateCredentialsInWorker?: pulumi.Input<boolean>;
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
     * Custom external ID (works only for private workers).
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * Generate AWS credentials in the private worker
     */
    readonly generateCredentialsInWorker?: pulumi.Input<boolean>;
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
