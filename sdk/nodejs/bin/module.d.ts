import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
/**
 * ## Import
 *
 * Import is supported using the following syntax
 *
 * ```sh
 *  $ pulumi import spacelift:index/module:Module k8s-module $MODULE_ID
 * ```
 */
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
     * Azure DevOps VCS settings
     */
    readonly azureDevops: pulumi.Output<outputs.ModuleAzureDevops | undefined>;
    /**
     * Bitbucket Cloud VCS settings
     */
    readonly bitbucketCloud: pulumi.Output<outputs.ModuleBitbucketCloud | undefined>;
    /**
     * Bitbucket Datacenter VCS settings
     */
    readonly bitbucketDatacenter: pulumi.Output<outputs.ModuleBitbucketDatacenter | undefined>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch: pulumi.Output<string>;
    /**
     * Free-form module description for users
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * GitHub Enterprise (self-hosted) VCS settings
     */
    readonly githubEnterprise: pulumi.Output<outputs.ModuleGithubEnterprise | undefined>;
    /**
     * GitLab VCS settings
     */
    readonly gitlab: pulumi.Output<outputs.ModuleGitlab | undefined>;
    readonly labels: pulumi.Output<string[] | undefined>;
    /**
     * The module name will by default be inferred from the repository name if it follows the terraform-provider-name naming
     * convention. However, if the repository doesn't follow this convention, or you want to give it a custom name, you can
     * provide it here.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Project root is the optional directory relative to the repository root containing the module source code.
     */
    readonly projectRoot: pulumi.Output<string | undefined>;
    /**
     * Protect this module from accidental deletion. If set, attempts to delete this module will fail.
     */
    readonly protectFromDeletion: pulumi.Output<boolean | undefined>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository: pulumi.Output<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    readonly sharedAccounts: pulumi.Output<string[] | undefined>;
    /**
     * The module provider will by default be inferred from the repository name if it follows the terraform-provider-name
     * naming convention. However, if the repository doesn't follow this convention, or you gave the module a custom name, you
     * can provide the provider name here.
     */
    readonly terraformProvider: pulumi.Output<string>;
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
     * Azure DevOps VCS settings
     */
    readonly azureDevops?: pulumi.Input<inputs.ModuleAzureDevops>;
    /**
     * Bitbucket Cloud VCS settings
     */
    readonly bitbucketCloud?: pulumi.Input<inputs.ModuleBitbucketCloud>;
    /**
     * Bitbucket Datacenter VCS settings
     */
    readonly bitbucketDatacenter?: pulumi.Input<inputs.ModuleBitbucketDatacenter>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * Free-form module description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * GitHub Enterprise (self-hosted) VCS settings
     */
    readonly githubEnterprise?: pulumi.Input<inputs.ModuleGithubEnterprise>;
    /**
     * GitLab VCS settings
     */
    readonly gitlab?: pulumi.Input<inputs.ModuleGitlab>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The module name will by default be inferred from the repository name if it follows the terraform-provider-name naming
     * convention. However, if the repository doesn't follow this convention, or you want to give it a custom name, you can
     * provide it here.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Project root is the optional directory relative to the repository root containing the module source code.
     */
    readonly projectRoot?: pulumi.Input<string>;
    /**
     * Protect this module from accidental deletion. If set, attempts to delete this module will fail.
     */
    readonly protectFromDeletion?: pulumi.Input<boolean>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    readonly sharedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The module provider will by default be inferred from the repository name if it follows the terraform-provider-name
     * naming convention. However, if the repository doesn't follow this convention, or you gave the module a custom name, you
     * can provide the provider name here.
     */
    readonly terraformProvider?: pulumi.Input<string>;
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
     * Azure DevOps VCS settings
     */
    readonly azureDevops?: pulumi.Input<inputs.ModuleAzureDevops>;
    /**
     * Bitbucket Cloud VCS settings
     */
    readonly bitbucketCloud?: pulumi.Input<inputs.ModuleBitbucketCloud>;
    /**
     * Bitbucket Datacenter VCS settings
     */
    readonly bitbucketDatacenter?: pulumi.Input<inputs.ModuleBitbucketDatacenter>;
    /**
     * GitHub branch to apply changes to
     */
    readonly branch: pulumi.Input<string>;
    /**
     * Free-form module description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * GitHub Enterprise (self-hosted) VCS settings
     */
    readonly githubEnterprise?: pulumi.Input<inputs.ModuleGithubEnterprise>;
    /**
     * GitLab VCS settings
     */
    readonly gitlab?: pulumi.Input<inputs.ModuleGitlab>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The module name will by default be inferred from the repository name if it follows the terraform-provider-name naming
     * convention. However, if the repository doesn't follow this convention, or you want to give it a custom name, you can
     * provide it here.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Project root is the optional directory relative to the repository root containing the module source code.
     */
    readonly projectRoot?: pulumi.Input<string>;
    /**
     * Protect this module from accidental deletion. If set, attempts to delete this module will fail.
     */
    readonly protectFromDeletion?: pulumi.Input<boolean>;
    /**
     * Name of the repository, without the owner part
     */
    readonly repository: pulumi.Input<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    readonly sharedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The module provider will by default be inferred from the repository name if it follows the terraform-provider-name
     * naming convention. However, if the repository doesn't follow this convention, or you gave the module a custom name, you
     * can provide the provider name here.
     */
    readonly terraformProvider?: pulumi.Input<string>;
    /**
     * ID of the worker pool to use
     */
    readonly workerPoolId?: pulumi.Input<string>;
}
