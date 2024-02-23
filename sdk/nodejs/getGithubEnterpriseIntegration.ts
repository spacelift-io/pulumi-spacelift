// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getGithubEnterpriseIntegration` returns details about Github Enterprise integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const githubEnterpriseIntegration = spacelift.getGithubEnterpriseIntegration({});
 * ```
 */
export function getGithubEnterpriseIntegration(args?: GetGithubEnterpriseIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetGithubEnterpriseIntegrationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getGithubEnterpriseIntegration:getGithubEnterpriseIntegration", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getGithubEnterpriseIntegration.
 */
export interface GetGithubEnterpriseIntegrationArgs {
    /**
     * Github integration id. If not provided, the default integration will be returned
     */
    id?: string;
}

/**
 * A collection of values returned by getGithubEnterpriseIntegration.
 */
export interface GetGithubEnterpriseIntegrationResult {
    /**
     * Github integration api host
     */
    readonly apiHost: string;
    /**
     * Github integration app id
     */
    readonly appId: string;
    /**
     * Github integration description
     */
    readonly description: string;
    /**
     * Github integration id. If not provided, the default integration will be returned
     */
    readonly id?: string;
    /**
     * Github integration is default
     */
    readonly isDefault: boolean;
    /**
     * Github integration labels
     */
    readonly labels: string[];
    /**
     * Github integration name
     */
    readonly name: string;
    /**
     * Github integration space id
     */
    readonly spaceId: string;
    /**
     * Github integration webhook secret
     */
    readonly webhookSecret: string;
    /**
     * Github integration webhook url
     */
    readonly webhookUrl: string;
}
/**
 * `spacelift.getGithubEnterpriseIntegration` returns details about Github Enterprise integration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const githubEnterpriseIntegration = spacelift.getGithubEnterpriseIntegration({});
 * ```
 */
export function getGithubEnterpriseIntegrationOutput(args?: GetGithubEnterpriseIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGithubEnterpriseIntegrationResult> {
    return pulumi.output(args).apply((a: any) => getGithubEnterpriseIntegration(a, opts))
}

/**
 * A collection of arguments for invoking getGithubEnterpriseIntegration.
 */
export interface GetGithubEnterpriseIntegrationOutputArgs {
    /**
     * Github integration id. If not provided, the default integration will be returned
     */
    id?: pulumi.Input<string>;
}
