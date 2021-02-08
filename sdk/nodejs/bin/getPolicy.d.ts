import * as pulumi from "@pulumi/pulumi";
export declare function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult>;
/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    readonly policyId: string;
}
/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    readonly body: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly policyId: string;
    readonly type: string;
}
