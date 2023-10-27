// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.getCurrentStack` is a data source that provides information about the current administrative stack if the run is executed within Spacelift by a stack or module. This allows clever tricks like attaching contexts or policies to the stack that manages them.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const this = spacelift.getCurrentStack({});
 * const core_kubeconfig = new spacelift.EnvironmentVariable("core-kubeconfig", {
 *     stackId: _this.then(_this => _this.id),
 *     value: "bacon",
 * });
 * ```
 */
export function getCurrentStack(opts?: pulumi.InvokeOptions): Promise<GetCurrentStackResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getCurrentStack:getCurrentStack", {
    }, opts);
}

/**
 * A collection of values returned by getCurrentStack.
 */
export interface GetCurrentStackResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * `spacelift.getCurrentStack` is a data source that provides information about the current administrative stack if the run is executed within Spacelift by a stack or module. This allows clever tricks like attaching contexts or policies to the stack that manages them.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const this = spacelift.getCurrentStack({});
 * const core_kubeconfig = new spacelift.EnvironmentVariable("core-kubeconfig", {
 *     stackId: _this.then(_this => _this.id),
 *     value: "bacon",
 * });
 * ```
 */
export function getCurrentStackOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetCurrentStackResult> {
    return pulumi.output(getCurrentStack(opts))
}
