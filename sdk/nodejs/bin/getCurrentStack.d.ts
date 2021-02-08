import * as pulumi from "@pulumi/pulumi";
export declare function getCurrentStack(opts?: pulumi.InvokeOptions): Promise<GetCurrentStackResult>;
/**
 * A collection of values returned by getCurrentStack.
 */
export interface GetCurrentStackResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
