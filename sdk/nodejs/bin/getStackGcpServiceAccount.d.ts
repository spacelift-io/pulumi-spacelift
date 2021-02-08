import * as pulumi from "@pulumi/pulumi";
export declare function getStackGcpServiceAccount(args?: GetStackGcpServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetStackGcpServiceAccountResult>;
/**
 * A collection of arguments for invoking getStackGcpServiceAccount.
 */
export interface GetStackGcpServiceAccountArgs {
    readonly moduleId?: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getStackGcpServiceAccount.
 */
export interface GetStackGcpServiceAccountResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly serviceAccountEmail: string;
    readonly stackId?: string;
    readonly tokenScopes: string[];
}
