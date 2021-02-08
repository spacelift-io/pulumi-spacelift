import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
export declare class Module extends pulumi.CustomResource {
    /**
     * Get an existing Module resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModuleState, opts?: pulumi.CustomResourceOptions): Module;
    /**
     * Returns true if the given object is an instance of Module.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Module;
    /**
     * Indicates whether this module can manage others
     */
    readonly administrative: pulumi.Output<boolean | undefined>;
    /**
     * AWS IAM assume role policy statement setting up trust relationship
     */
    readonly awsAssumeRolePolicyStatement: pulumi.Output<string>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch: pulumi.Output<string>;
    /**
     * Free-form module description for users
     */
    readonly description: pulumi.Output<string | undefined>;
    readonly gitlab: pulumi.Output<outputs.ModuleGitlab | undefined>;
    readonly labels: pulumi.Output<string[] | undefined>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository: pulumi.Output<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    readonly sharedAccounts: pulumi.Output<string[] | undefined>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId: pulumi.Output<string | undefined>;
    /**
     * Create a Module resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Module resources.
 */
export interface ModuleState {
    /**
     * Indicates whether this module can manage others
     */
    readonly administrative?: pulumi.Input<boolean>;
    /**
     * AWS IAM assume role policy statement setting up trust relationship
     */
    readonly awsAssumeRolePolicyStatement?: pulumi.Input<string>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * Free-form module description for users
     */
    readonly description?: pulumi.Input<string>;
    readonly gitlab?: pulumi.Input<inputs.ModuleGitlab>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    readonly sharedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Module resource.
 */
export interface ModuleArgs {
    /**
     * Indicates whether this module can manage others
     */
    readonly administrative?: pulumi.Input<boolean>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch: pulumi.Input<string>;
    /**
     * Free-form module description for users
     */
    readonly description?: pulumi.Input<string>;
    readonly gitlab?: pulumi.Input<inputs.ModuleGitlab>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository: pulumi.Input<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    readonly sharedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId?: pulumi.Input<string>;
}
