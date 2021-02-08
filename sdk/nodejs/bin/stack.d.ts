import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
export declare class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackState, opts?: pulumi.CustomResourceOptions): Stack;
    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Stack;
    /**
     * Indicates whether this stack can manage others
     */
    readonly administrative: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether changes to this stack can be automatically deployed
     */
    readonly autodeploy: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether obsolete proposed changes should automatically be retried
     */
    readonly autoretry: pulumi.Output<boolean | undefined>;
    /**
     * AWS IAM assume role policy statement setting up trust relationship
     */
    readonly awsAssumeRolePolicyStatement: pulumi.Output<string>;
    /**
     * List of before-init scripts
     */
    readonly beforeInits: pulumi.Output<string[] | undefined>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch: pulumi.Output<string>;
    /**
     * CloudFormation-specific configuration. Presence means this Stack is a CloudFormation Stack.
     */
    readonly cloudformation: pulumi.Output<outputs.StackCloudformation | undefined>;
    /**
     * Free-form stack description for users
     */
    readonly description: pulumi.Output<string | undefined>;
    readonly gitlab: pulumi.Output<outputs.StackGitlab | undefined>;
    /**
     * State file to upload when creating a new stack
     */
    readonly importState: pulumi.Output<string | undefined>;
    readonly labels: pulumi.Output<string[] | undefined>;
    /**
     * Determines if Spacelift should manage state for this stack
     */
    readonly manageState: pulumi.Output<boolean | undefined>;
    /**
     * Name of the stack - should be unique in one account
     */
    readonly name: pulumi.Output<string>;
    /**
     * Project root is the optional directory relative to the workspace root containing the entrypoint to the Stack.
     */
    readonly projectRoot: pulumi.Output<string | undefined>;
    /**
     * Pulumi-specific configuration. Presence means this Stack is a Pulumi Stack.
     */
    readonly pulumi: pulumi.Output<outputs.StackPulumi | undefined>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository: pulumi.Output<string>;
    /**
     * Name of the Docker image used to process Runs
     */
    readonly runnerImage: pulumi.Output<string | undefined>;
    /**
     * Terraform version to use
     */
    readonly terraformVersion: pulumi.Output<string | undefined>;
    /**
     * Terraform workspace to select
     */
    readonly terraformWorkspace: pulumi.Output<string | undefined>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId: pulumi.Output<string | undefined>;
    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Stack resources.
 */
export interface StackState {
    /**
     * Indicates whether this stack can manage others
     */
    readonly administrative?: pulumi.Input<boolean>;
    /**
     * Indicates whether changes to this stack can be automatically deployed
     */
    readonly autodeploy?: pulumi.Input<boolean>;
    /**
     * Indicates whether obsolete proposed changes should automatically be retried
     */
    readonly autoretry?: pulumi.Input<boolean>;
    /**
     * AWS IAM assume role policy statement setting up trust relationship
     */
    readonly awsAssumeRolePolicyStatement?: pulumi.Input<string>;
    /**
     * List of before-init scripts
     */
    readonly beforeInits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * CloudFormation-specific configuration. Presence means this Stack is a CloudFormation Stack.
     */
    readonly cloudformation?: pulumi.Input<inputs.StackCloudformation>;
    /**
     * Free-form stack description for users
     */
    readonly description?: pulumi.Input<string>;
    readonly gitlab?: pulumi.Input<inputs.StackGitlab>;
    /**
     * State file to upload when creating a new stack
     */
    readonly importState?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines if Spacelift should manage state for this stack
     */
    readonly manageState?: pulumi.Input<boolean>;
    /**
     * Name of the stack - should be unique in one account
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Project root is the optional directory relative to the workspace root containing the entrypoint to the Stack.
     */
    readonly projectRoot?: pulumi.Input<string>;
    /**
     * Pulumi-specific configuration. Presence means this Stack is a Pulumi Stack.
     */
    readonly pulumi?: pulumi.Input<inputs.StackPulumi>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * Name of the Docker image used to process Runs
     */
    readonly runnerImage?: pulumi.Input<string>;
    /**
     * Terraform version to use
     */
    readonly terraformVersion?: pulumi.Input<string>;
    /**
     * Terraform workspace to select
     */
    readonly terraformWorkspace?: pulumi.Input<string>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    /**
     * Indicates whether this stack can manage others
     */
    readonly administrative?: pulumi.Input<boolean>;
    /**
     * Indicates whether changes to this stack can be automatically deployed
     */
    readonly autodeploy?: pulumi.Input<boolean>;
    /**
     * Indicates whether obsolete proposed changes should automatically be retried
     */
    readonly autoretry?: pulumi.Input<boolean>;
    /**
     * List of before-init scripts
     */
    readonly beforeInits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch: pulumi.Input<string>;
    /**
     * CloudFormation-specific configuration. Presence means this Stack is a CloudFormation Stack.
     */
    readonly cloudformation?: pulumi.Input<inputs.StackCloudformation>;
    /**
     * Free-form stack description for users
     */
    readonly description?: pulumi.Input<string>;
    readonly gitlab?: pulumi.Input<inputs.StackGitlab>;
    /**
     * State file to upload when creating a new stack
     */
    readonly importState?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines if Spacelift should manage state for this stack
     */
    readonly manageState?: pulumi.Input<boolean>;
    /**
     * Name of the stack - should be unique in one account
     */
    readonly name: pulumi.Input<string>;
    /**
     * Project root is the optional directory relative to the workspace root containing the entrypoint to the Stack.
     */
    readonly projectRoot?: pulumi.Input<string>;
    /**
     * Pulumi-specific configuration. Presence means this Stack is a Pulumi Stack.
     */
    readonly pulumi?: pulumi.Input<inputs.StackPulumi>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository: pulumi.Input<string>;
    /**
     * Name of the Docker image used to process Runs
     */
    readonly runnerImage?: pulumi.Input<string>;
    /**
     * Terraform version to use
     */
    readonly terraformVersion?: pulumi.Input<string>;
    /**
     * Terraform workspace to select
     */
    readonly terraformWorkspace?: pulumi.Input<string>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId?: pulumi.Input<string>;
}
