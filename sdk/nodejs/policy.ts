// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.Policy` represents a Spacelift **policy** - a collection of customer-defined rules that are applied by Spacelift at one of the decision points within the application.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const no_weekend_deploysPolicy = new spacelift.Policy("no-weekend-deploysPolicy", {
 *     body: fs.readFileSync(`${path.module}/policies/no-weekend-deploys.rego`, "utf8"),
 *     type: "PLAN",
 * });
 * const core_infra_production = new spacelift.Stack("core-infra-production", {
 *     branch: "master",
 *     repository: "core-infra",
 * });
 * const no_weekend_deploysPolicyAttachment = new spacelift.PolicyAttachment("no-weekend-deploysPolicyAttachment", {
 *     policyId: no_weekend_deploysPolicy.id,
 *     stackId: core_infra_production.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import spacelift:index/policy:Policy no-weekend-deploys $POLICY_ID
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Body of the policy
     */
    public readonly body!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the policy - should be unique in one account
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID (slug) of the space the policy is in
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * Type of the policy. Possible values are `ACCESS`, `APPROVAL`, `GIT_PUSH`, `INITIALIZATION`, `LOGIN`, `PLAN`, `TASK`, `TRIGGER` and `NOTIFICATION`. Deprecated values are `STACK_ACCESS` (use `ACCESS` instead), `TASK_RUN` (use `TASK` instead), and `TERRAFORM_PLAN` (use `PLAN` instead).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.body === undefined) && !opts.urn) {
                throw new Error("Missing required property 'body'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Body of the policy
     */
    body?: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the policy - should be unique in one account
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the policy is in
     */
    spaceId?: pulumi.Input<string>;
    /**
     * Type of the policy. Possible values are `ACCESS`, `APPROVAL`, `GIT_PUSH`, `INITIALIZATION`, `LOGIN`, `PLAN`, `TASK`, `TRIGGER` and `NOTIFICATION`. Deprecated values are `STACK_ACCESS` (use `ACCESS` instead), `TASK_RUN` (use `TASK` instead), and `TERRAFORM_PLAN` (use `PLAN` instead).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Body of the policy
     */
    body: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the policy - should be unique in one account
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the policy is in
     */
    spaceId?: pulumi.Input<string>;
    /**
     * Type of the policy. Possible values are `ACCESS`, `APPROVAL`, `GIT_PUSH`, `INITIALIZATION`, `LOGIN`, `PLAN`, `TASK`, `TRIGGER` and `NOTIFICATION`. Deprecated values are `STACK_ACCESS` (use `ACCESS` instead), `TASK_RUN` (use `TASK` instead), and `TERRAFORM_PLAN` (use `PLAN` instead).
     */
    type: pulumi.Input<string>;
}
