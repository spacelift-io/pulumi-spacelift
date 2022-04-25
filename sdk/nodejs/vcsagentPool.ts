// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.VCSAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const ghe = new spacelift.VCSAgentPool("ghe", {
 *     description: "VCS agent pool for our internal GitHub Enterpris",
 *     name: "ghe",
 * });
 * ```
 *
 * <!-- schema generated by tfplugindocs -->
 * ## Schema
 *
 * ### Required
 *
 * - **name** (String) Name of the VCS agent pool, must be unique within an account
 *
 * ### Optional
 *
 * - **description** (String) Free-form VCS agent pool description for users
 * - **id** (String) The ID of this resource.
 *
 * ### Read-Only
 *
 * - **config** (String, Sensitive) VCS agent pool configuration, encoded using base64
 *
 * ## Import
 *
 * Import is supported using the following syntax
 *
 * ```sh
 *  $ pulumi import spacelift:index/vCSAgentPool:VCSAgentPool ghe $VCS_AGENT_POOL_ID
 * ```
 */
export class VCSAgentPool extends pulumi.CustomResource {
    /**
     * Get an existing VCSAgentPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VCSAgentPoolState, opts?: pulumi.CustomResourceOptions): VCSAgentPool {
        return new VCSAgentPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/vCSAgentPool:VCSAgentPool';

    /**
     * Returns true if the given object is an instance of VCSAgentPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VCSAgentPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VCSAgentPool.__pulumiType;
    }

    /**
     * VCS agent pool configuration, encoded using base64
     */
    public /*out*/ readonly config!: pulumi.Output<string>;
    /**
     * Free-form VCS agent pool description for users
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the VCS agent pool, must be unique within an account
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a VCSAgentPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VCSAgentPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VCSAgentPoolArgs | VCSAgentPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VCSAgentPoolState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as VCSAgentPoolArgs | undefined;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["config"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VCSAgentPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VCSAgentPool resources.
 */
export interface VCSAgentPoolState {
    /**
     * VCS agent pool configuration, encoded using base64
     */
    readonly config?: pulumi.Input<string>;
    /**
     * Free-form VCS agent pool description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the VCS agent pool, must be unique within an account
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VCSAgentPool resource.
 */
export interface VCSAgentPoolArgs {
    /**
     * Free-form VCS agent pool description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the VCS agent pool, must be unique within an account
     */
    readonly name: pulumi.Input<string>;
}