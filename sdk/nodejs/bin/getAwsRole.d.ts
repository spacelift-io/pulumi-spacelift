import * as pulumi from "@pulumi/pulumi";
export declare function getAwsRole(args?: GetAwsRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetAwsRoleResult>;
/**
 * A collection of arguments for invoking getAwsRole.
 */
export interface GetAwsRoleArgs {
    readonly moduleId?: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getAwsRole.
 */
export interface GetAwsRoleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly roleArn: string;
    readonly stackId?: string;
}
