import * as pulumi from "@pulumi/pulumi";
import { output as outputs } from "./types";
export declare function getWorkerPools(opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolsResult>;
/**
 * A collection of values returned by getWorkerPools.
 */
export interface GetWorkerPoolsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly workerPools: outputs.GetWorkerPoolsWorkerPool[];
}
