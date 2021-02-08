import * as pulumi from "@pulumi/pulumi";
export interface ModuleGitlab {
    namespace: pulumi.Input<string>;
}
export interface StackCloudformation {
    entryTemplateFile: pulumi.Input<string>;
    region: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
    templateBucket: pulumi.Input<string>;
}
export interface StackGitlab {
    namespace: pulumi.Input<string>;
}
export interface StackPulumi {
    loginUrl: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
}
