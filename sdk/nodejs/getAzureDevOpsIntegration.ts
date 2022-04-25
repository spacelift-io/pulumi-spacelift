// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.getAzureDevOpsIntegration` returns details about Azure DevOps integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const azureDevopsIntegration = pulumi.output(spacelift.getAzureDevOpsIntegration({ async: true }));
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
 * - **organization_url** (String) Azure DevOps integration organization url
 * - **webhook_password** (String) Azure DevOps integration webhook password
 */
export function getAzureDevOpsIntegration(opts?: pulumi.InvokeOptions): Promise<GetAzureDevOpsIntegrationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getAzureDevOpsIntegration:getAzureDevOpsIntegration", {
    }, opts);
}

/**
 * A collection of values returned by getAzureDevOpsIntegration.
 */
export interface GetAzureDevOpsIntegrationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationUrl: string;
    readonly webhookPassword: string;
}