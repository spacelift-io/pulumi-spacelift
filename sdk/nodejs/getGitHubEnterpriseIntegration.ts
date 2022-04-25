// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

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
export function getGitHubEnterpriseIntegration(opts?: pulumi.InvokeOptions): Promise<GetGitHubEnterpriseIntegrationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getGitHubEnterpriseIntegration:getGitHubEnterpriseIntegration", {
    }, opts);
}

/**
 * A collection of values returned by getGitHubEnterpriseIntegration.
 */
export interface GetGitHubEnterpriseIntegrationResult {
    readonly apiHost: string;
    readonly appId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly webhookSecret: string;
}
