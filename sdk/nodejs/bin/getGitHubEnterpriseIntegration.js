"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
/**
 * `spacelift.getGitHubEnterpriseIntegration` returns details about Github Enterprise integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const githubEnterpriseIntegration = pulumi.output(spacelift.getGitHubEnterpriseIntegration({ async: true }));
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
 * - **api_host** (String) Github integration api host
 * - **app_id** (String) Github integration app id
 * - **webhook_secret** (String) Github integration webhook secret
 */
function getGitHubEnterpriseIntegration(opts) {
    if (!opts) {
        opts = {};
    }
    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getGitHubEnterpriseIntegration:getGitHubEnterpriseIntegration", {}, opts);
}
exports.getGitHubEnterpriseIntegration = getGitHubEnterpriseIntegration;
//# sourceMappingURL=getGitHubEnterpriseIntegration.js.map