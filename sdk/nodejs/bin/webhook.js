"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
class Webhook extends pulumi.CustomResource {
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
        }
        else {
            const args = argsOrState;
            if ((!args || args.endpoint === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'endpoint'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["secret"] = args ? args.secret : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
        }
        if (!opts) {
            opts = {};
        }
        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Webhook.__pulumiType, name, inputs, opts);
    }
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, state, opts) {
        return new Webhook(name, state, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }
}
exports.Webhook = Webhook;
/** @internal */
Webhook.__pulumiType = 'spacelift:index/webhook:Webhook';
//# sourceMappingURL=webhook.js.map