import * as pulumi from "@pulumi/pulumi";
import { output as outputs } from "./types";
export declare function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult>;
/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackArgs {
    readonly afterApplies?: string[];
    readonly afterDestroys?: string[];
    readonly afterInits?: string[];
    readonly afterPerforms?: string[];
    readonly afterPlans?: string[];
    readonly beforeApplies?: string[];
    readonly beforeDestroys?: string[];
    readonly beforeInits?: string[];
    readonly beforePerforms?: string[];
    readonly beforePlans?: string[];
    readonly stackId: string;
}
/**
 * A collection of values returned by getStack.
 */
export interface GetStackResult {
    readonly administrative: boolean;
    readonly afterApplies?: string[];
    readonly afterDestroys?: string[];
    readonly afterInits?: string[];
    readonly afterPerforms?: string[];
    readonly afterPlans?: string[];
    readonly autodeploy: boolean;
    readonly autoretry: boolean;
    readonly awsAssumeRolePolicyStatement: string;
    readonly azureDevops: outputs.GetStackAzureDevop[];
    readonly beforeApplies?: string[];
    readonly beforeDestroys?: string[];
    readonly beforeInits?: string[];
    readonly beforePerforms?: string[];
    readonly beforePlans?: string[];
    readonly bitbucketClouds: outputs.GetStackBitbucketCloud[];
    readonly bitbucketDatacenters: outputs.GetStackBitbucketDatacenter[];
    readonly branch: string;
    readonly cloudformations: outputs.GetStackCloudformation[];
    readonly description: string;
    readonly enableLocalPreview: boolean;
    readonly githubEnterprises: outputs.GetStackGithubEnterprise[];
    readonly gitlabs: outputs.GetStackGitlab[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kubernetes: outputs.GetStackKubernete[];
    readonly labels: string[];
    readonly manageState: boolean;
    readonly name: string;
    readonly projectRoot: string;
    readonly protectFromDeletion: boolean;
    readonly pulumis: outputs.GetStackPulumi[];
    readonly repository: string;
    readonly runnerImage: string;
    readonly showcases: outputs.GetStackShowcase[];
    readonly stackId: string;
    readonly terraformVersion: string;
    readonly terraformWorkspace: string;
    readonly workerPoolId: string;
}
