import * as pulumi from "@pulumi/pulumi";
/**
 * `spacelift.WorkerPool` represents a worker pool assigned to the Spacelift account.
 *
 * ## Schema
 *
 * ### Required
 *
 * - **name** (String) name of the worker pool
 *
 * ### Optional
 *
 * - **csr** (String, Sensitive) certificate signing request in base64
 * - **description** (String) description of the worker pool
 * - **id** (String) The ID of this resource.
 * - **labels** (Set of String)
 *
 * ### Read-Only
 *
 * - **config** (String, Sensitive) credentials necessary to connect WorkerPool's workers to the control plane
 * - **private_key** (String, Sensitive) private key in base64
 *
 * ## Import
 *
 * Import is supported using the following syntax
 *
 * ```sh
 *  $ pulumi import spacelift:index/workerPool:WorkerPool k8s-core $WORKER_POOL_ID
 * ```
 */
export declare class WorkerPool extends pulumi.CustomResource {
    /**
     * Get an existing WorkerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkerPoolState, opts?: pulumi.CustomResourceOptions): WorkerPool;
    /**
     * Returns true if the given object is an instance of WorkerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is WorkerPool;
    /**
     * credentials necessary to connect WorkerPool's workers to the control plane
     */
    readonly config: pulumi.Output<string>;
    /**
     * certificate signing request in base64
     */
    readonly csr: pulumi.Output<string>;
    /**
     * description of the worker pool
     */
    readonly description: pulumi.Output<string | undefined>;
    readonly labels: pulumi.Output<string[] | undefined>;
    /**
     * name of the worker pool
     */
    readonly name: pulumi.Output<string>;
    /**
     * private key in base64
     */
    readonly privateKey: pulumi.Output<string>;
    /**
     * Create a WorkerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkerPoolArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering WorkerPool resources.
 */
export interface WorkerPoolState {
    /**
     * credentials necessary to connect WorkerPool's workers to the control plane
     */
    readonly config?: pulumi.Input<string>;
    /**
     * certificate signing request in base64
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * description of the worker pool
     */
    readonly description?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * name of the worker pool
     */
    readonly name?: pulumi.Input<string>;
    /**
     * private key in base64
     */
    readonly privateKey?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a WorkerPool resource.
 */
export interface WorkerPoolArgs {
    /**
     * certificate signing request in base64
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * description of the worker pool
     */
    readonly description?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * name of the worker pool
     */
    readonly name: pulumi.Input<string>;
}
