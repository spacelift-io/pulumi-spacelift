"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
class PolicyAttachment extends pulumi.CustomResource {
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["customInput"] = state ? state.customInput : undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
        }
        else {
            const args = argsOrState;
            if ((!args || args.policyId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyId'");
            }
            inputs["customInput"] = args ? args.customInput : undefined;
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
        }
        if (!opts) {
            opts = {};
        }
        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PolicyAttachment.__pulumiType, name, inputs, opts);
    }
    /**
     * Get an existing PolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, state, opts) {
        return new PolicyAttachment(name, state, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of PolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAttachment.__pulumiType;
    }
}
exports.PolicyAttachment = PolicyAttachment;
/** @internal */
PolicyAttachment.__pulumiType = 'spacelift:index/policyAttachment:PolicyAttachment';
//# sourceMappingURL=policyAttachment.js.map