"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
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
function getWorkerPools(opts) {
    if (!opts) {
        opts = {};
    }
    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getWorkerPools:getWorkerPools", {}, opts);
}
exports.getWorkerPools = getWorkerPools;
//# sourceMappingURL=getWorkerPools.js.map