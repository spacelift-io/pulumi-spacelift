"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
function getStackAwsRole(args, opts) {
    args = args || {};
    if (!opts) {
        opts = {};
    }
    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getStackAwsRole:getStackAwsRole", {
        "moduleId": args.moduleId,
        "stackId": args.stackId,
    }, opts);
}
exports.getStackAwsRole = getStackAwsRole;
//# sourceMappingURL=getStackAwsRole.js.map