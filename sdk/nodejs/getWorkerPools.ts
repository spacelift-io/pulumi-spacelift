// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.getWorkerPools` represents the worker pools assigned to the Spacelift account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const worker_pools = pulumi.output(spacelift.getWorkerPools({ async: true }));
 * ```
 *
 * <!-- schema generated by tfplugindocs -->
 * ## Schema
 *
 * ### Optional
 *
 * - **id** (String) The ID of this resource.
 *
 * ### Read-Only
 *
 * - **worker_pools** (List of Object) (see below for nested schema)
 *
 * <a id="nestedatt--worker_pools"></a>
 * ### Nested Schema for `workerPools`
 *
 * Read-Only:
 *
 * - **config** (String)
 * - **description** (String)
 * - **name** (String)
 * - **worker_pool_id** (String)
 */
export function getWorkerPools(opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getWorkerPools:getWorkerPools", {
    }, opts);
}

/**
 * A collection of values returned by getWorkerPools.
 */
export interface GetWorkerPoolsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly workerPools: outputs.GetWorkerPoolsWorkerPool[];
}
