// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getGitlabWebhookEndpoint` returns details about Gitlab webhook endpoint
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const this = spacelift.getGitlabWebhookEndpoint({});
 * ```
 */
export function getGitlabWebhookEndpoint(opts?: pulumi.InvokeOptions): Promise<GetGitlabWebhookEndpointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getGitlabWebhookEndpoint:getGitlabWebhookEndpoint", {
    }, opts);
}

/**
 * A collection of values returned by getGitlabWebhookEndpoint.
 */
export interface GetGitlabWebhookEndpointResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Gitlab webhook endpoint address
     */
    readonly webhookEndpoint: string;
}
/**
 * `spacelift.getGitlabWebhookEndpoint` returns details about Gitlab webhook endpoint
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const this = spacelift.getGitlabWebhookEndpoint({});
 * ```
 */
export function getGitlabWebhookEndpointOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetGitlabWebhookEndpointResult> {
    return pulumi.output(getGitlabWebhookEndpoint(opts))
}