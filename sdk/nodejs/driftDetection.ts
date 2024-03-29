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
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const core_infra_production = new spacelift.Stack("core-infra-production", {
 *     branch: "master",
 *     repository: "core-infra",
 * });
 * const core_infra_production_drift_detection = new spacelift.DriftDetection("core-infra-production-drift-detection", {
 *     reconcile: true,
 *     stackId: core_infra_production.id,
 *     schedules: ["*&#47;15 * * * *"],
 * });
 * // Every 15 minutes
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import spacelift:index/driftDetection:DriftDetection core-infra-production-drift-detection stack/$STACK_ID
 * ```
 *
 * ```sh
 *  $ pulumi import spacelift:index/driftDetection:DriftDetection core-infra-production-drift-detection module/$MODULE_ID
 * ```
 */
export class DriftDetection extends pulumi.CustomResource {
    /**
     * Get an existing DriftDetection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DriftDetectionState, opts?: pulumi.CustomResourceOptions): DriftDetection {
        return new DriftDetection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/driftDetection:DriftDetection';

    /**
     * Returns true if the given object is an instance of DriftDetection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DriftDetection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DriftDetection.__pulumiType;
    }

    /**
     * Controls whether drift detection should be performed on a stack in any final state instead of just 'Finished'.
     */
    public readonly ignoreState!: pulumi.Output<boolean | undefined>;
    /**
     * Whether a tracked run should be triggered when drift is detected.
     */
    public readonly reconcile!: pulumi.Output<boolean | undefined>;
    /**
     * List of cron schedule expressions based on which drift detection should be triggered.
     */
    public readonly schedules!: pulumi.Output<string[]>;
    /**
     * ID of the stack for which to set up drift detection
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * Timezone in which the schedule is expressed. Defaults to `UTC`.
     */
    public readonly timezone!: pulumi.Output<string | undefined>;

    /**
     * Create a DriftDetection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DriftDetectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DriftDetectionArgs | DriftDetectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DriftDetectionState | undefined;
            resourceInputs["ignoreState"] = state ? state.ignoreState : undefined;
            resourceInputs["reconcile"] = state ? state.reconcile : undefined;
            resourceInputs["schedules"] = state ? state.schedules : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
            resourceInputs["timezone"] = state ? state.timezone : undefined;
        } else {
            const args = argsOrState as DriftDetectionArgs | undefined;
            if ((!args || args.schedules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedules'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            resourceInputs["ignoreState"] = args ? args.ignoreState : undefined;
            resourceInputs["reconcile"] = args ? args.reconcile : undefined;
            resourceInputs["schedules"] = args ? args.schedules : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["timezone"] = args ? args.timezone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DriftDetection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DriftDetection resources.
 */
export interface DriftDetectionState {
    /**
     * Controls whether drift detection should be performed on a stack in any final state instead of just 'Finished'.
     */
    ignoreState?: pulumi.Input<boolean>;
    /**
     * Whether a tracked run should be triggered when drift is detected.
     */
    reconcile?: pulumi.Input<boolean>;
    /**
     * List of cron schedule expressions based on which drift detection should be triggered.
     */
    schedules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the stack for which to set up drift detection
     */
    stackId?: pulumi.Input<string>;
    /**
     * Timezone in which the schedule is expressed. Defaults to `UTC`.
     */
    timezone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DriftDetection resource.
 */
export interface DriftDetectionArgs {
    /**
     * Controls whether drift detection should be performed on a stack in any final state instead of just 'Finished'.
     */
    ignoreState?: pulumi.Input<boolean>;
    /**
     * Whether a tracked run should be triggered when drift is detected.
     */
    reconcile?: pulumi.Input<boolean>;
    /**
     * List of cron schedule expressions based on which drift detection should be triggered.
     */
    schedules: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the stack for which to set up drift detection
     */
    stackId: pulumi.Input<string>;
    /**
     * Timezone in which the schedule is expressed. Defaults to `UTC`.
     */
    timezone?: pulumi.Input<string>;
}
