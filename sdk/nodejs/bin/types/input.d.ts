import * as pulumi from "@pulumi/pulumi";
export interface ModuleAzureDevops {
    project: pulumi.Input<string>;
}
export interface ModuleBitbucketCloud {
    namespace: pulumi.Input<string>;
}
export interface ModuleBitbucketDatacenter {
    namespace: pulumi.Input<string>;
}
export interface ModuleGithubEnterprise {
    namespace: pulumi.Input<string>;
}
export interface ModuleGitlab {
    namespace: pulumi.Input<string>;
}
export interface StackAzureDevops {
    project: pulumi.Input<string>;
}
export interface StackBitbucketCloud {
    namespace: pulumi.Input<string>;
}
export interface StackBitbucketDatacenter {
    namespace: pulumi.Input<string>;
}
export interface StackCloudformation {
    entryTemplateFile: pulumi.Input<string>;
    region: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
    templateBucket: pulumi.Input<string>;
}
export interface StackGithubEnterprise {
    namespace: pulumi.Input<string>;
}
export interface StackGitlab {
    namespace: pulumi.Input<string>;
}
export interface StackKubernetes {
    namespace?: pulumi.Input<string>;
}
export interface StackPulumi {
    loginUrl: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
}
export interface StackShowcase {
    namespace: pulumi.Input<string>;
}
