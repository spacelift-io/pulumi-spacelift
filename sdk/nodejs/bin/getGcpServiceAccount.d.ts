import * as pulumi from "@pulumi/pulumi";
export declare function getGcpServiceAccount(args?: GetGcpServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetGcpServiceAccountResult>;
/**
 * A collection of arguments for invoking getGcpServiceAccount.
 */
export interface GetGcpServiceAccountArgs {
    readonly moduleId?: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getGcpServiceAccount.
 */
export interface GetGcpServiceAccountResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly serviceAccountEmail: string;
    readonly stackId?: string;
    readonly tokenScopes: string[];
}
