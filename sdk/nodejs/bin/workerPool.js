"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
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
class WorkerPool extends pulumi.CustomResource {
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["config"] = state ? state.config : undefined;
            inputs["csr"] = state ? state.csr : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
        }
        else {
            const args = argsOrState;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            inputs["csr"] = args ? args.csr : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["config"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {};
        }
        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WorkerPool.__pulumiType, name, inputs, opts);
    }
    /**
     * Get an existing WorkerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, state, opts) {
        return new WorkerPool(name, state, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of WorkerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkerPool.__pulumiType;
    }
}
exports.WorkerPool = WorkerPool;
/** @internal */
WorkerPool.__pulumiType = 'spacelift:index/workerPool:WorkerPool';
//# sourceMappingURL=workerPool.js.map