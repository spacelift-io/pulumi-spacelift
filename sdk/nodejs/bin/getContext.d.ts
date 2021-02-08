import * as pulumi from "@pulumi/pulumi";
export declare function getContext(args: GetContextArgs, opts?: pulumi.InvokeOptions): Promise<GetContextResult>;
/**
 * A collection of arguments for invoking getContext.
 */
export interface GetContextArgs {
    readonly contextId: string;
}
/**
 * A collection of values returned by getContext.
 */
export interface GetContextResult {
    readonly contextId: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
