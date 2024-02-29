// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.WorkerPool` represents a worker pool assigned to the Spacelift account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const k8s_core = new spacelift.WorkerPool("k8s-core", {
 *     csr: fs.readFileSync("/path/to/csr", { encoding: "base64" }),
 *     description: "Used for all type jobs",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import spacelift:index/workerPool:WorkerPool k8s-core $WORKER_POOL_ID
 * ```
 */
export class WorkerPool extends pulumi.CustomResource {
    /**
     * Get an existing WorkerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkerPoolState, opts?: pulumi.CustomResourceOptions): WorkerPool {
        return new WorkerPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/workerPool:WorkerPool';

    /**
     * Returns true if the given object is an instance of WorkerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkerPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkerPool.__pulumiType;
    }

    /**
     * credentials necessary to connect WorkerPool's workers to the control plane
     */
    public /*out*/ readonly config!: pulumi.Output<string>;
    /**
     * certificate signing request in base64
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * description of the worker pool
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * name of the worker pool
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * private key in base64
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * ID (slug) of the space the worker pool is in
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a WorkerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkerPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkerPoolArgs | WorkerPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkerPoolState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["csr"] = state ? state.csr : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as WorkerPoolArgs | undefined;
            resourceInputs["csr"] = args?.csr ? pulumi.secret(args.csr) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["privateKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["config", "csr", "privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(WorkerPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkerPool resources.
 */
export interface WorkerPoolState {
    /**
     * credentials necessary to connect WorkerPool's workers to the control plane
     */
    config?: pulumi.Input<string>;
    /**
     * certificate signing request in base64
     */
    csr?: pulumi.Input<string>;
    /**
     * description of the worker pool
     */
    description?: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * name of the worker pool
     */
    name?: pulumi.Input<string>;
    /**
     * private key in base64
     */
    privateKey?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the worker pool is in
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkerPool resource.
 */
export interface WorkerPoolArgs {
    /**
     * certificate signing request in base64
     */
    csr?: pulumi.Input<string>;
    /**
     * description of the worker pool
     */
    description?: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * name of the worker pool
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the worker pool is in
     */
    spaceId?: pulumi.Input<string>;
}
