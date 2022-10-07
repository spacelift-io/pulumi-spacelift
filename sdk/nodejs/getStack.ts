// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.Stack` combines source code and configuration to create a runtime environment where resources are managed. In this way it's similar to a stack in AWS CloudFormation, or a project on generic CI/CD platforms.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const k8s_core = pulumi.output(spacelift.getStack({
 *     stackId: "k8s-core",
 * }));
 * ```
 */
export function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getStack:getStack", {
        "afterApplies": args.afterApplies,
        "afterDestroys": args.afterDestroys,
        "afterInits": args.afterInits,
        "afterPerforms": args.afterPerforms,
        "afterPlans": args.afterPlans,
        "beforeApplies": args.beforeApplies,
        "beforeDestroys": args.beforeDestroys,
        "beforeInits": args.beforeInits,
        "beforePerforms": args.beforePerforms,
        "beforePlans": args.beforePlans,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackArgs {
    afterApplies?: string[];
    afterDestroys?: string[];
    afterInits?: string[];
    afterPerforms?: string[];
    afterPlans?: string[];
    beforeApplies?: string[];
    beforeDestroys?: string[];
    beforeInits?: string[];
    beforePerforms?: string[];
    beforePlans?: string[];
    stackId: string;
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
    readonly ansibles: outputs.GetStackAnsible[];
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
    readonly spaceId: string;
    readonly stackId: string;
    readonly terraformVersion: string;
    readonly terraformWorkspace: string;
    readonly workerPoolId: string;
}

export function getStackOutput(args: GetStackOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStackResult> {
    return pulumi.output(args).apply(a => getStack(a, opts))
}

/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackOutputArgs {
    afterApplies?: pulumi.Input<pulumi.Input<string>[]>;
    afterDestroys?: pulumi.Input<pulumi.Input<string>[]>;
    afterInits?: pulumi.Input<pulumi.Input<string>[]>;
    afterPerforms?: pulumi.Input<pulumi.Input<string>[]>;
    afterPlans?: pulumi.Input<pulumi.Input<string>[]>;
    beforeApplies?: pulumi.Input<pulumi.Input<string>[]>;
    beforeDestroys?: pulumi.Input<pulumi.Input<string>[]>;
    beforeInits?: pulumi.Input<pulumi.Input<string>[]>;
    beforePerforms?: pulumi.Input<pulumi.Input<string>[]>;
    beforePlans?: pulumi.Input<pulumi.Input<string>[]>;
    stackId: pulumi.Input<string>;
}
