import * as pulumi from "@pulumi/pulumi";
export declare function getIps(opts?: pulumi.InvokeOptions): Promise<GetIpsResult>;
/**
 * A collection of values returned by getIps.
 */
export interface GetIpsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ips: string[];
}
