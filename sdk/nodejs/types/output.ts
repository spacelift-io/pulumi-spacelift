// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface GetModuleAzureDevop {
    project: string;
}

export interface GetModuleBitbucketCloud {
    namespace: string;
}

export interface GetModuleBitbucketDatacenter {
    namespace: string;
}

export interface GetModuleGithubEnterprise {
    namespace: string;
}

export interface GetModuleGitlab {
    namespace: string;
}

export interface GetPoliciesPolicy {
    id: string;
    labels: string[];
    name: string;
    spaceId: string;
    type: string;
}

export interface GetStackAnsible {
    playbook: string;
}

export interface GetStackAzureDevop {
    project: string;
}

export interface GetStackBitbucketCloud {
    namespace: string;
}

export interface GetStackBitbucketDatacenter {
    namespace: string;
}

export interface GetStackCloudformation {
    entryTemplateFile: string;
    region: string;
    stackName: string;
    templateBucket: string;
}

export interface GetStackGithubEnterprise {
    namespace: string;
}

export interface GetStackGitlab {
    namespace: string;
}

export interface GetStackKubernete {
    namespace: string;
}

export interface GetStackPulumi {
    loginUrl: string;
    stackName: string;
}

export interface GetStackShowcase {
    namespace: string;
}

export interface GetVcsAgentPoolsVcsAgentPool {
    description: string;
    name: string;
    vcsAgentPoolId: string;
}

export interface GetWorkerPoolsWorkerPool {
    config: string;
    description: string;
    name: string;
    spaceId: string;
    workerPoolId: string;
}

export interface ModuleAzureDevops {
    project: string;
}

export interface ModuleBitbucketCloud {
    namespace: string;
}

export interface ModuleBitbucketDatacenter {
    namespace: string;
}

export interface ModuleGithubEnterprise {
    namespace: string;
}

export interface ModuleGitlab {
    namespace: string;
}

export interface StackAnsible {
    playbook: string;
}

export interface StackAzureDevops {
    project: string;
}

export interface StackBitbucketCloud {
    namespace: string;
}

export interface StackBitbucketDatacenter {
    namespace: string;
}

export interface StackCloudformation {
    entryTemplateFile: string;
    region: string;
    stackName: string;
    templateBucket: string;
}

export interface StackGithubEnterprise {
    namespace: string;
}

export interface StackGitlab {
    namespace: string;
}

export interface StackKubernetes {
    namespace?: string;
}

export interface StackPulumi {
    loginUrl: string;
    stackName: string;
}

export interface StackShowcase {
    namespace: string;
}
