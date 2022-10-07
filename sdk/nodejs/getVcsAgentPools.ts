// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.getVcsAgentPools` represents the VCS agent pools assigned to the Spacelift account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const vcs_agent_pools = pulumi.output(spacelift.getVcsAgentPools());
 * ```
 */
export function getVcsAgentPools(opts?: pulumi.InvokeOptions): Promise<GetVcsAgentPoolsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getVcsAgentPools:getVcsAgentPools", {
    }, opts);
}

/**
 * A collection of values returned by getVcsAgentPools.
 */
export interface GetVcsAgentPoolsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vcsAgentPools: outputs.GetVcsAgentPoolsVcsAgentPool[];
}
