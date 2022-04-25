import * as pulumi from "@pulumi/pulumi";
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
export declare function getGitHubEnterpriseIntegration(opts?: pulumi.InvokeOptions): Promise<GetGitHubEnterpriseIntegrationResult>;
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
