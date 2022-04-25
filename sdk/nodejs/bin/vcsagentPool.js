"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
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
class VCSAgentPool extends pulumi.CustomResource {
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["config"] = state ? state.config : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
        }
        else {
            const args = argsOrState;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["config"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {};
        }
        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VCSAgentPool.__pulumiType, name, inputs, opts);
    }
    /**
     * Get an existing VCSAgentPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, state, opts) {
        return new VCSAgentPool(name, state, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of VCSAgentPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VCSAgentPool.__pulumiType;
    }
}
exports.VCSAgentPool = VCSAgentPool;
/** @internal */
VCSAgentPool.__pulumiType = 'spacelift:index/vCSAgentPool:VCSAgentPool';
//# sourceMappingURL=vcsagentPool.js.map