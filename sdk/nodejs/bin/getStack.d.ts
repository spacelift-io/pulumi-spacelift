import * as pulumi from "@pulumi/pulumi";
import { output as outputs } from "./types";
export declare function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult>;
/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackArgs {
    readonly stackId: string;
}
/**
 * A collection of values returned by getStack.
 */
export interface GetStackResult {
    readonly administrative: boolean;
    readonly autodeploy: boolean;
    readonly autoretry: boolean;
    readonly awsAssumeRolePolicyStatement: string;
    readonly beforeInits: string[];
    readonly branch: string;
    readonly cloudformations: outputs.GetStackCloudformation[];
    readonly description: string;
    readonly gitlabs: outputs.GetStackGitlab[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: string[];
    readonly manageState: boolean;
    readonly name: string;
    readonly projectRoot: string;
    readonly pulumis: outputs.GetStackPulumi[];
    readonly repository: string;
    readonly runnerImage: string;
    readonly stackId: string;
    readonly terraformVersion: string;
    readonly terraformWorkspace: string;
    readonly workerPoolId: string;
}
