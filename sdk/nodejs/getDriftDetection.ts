// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const core_infra_production_drift_detection = pulumi.output(spacelift.getDriftDetection({
 *     stackId: "core-infra-production",
 * }));
 * ```
 */
export function getDriftDetection(args: GetDriftDetectionArgs, opts?: pulumi.InvokeOptions): Promise<GetDriftDetectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getDriftDetection:getDriftDetection", {
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDriftDetection.
 */
export interface GetDriftDetectionArgs {
    stackId: string;
}

/**
 * A collection of values returned by getDriftDetection.
 */
export interface GetDriftDetectionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly reconcile: boolean;
    readonly schedules: string[];
    readonly stackId: string;
    readonly timezone: string;
}

export function getDriftDetectionOutput(args: GetDriftDetectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDriftDetectionResult> {
    return pulumi.output(args).apply(a => getDriftDetection(a, opts))
}

/**
 * A collection of arguments for invoking getDriftDetection.
 */
export interface GetDriftDetectionOutputArgs {
    stackId: pulumi.Input<string>;
}
