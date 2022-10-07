// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getIPs` returns the list of Spacelift's outgoing IP addresses, which you can use to whitelist connections coming from the Spacelift's "mothership".
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const ips = pulumi.output(spacelift.getIPs());
 * ```
 */
export function getIPs(opts?: pulumi.InvokeOptions): Promise<GetIPsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getIPs:getIPs", {
    }, opts);
}

/**
 * A collection of values returned by getIPs.
 */
export interface GetIPsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ips: string[];
}