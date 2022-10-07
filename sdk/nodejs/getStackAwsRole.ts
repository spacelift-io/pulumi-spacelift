// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > **Note:** `spacelift.StackAwsRole` is deprecated. Please use `spacelift.AwsRole` instead. The functionality is identical.
 *
 * `spacelift.StackAwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.
 *
 * If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).
 *
 * Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * // For a Module
 * const k8s_module = pulumi.output(spacelift.getStackAwsRole({
 *     moduleId: "k8s-module",
 * }));
 * // For a Stack
 * const k8s_core = pulumi.output(spacelift.getStackAwsRole({
 *     stackId: "k8s-core",
 * }));
 * ```
 */
export function getStackAwsRole(args?: GetStackAwsRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetStackAwsRoleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getStackAwsRole:getStackAwsRole", {
        "moduleId": args.moduleId,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getStackAwsRole.
 */
export interface GetStackAwsRoleArgs {
    moduleId?: string;
    stackId?: string;
}

/**
 * A collection of values returned by getStackAwsRole.
 */
export interface GetStackAwsRoleResult {
    readonly durationSeconds: number;
    readonly externalId: string;
    readonly generateCredentialsInWorker: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly roleArn: string;
    readonly stackId?: string;
}

export function getStackAwsRoleOutput(args?: GetStackAwsRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStackAwsRoleResult> {
    return pulumi.output(args).apply(a => getStackAwsRole(a, opts))
}

/**
 * A collection of arguments for invoking getStackAwsRole.
 */
export interface GetStackAwsRoleOutputArgs {
    moduleId?: pulumi.Input<string>;
    stackId?: pulumi.Input<string>;
}
