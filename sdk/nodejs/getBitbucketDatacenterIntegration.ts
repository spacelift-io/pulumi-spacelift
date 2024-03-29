// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getBitbucketDatacenterIntegration` returns details about Bitbucket Datacenter integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const bitbucketDatacenterIntegration = spacelift.getBitbucketDatacenterIntegration({});
 * ```
 */
export function getBitbucketDatacenterIntegration(args?: GetBitbucketDatacenterIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetBitbucketDatacenterIntegrationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getBitbucketDatacenterIntegration:getBitbucketDatacenterIntegration", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getBitbucketDatacenterIntegration.
 */
export interface GetBitbucketDatacenterIntegrationArgs {
    /**
     * Bitbucket Datacenter integration id. If not provided, the default integration will be returned
     */
    id?: string;
}

/**
 * A collection of values returned by getBitbucketDatacenterIntegration.
 */
export interface GetBitbucketDatacenterIntegrationResult {
    /**
     * Bitbucket Datacenter integration api host
     */
    readonly apiHost: string;
    /**
     * Bitbucket Datacenter integration description
     */
    readonly description: string;
    /**
     * Bitbucket Datacenter integration id. If not provided, the default integration will be returned
     */
    readonly id?: string;
    /**
     * Bitbucket Datacenter integration is default
     */
    readonly isDefault: boolean;
    /**
     * Bitbucket Datacenter integration labels
     */
    readonly labels: string[];
    /**
     * Bitbucket Datacenter integration name
     */
    readonly name: string;
    /**
     * Bitbucket Datacenter integration space id
     */
    readonly spaceId: string;
    /**
     * Bitbucket Datacenter integration user facing host
     */
    readonly userFacingHost: string;
    /**
     * Bitbucket Datacenter username
     */
    readonly username: string;
    /**
     * Bitbucket Datacenter integration webhook secret
     */
    readonly webhookSecret: string;
    /**
     * Bitbucket Datacenter integration webhook URL
     */
    readonly webhookUrl: string;
}
/**
 * `spacelift.getBitbucketDatacenterIntegration` returns details about Bitbucket Datacenter integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const bitbucketDatacenterIntegration = spacelift.getBitbucketDatacenterIntegration({});
 * ```
 */
export function getBitbucketDatacenterIntegrationOutput(args?: GetBitbucketDatacenterIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBitbucketDatacenterIntegrationResult> {
    return pulumi.output(args).apply((a: any) => getBitbucketDatacenterIntegration(a, opts))
}

/**
 * A collection of arguments for invoking getBitbucketDatacenterIntegration.
 */
export interface GetBitbucketDatacenterIntegrationOutputArgs {
    /**
     * Bitbucket Datacenter integration id. If not provided, the default integration will be returned
     */
    id?: pulumi.Input<string>;
}
