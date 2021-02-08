import * as pulumi from "@pulumi/pulumi";
export declare function getEnvironmentVariable(args: GetEnvironmentVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentVariableResult>;
/**
 * A collection of arguments for invoking getEnvironmentVariable.
 */
export interface GetEnvironmentVariableArgs {
    readonly contextId?: string;
    readonly moduleId?: string;
    readonly name: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getEnvironmentVariable.
 */
export interface GetEnvironmentVariableResult {
    readonly checksum: string;
    readonly contextId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly name: string;
    readonly stackId?: string;
    readonly value: string;
    readonly writeOnly: boolean;
}
