// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.ScheduledDeleteTask` represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const ireland-kubeconfig-delete = spacelift.getScheduledDeleteStack({
 *     scheduledDeleteStackId: "$STACK_ID/$SCHEDULED_DELETE_STACK_ID",
 * });
 * ```
 */
export function getScheduledDeleteStack(args: GetScheduledDeleteStackArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduledDeleteStackResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getScheduledDeleteStack:getScheduledDeleteStack", {
        "scheduledDeleteStackId": args.scheduledDeleteStackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getScheduledDeleteStack.
 */
export interface GetScheduledDeleteStackArgs {
    /**
     * ID of the scheduled delete*stack (stack*id/schedule_id)
     */
    scheduledDeleteStackId: string;
}

/**
 * A collection of values returned by getScheduledDeleteStack.
 */
export interface GetScheduledDeleteStackResult {
    /**
     * Timestamp (unix timestamp) at which time the scheduling should happen.
     */
    readonly at: number;
    /**
     * Indicates whether the resources of the stack should be deleted.
     */
    readonly deleteResources: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ID of the schedule
     */
    readonly scheduleId: string;
    /**
     * ID of the scheduled delete*stack (stack*id/schedule_id)
     */
    readonly scheduledDeleteStackId: string;
    /**
     * Stack ID of the scheduling config
     */
    readonly stackId: string;
}
/**
 * `spacelift.ScheduledDeleteTask` represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const ireland-kubeconfig-delete = spacelift.getScheduledDeleteStack({
 *     scheduledDeleteStackId: "$STACK_ID/$SCHEDULED_DELETE_STACK_ID",
 * });
 * ```
 */
export function getScheduledDeleteStackOutput(args: GetScheduledDeleteStackOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScheduledDeleteStackResult> {
    return pulumi.output(args).apply((a: any) => getScheduledDeleteStack(a, opts))
}

/**
 * A collection of arguments for invoking getScheduledDeleteStack.
 */
export interface GetScheduledDeleteStackOutputArgs {
    /**
     * ID of the scheduled delete*stack (stack*id/schedule_id)
     */
    scheduledDeleteStackId: pulumi.Input<string>;
}
