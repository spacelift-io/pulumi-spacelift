// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Import is supported using the following syntax
 *
 * ```sh
 *  $ pulumi import spacelift:index/gcpServiceAccount:GcpServiceAccount k8s-core stack/$STACK_ID
 * ```
 *
 * ```sh
 *  $ pulumi import spacelift:index/gcpServiceAccount:GcpServiceAccount k8s-core module/$MODULE_ID
 * ```
 */
export class GcpServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing GcpServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GcpServiceAccountState, opts?: pulumi.CustomResourceOptions): GcpServiceAccount {
        return new GcpServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/gcpServiceAccount:GcpServiceAccount';

    /**
     * Returns true if the given object is an instance of GcpServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GcpServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GcpServiceAccount.__pulumiType;
    }

    /**
     * ID of the module which uses GCP service account credentials
     */
    public readonly moduleId!: pulumi.Output<string | undefined>;
    /**
     * Email address of the GCP service account dedicated for this stack
     */
    public /*out*/ readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    public readonly stackId!: pulumi.Output<string | undefined>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    public readonly tokenScopes!: pulumi.Output<string[]>;

    /**
     * Create a GcpServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GcpServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GcpServiceAccountArgs | GcpServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GcpServiceAccountState | undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["tokenScopes"] = state ? state.tokenScopes : undefined;
        } else {
            const args = argsOrState as GcpServiceAccountArgs | undefined;
            if ((!args || args.tokenScopes === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'tokenScopes'");
            }
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["tokenScopes"] = args ? args.tokenScopes : undefined;
            inputs["serviceAccountEmail"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GcpServiceAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GcpServiceAccount resources.
 */
export interface GcpServiceAccountState {
    /**
     * ID of the module which uses GCP service account credentials
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Email address of the GCP service account dedicated for this stack
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    readonly tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a GcpServiceAccount resource.
 */
export interface GcpServiceAccountArgs {
    /**
     * ID of the module which uses GCP service account credentials
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    readonly tokenScopes: pulumi.Input<pulumi.Input<string>[]>;
}
