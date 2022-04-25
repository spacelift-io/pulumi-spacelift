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
export interface GetWorkerPoolsWorkerPool {
    config: string;
    description: string;
    name: string;
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
