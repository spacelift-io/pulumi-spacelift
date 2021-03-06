// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
     * Free-form context description for users
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the context - should be unique in one account
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Context resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContextArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContextArgs | ContextState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ContextState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ContextArgs | undefined;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Context.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Context resources.
 */
export interface ContextState {
    /**
     * Free-form context description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the context - should be unique in one account
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Context resource.
 */
export interface ContextArgs {
    /**
     * Free-form context description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the context - should be unique in one account
     */
    readonly name: pulumi.Input<string>;
}
