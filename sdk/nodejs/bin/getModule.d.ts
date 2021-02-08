import * as pulumi from "@pulumi/pulumi";
import { output as outputs } from "./types";
export declare function getModule(args: GetModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetModuleResult>;
/**
 * A collection of arguments for invoking getModule.
 */
export interface GetModuleArgs {
    readonly moduleId: string;
}
/**
 * A collection of values returned by getModule.
 */
export interface GetModuleResult {
    readonly administrative: boolean;
    readonly awsAssumeRolePolicyStatement: string;
    readonly branch: string;
    readonly description: string;
    readonly gitlabs: outputs.GetModuleGitlab[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: string[];
    readonly moduleId: string;
    readonly repository: string;
    readonly sharedAccounts: string[];
    readonly workerPoolId: string;
}
