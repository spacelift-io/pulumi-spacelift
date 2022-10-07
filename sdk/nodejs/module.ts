// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * // Explicit module name and provider:
 * const k8s_module = new spacelift.Module("k8s-module", {
 *     administrative: true,
 *     branch: "master",
 *     description: "Infra terraform module",
 *     repository: "terraform-super-module",
 *     terraformProvider: "aws",
 * });
 * // Unspecified module name and provider (repository naming scheme terraform-${provider}-${name})
 * const example_module = new spacelift.Module("example-module", {
 *     administrative: true,
 *     branch: "master",
 *     description: "Example terraform module",
 *     projectRoot: "example",
 *     repository: "terraform-aws-example",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import spacelift:index/module:Module k8s-module $MODULE_ID
 * ```
 */
export class Module extends pulumi.CustomResource {
    /**
     * Get an existing Module resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModuleState, opts?: pulumi.CustomResourceOptions): Module {
        return new Module(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/module:Module';

    /**
     * Returns true if the given object is an instance of Module.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Module {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Module.__pulumiType;
    }

    /**
     * Indicates whether this module can manage others. Defaults to `false`.
     */
    public readonly administrative!: pulumi.Output<boolean | undefined>;
    /**
     * AWS IAM assume role policy statement setting up trust relationship
     */
    public /*out*/ readonly awsAssumeRolePolicyStatement!: pulumi.Output<string>;
    /**
     * Azure DevOps VCS settings
     */
    public readonly azureDevops!: pulumi.Output<outputs.ModuleAzureDevops | undefined>;
    /**
     * Bitbucket Cloud VCS settings
     */
    public readonly bitbucketCloud!: pulumi.Output<outputs.ModuleBitbucketCloud | undefined>;
    /**
     * Bitbucket Datacenter VCS settings
     */
    public readonly bitbucketDatacenter!: pulumi.Output<outputs.ModuleBitbucketDatacenter | undefined>;
    /**
     * GitHub branch to apply changes to
     */
    public readonly branch!: pulumi.Output<string>;
    /**
     * Free-form module description for users
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * GitHub Enterprise (self-hosted) VCS settings
     */
    public readonly githubEnterprise!: pulumi.Output<outputs.ModuleGithubEnterprise | undefined>;
    /**
     * GitLab VCS settings
     */
    public readonly gitlab!: pulumi.Output<outputs.ModuleGitlab | undefined>;
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The module name will by default be inferred from the repository name if it follows the terraform-provider-name naming
     * convention. However, if the repository doesn't follow this convention, or you want to give it a custom name, you can
     * provide it here.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project root is the optional directory relative to the repository root containing the module source code.
     */
    public readonly projectRoot!: pulumi.Output<string | undefined>;
    /**
     * Protect this module from accidental deletion. If set, attempts to delete this module will fail. Defaults to `false`.
     */
    public readonly protectFromDeletion!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the repository, without the owner part
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    public readonly sharedAccounts!: pulumi.Output<string[] | undefined>;
    /**
     * ID (slug) of the space the module is in
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The module provider will by default be inferred from the repository name if it follows the terraform-provider-name
     * naming convention. However, if the repository doesn't follow this convention, or you gave the module a custom name, you
     * can provide the provider name here.
     */
    public readonly terraformProvider!: pulumi.Output<string>;
    /**
     * ID of the worker pool to use
     */
    public readonly workerPoolId!: pulumi.Output<string | undefined>;

    /**
     * Create a Module resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModuleArgs | ModuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModuleState | undefined;
            resourceInputs["administrative"] = state ? state.administrative : undefined;
            resourceInputs["awsAssumeRolePolicyStatement"] = state ? state.awsAssumeRolePolicyStatement : undefined;
            resourceInputs["azureDevops"] = state ? state.azureDevops : undefined;
            resourceInputs["bitbucketCloud"] = state ? state.bitbucketCloud : undefined;
            resourceInputs["bitbucketDatacenter"] = state ? state.bitbucketDatacenter : undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["githubEnterprise"] = state ? state.githubEnterprise : undefined;
            resourceInputs["gitlab"] = state ? state.gitlab : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectRoot"] = state ? state.projectRoot : undefined;
            resourceInputs["protectFromDeletion"] = state ? state.protectFromDeletion : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["sharedAccounts"] = state ? state.sharedAccounts : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["terraformProvider"] = state ? state.terraformProvider : undefined;
            resourceInputs["workerPoolId"] = state ? state.workerPoolId : undefined;
        } else {
            const args = argsOrState as ModuleArgs | undefined;
            if ((!args || args.branch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'branch'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["administrative"] = args ? args.administrative : undefined;
            resourceInputs["azureDevops"] = args ? args.azureDevops : undefined;
            resourceInputs["bitbucketCloud"] = args ? args.bitbucketCloud : undefined;
            resourceInputs["bitbucketDatacenter"] = args ? args.bitbucketDatacenter : undefined;
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["githubEnterprise"] = args ? args.githubEnterprise : undefined;
            resourceInputs["gitlab"] = args ? args.gitlab : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectRoot"] = args ? args.projectRoot : undefined;
            resourceInputs["protectFromDeletion"] = args ? args.protectFromDeletion : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["sharedAccounts"] = args ? args.sharedAccounts : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["terraformProvider"] = args ? args.terraformProvider : undefined;
            resourceInputs["workerPoolId"] = args ? args.workerPoolId : undefined;
            resourceInputs["awsAssumeRolePolicyStatement"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Module.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Module resources.
 */
export interface ModuleState {
    /**
     * Indicates whether this module can manage others. Defaults to `false`.
     */
    administrative?: pulumi.Input<boolean>;
    /**
     * AWS IAM assume role policy statement setting up trust relationship
     */
    awsAssumeRolePolicyStatement?: pulumi.Input<string>;
    /**
     * Azure DevOps VCS settings
     */
    azureDevops?: pulumi.Input<inputs.ModuleAzureDevops>;
    /**
     * Bitbucket Cloud VCS settings
     */
    bitbucketCloud?: pulumi.Input<inputs.ModuleBitbucketCloud>;
    /**
     * Bitbucket Datacenter VCS settings
     */
    bitbucketDatacenter?: pulumi.Input<inputs.ModuleBitbucketDatacenter>;
    /**
     * GitHub branch to apply changes to
     */
    branch?: pulumi.Input<string>;
    /**
     * Free-form module description for users
     */
    description?: pulumi.Input<string>;
    /**
     * GitHub Enterprise (self-hosted) VCS settings
     */
    githubEnterprise?: pulumi.Input<inputs.ModuleGithubEnterprise>;
    /**
     * GitLab VCS settings
     */
    gitlab?: pulumi.Input<inputs.ModuleGitlab>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The module name will by default be inferred from the repository name if it follows the terraform-provider-name naming
     * convention. However, if the repository doesn't follow this convention, or you want to give it a custom name, you can
     * provide it here.
     */
    name?: pulumi.Input<string>;
    /**
     * Project root is the optional directory relative to the repository root containing the module source code.
     */
    projectRoot?: pulumi.Input<string>;
    /**
     * Protect this module from accidental deletion. If set, attempts to delete this module will fail. Defaults to `false`.
     */
    protectFromDeletion?: pulumi.Input<boolean>;
    /**
     * Name of the repository, without the owner part
     */
    repository?: pulumi.Input<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    sharedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID (slug) of the space the module is in
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The module provider will by default be inferred from the repository name if it follows the terraform-provider-name
     * naming convention. However, if the repository doesn't follow this convention, or you gave the module a custom name, you
     * can provide the provider name here.
     */
    terraformProvider?: pulumi.Input<string>;
    /**
     * ID of the worker pool to use
     */
    workerPoolId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Module resource.
 */
export interface ModuleArgs {
    /**
     * Indicates whether this module can manage others. Defaults to `false`.
     */
    administrative?: pulumi.Input<boolean>;
    /**
     * Azure DevOps VCS settings
     */
    azureDevops?: pulumi.Input<inputs.ModuleAzureDevops>;
    /**
     * Bitbucket Cloud VCS settings
     */
    bitbucketCloud?: pulumi.Input<inputs.ModuleBitbucketCloud>;
    /**
     * Bitbucket Datacenter VCS settings
     */
    bitbucketDatacenter?: pulumi.Input<inputs.ModuleBitbucketDatacenter>;
    /**
     * GitHub branch to apply changes to
     */
    branch: pulumi.Input<string>;
    /**
     * Free-form module description for users
     */
    description?: pulumi.Input<string>;
    /**
     * GitHub Enterprise (self-hosted) VCS settings
     */
    githubEnterprise?: pulumi.Input<inputs.ModuleGithubEnterprise>;
    /**
     * GitLab VCS settings
     */
    gitlab?: pulumi.Input<inputs.ModuleGitlab>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The module name will by default be inferred from the repository name if it follows the terraform-provider-name naming
     * convention. However, if the repository doesn't follow this convention, or you want to give it a custom name, you can
     * provide it here.
     */
    name?: pulumi.Input<string>;
    /**
     * Project root is the optional directory relative to the repository root containing the module source code.
     */
    projectRoot?: pulumi.Input<string>;
    /**
     * Protect this module from accidental deletion. If set, attempts to delete this module will fail. Defaults to `false`.
     */
    protectFromDeletion?: pulumi.Input<boolean>;
    /**
     * Name of the repository, without the owner part
     */
    repository: pulumi.Input<string>;
    /**
     * List of the accounts (subdomains) which should have access to the Module
     */
    sharedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID (slug) of the space the module is in
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The module provider will by default be inferred from the repository name if it follows the terraform-provider-name
     * naming convention. However, if the repository doesn't follow this convention, or you gave the module a custom name, you
     * can provide the provider name here.
     */
    terraformProvider?: pulumi.Input<string>;
    /**
     * ID of the worker pool to use
     */
    workerPoolId?: pulumi.Input<string>;
}
