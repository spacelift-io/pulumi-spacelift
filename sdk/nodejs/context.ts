// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`spacelift.Stack`) or modules (`spacelift.Module`) using a context attachment (`spacelift.ContextAttachment`)`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const prod_k8s_ie = new spacelift.Context("prod-k8s-ie", {description: "Configuration details for the compute cluster in 🇮🇪"});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import spacelift:index/context:Context prod-k8s-ie $CONTEXT_ID
 * ```
 */
export class Context extends pulumi.CustomResource {
    /**
     * Get an existing Context resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContextState, opts?: pulumi.CustomResourceOptions): Context {
        return new Context(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/context:Context';

    /**
     * Returns true if the given object is an instance of Context.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Context {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Context.__pulumiType;
    }

    /**
     * List of after-apply scripts
     */
    public readonly afterApplies!: pulumi.Output<string[] | undefined>;
    /**
     * List of after-destroy scripts
     */
    public readonly afterDestroys!: pulumi.Output<string[] | undefined>;
    /**
     * List of after-init scripts
     */
    public readonly afterInits!: pulumi.Output<string[] | undefined>;
    /**
     * List of after-perform scripts
     */
    public readonly afterPerforms!: pulumi.Output<string[] | undefined>;
    /**
     * List of after-plan scripts
     */
    public readonly afterPlans!: pulumi.Output<string[] | undefined>;
    /**
     * List of after-run scripts
     */
    public readonly afterRuns!: pulumi.Output<string[] | undefined>;
    /**
     * List of before-apply scripts
     */
    public readonly beforeApplies!: pulumi.Output<string[] | undefined>;
    /**
     * List of before-destroy scripts
     */
    public readonly beforeDestroys!: pulumi.Output<string[] | undefined>;
    /**
     * List of before-init scripts
     */
    public readonly beforeInits!: pulumi.Output<string[] | undefined>;
    /**
     * List of before-perform scripts
     */
    public readonly beforePerforms!: pulumi.Output<string[] | undefined>;
    /**
     * List of before-plan scripts
     */
    public readonly beforePlans!: pulumi.Output<string[] | undefined>;
    /**
     * Free-form context description for users
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the context - should be unique in one account
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID (slug) of the space the context is in
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a Context resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContextArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContextArgs | ContextState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContextState | undefined;
            resourceInputs["afterApplies"] = state ? state.afterApplies : undefined;
            resourceInputs["afterDestroys"] = state ? state.afterDestroys : undefined;
            resourceInputs["afterInits"] = state ? state.afterInits : undefined;
            resourceInputs["afterPerforms"] = state ? state.afterPerforms : undefined;
            resourceInputs["afterPlans"] = state ? state.afterPlans : undefined;
            resourceInputs["afterRuns"] = state ? state.afterRuns : undefined;
            resourceInputs["beforeApplies"] = state ? state.beforeApplies : undefined;
            resourceInputs["beforeDestroys"] = state ? state.beforeDestroys : undefined;
            resourceInputs["beforeInits"] = state ? state.beforeInits : undefined;
            resourceInputs["beforePerforms"] = state ? state.beforePerforms : undefined;
            resourceInputs["beforePlans"] = state ? state.beforePlans : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as ContextArgs | undefined;
            resourceInputs["afterApplies"] = args ? args.afterApplies : undefined;
            resourceInputs["afterDestroys"] = args ? args.afterDestroys : undefined;
            resourceInputs["afterInits"] = args ? args.afterInits : undefined;
            resourceInputs["afterPerforms"] = args ? args.afterPerforms : undefined;
            resourceInputs["afterPlans"] = args ? args.afterPlans : undefined;
            resourceInputs["afterRuns"] = args ? args.afterRuns : undefined;
            resourceInputs["beforeApplies"] = args ? args.beforeApplies : undefined;
            resourceInputs["beforeDestroys"] = args ? args.beforeDestroys : undefined;
            resourceInputs["beforeInits"] = args ? args.beforeInits : undefined;
            resourceInputs["beforePerforms"] = args ? args.beforePerforms : undefined;
            resourceInputs["beforePlans"] = args ? args.beforePlans : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Context.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Context resources.
 */
export interface ContextState {
    /**
     * List of after-apply scripts
     */
    afterApplies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-destroy scripts
     */
    afterDestroys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-init scripts
     */
    afterInits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-perform scripts
     */
    afterPerforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-plan scripts
     */
    afterPlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-run scripts
     */
    afterRuns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-apply scripts
     */
    beforeApplies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-destroy scripts
     */
    beforeDestroys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-init scripts
     */
    beforeInits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-perform scripts
     */
    beforePerforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-plan scripts
     */
    beforePlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Free-form context description for users
     */
    description?: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the context - should be unique in one account
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the context is in
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Context resource.
 */
export interface ContextArgs {
    /**
     * List of after-apply scripts
     */
    afterApplies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-destroy scripts
     */
    afterDestroys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-init scripts
     */
    afterInits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-perform scripts
     */
    afterPerforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-plan scripts
     */
    afterPlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of after-run scripts
     */
    afterRuns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-apply scripts
     */
    beforeApplies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-destroy scripts
     */
    beforeDestroys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-init scripts
     */
    beforeInits?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-perform scripts
     */
    beforePerforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of before-plan scripts
     */
    beforePlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Free-form context description for users
     */
    description?: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the context - should be unique in one account
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the context is in
     */
    spaceId?: pulumi.Input<string>;
}
