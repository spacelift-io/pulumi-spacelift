// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getBitbucketCloudIntegration` returns details about Bitbucket Cloud integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const bitbucketCloudIntegration = pulumi.output(spacelift.getBitbucketCloudIntegration());
 * ```
 */
export function getBitbucketCloudIntegration(opts?: pulumi.InvokeOptions): Promise<GetBitbucketCloudIntegrationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getBitbucketCloudIntegration:getBitbucketCloudIntegration", {
    }, opts);
}

/**
 * A collection of values returned by getBitbucketCloudIntegration.
 */
export interface GetBitbucketCloudIntegrationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Bitbucket Cloud username
     */
    readonly username: string;
}
