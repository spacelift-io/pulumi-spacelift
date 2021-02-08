// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

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
