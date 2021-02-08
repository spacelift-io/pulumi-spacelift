import * as pulumi from "@pulumi/pulumi";
export declare function getWorkerPool(args: GetWorkerPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolResult>;
/**
 * A collection of arguments for invoking getWorkerPool.
 */
export interface GetWorkerPoolArgs {
    readonly workerPoolId: string;
}
/**
 * A collection of values returned by getWorkerPool.
 */
export interface GetWorkerPoolResult {
    readonly config: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workerPoolId: string;
}
