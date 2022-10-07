// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getAzureDevopsIntegration` returns details about Azure DevOps integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const azureDevopsIntegration = pulumi.output(spacelift.getAzureDevopsIntegration());
 * ```
 */
export function getAzureDevopsIntegration(opts?: pulumi.InvokeOptions): Promise<GetAzureDevopsIntegrationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getAzureDevopsIntegration:getAzureDevopsIntegration", {
    }, opts);
}

/**
 * A collection of values returned by getAzureDevopsIntegration.
 */
export interface GetAzureDevopsIntegrationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationUrl: string;
    readonly webhookPassword: string;
}