import * as pulumi from "@pulumi/pulumi";
export declare function getStackAwsRole(args?: GetStackAwsRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetStackAwsRoleResult>;
/**
 * A collection of arguments for invoking getStackAwsRole.
 */
export interface GetStackAwsRoleArgs {
    readonly moduleId?: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getStackAwsRole.
 */
export interface GetStackAwsRoleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly roleArn: string;
    readonly stackId?: string;
}
