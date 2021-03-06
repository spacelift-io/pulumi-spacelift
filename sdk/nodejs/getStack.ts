// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getStack:getStack", {
        "stackId": args.stackId,
    }, opts);
}

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
