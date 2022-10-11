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
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const k8s_core = pulumi.output(spacelift.getWorkerPool({
 *     workerPoolId: "k8s-core",
 * }));
 * ```
 */
export function getWorkerPool(args: GetWorkerPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getWorkerPool:getWorkerPool", {
        "workerPoolId": args.workerPoolId,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkerPool.
 */
export interface GetWorkerPoolArgs {
    /**
     * ID of the worker pool
     */
    workerPoolId: string;
}

/**
 * A collection of values returned by getWorkerPool.
 */
export interface GetWorkerPoolResult {
    /**
     * credentials necessary to connect WorkerPool's workers to the control plane
     */
    readonly config: string;
    /**
     * description of the worker pool
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: string[];
    /**
     * name of the worker pool
     */
    readonly name: string;
    /**
     * ID (slug) of the space the worker pool is in
     */
    readonly spaceId: string;
    /**
     * ID of the worker pool
     */
    readonly workerPoolId: string;
}

export function getWorkerPoolOutput(args: GetWorkerPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkerPoolResult> {
    return pulumi.output(args).apply(a => getWorkerPool(a, opts))
}

/**
 * A collection of arguments for invoking getWorkerPool.
 */
export interface GetWorkerPoolOutputArgs {
    /**
     * ID of the worker pool
     */
    workerPoolId: pulumi.Input<string>;
}
