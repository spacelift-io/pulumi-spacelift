export interface GetModuleGitlab {
    namespace: string;
}
export interface GetStackCloudformation {
    entryTemplateFile: string;
    region: string;
    stackName: string;
    templateBucket: string;
}
export interface GetStackGitlab {
    namespace: string;
}
export interface GetStackPulumi {
    loginUrl: string;
    stackName: string;
}
export interface GetWorkerPoolsWorkerPool {
    config: string;
    description: string;
    name: string;
    workerPoolId: string;
}
export interface ModuleGitlab {
    namespace: string;
}
export interface StackCloudformation {
    entryTemplateFile: string;
    region: string;
    stackName: string;
    templateBucket: string;
}
export interface StackGitlab {
    namespace: string;
}
export interface StackPulumi {
    loginUrl: string;
    stackName: string;
}
