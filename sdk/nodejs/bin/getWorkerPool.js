"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
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
 * }, { async: true }));
 * ```
 *
 * <!-- schema generated by tfplugindocs -->
 * ## Schema
 *
 * ### Required
 *
 * - **worker_pool_id** (String) ID of the worker pool
 *
 * ### Optional
 *
 * - **id** (String) The ID of this resource.
 *
 * ### Read-Only
 *
 * - **config** (String, Sensitive) credentials necessary to connect WorkerPool's workers to the control plane
 * - **description** (String) description of the worker pool
 * - **labels** (Set of String)
 * - **name** (String) name of the worker pool
 */
function getWorkerPool(args, opts) {
    if (!opts) {
        opts = {};
    }
    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getWorkerPool:getWorkerPool", {
        "workerPoolId": args.workerPoolId,
    }, opts);
}
exports.getWorkerPool = getWorkerPool;
//# sourceMappingURL=getWorkerPool.js.map