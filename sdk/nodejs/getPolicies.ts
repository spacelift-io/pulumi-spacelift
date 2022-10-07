// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.getPolicies` can find all policies that have certain labels.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const all = spacelift.getPolicies({});
 * const planAutoattach = spacelift.getPolicies({
 *     type: "PLAN",
 *     labels: ["autoattach"],
 * });
 * export const policyIds = data.spacelift_policies["this"].policies.map(__item => __item.id);
 * ```
 */
export function getPolicies(args?: GetPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetPoliciesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getPolicies:getPolicies", {
        "labels": args.labels,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicies.
 */
export interface GetPoliciesArgs {
    labels?: string[];
    type?: string;
}

/**
 * A collection of values returned by getPolicies.
 */
export interface GetPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels?: string[];
    readonly policies: outputs.GetPoliciesPolicy[];
    readonly type?: string;
}

export function getPoliciesOutput(args?: GetPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPoliciesResult> {
    return pulumi.output(args).apply(a => getPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getPolicies.
 */
export interface GetPoliciesOutputArgs {
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
}