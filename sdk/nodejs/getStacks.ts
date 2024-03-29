// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `spacelift.getStacks` represents all the stacks in the Spacelift account visible to the API user, matching predicates.
 */
export function getStacks(args?: GetStacksArgs, opts?: pulumi.InvokeOptions): Promise<GetStacksResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getStacks:getStacks", {
        "administrative": args.administrative,
        "branch": args.branch,
        "commit": args.commit,
        "labels": args.labels,
        "locked": args.locked,
        "name": args.name,
        "projectRoot": args.projectRoot,
        "repository": args.repository,
        "state": args.state,
        "vendor": args.vendor,
        "workerPool": args.workerPool,
    }, opts);
}

/**
 * A collection of arguments for invoking getStacks.
 */
export interface GetStacksArgs {
    administrative?: inputs.GetStacksAdministrative;
    branch?: inputs.GetStacksBranch;
    /**
     * Require stacks to be on one of the commits
     */
    commit?: inputs.GetStacksCommit;
    labels?: inputs.GetStacksLabel[];
    /**
     * Require stacks to be locked
     */
    locked?: inputs.GetStacksLocked;
    name?: inputs.GetStacksName;
    projectRoot?: inputs.GetStacksProjectRoot;
    repository?: inputs.GetStacksRepository;
    /**
     * Require stacks to have one of the states
     */
    state?: inputs.GetStacksState;
    /**
     * Require stacks to use one of the IaC vendors
     */
    vendor?: inputs.GetStacksVendor;
    /**
     * Require stacks to use one of the worker pools
     */
    workerPool?: inputs.GetStacksWorkerPool;
}

/**
 * A collection of values returned by getStacks.
 */
export interface GetStacksResult {
    /**
     * Require stacks to be administrative or not
     */
    readonly administrative?: outputs.GetStacksAdministrative;
    /**
     * Require stacks to be on one of the branches
     */
    readonly branch?: outputs.GetStacksBranch;
    /**
     * Require stacks to be on one of the commits
     */
    readonly commit?: outputs.GetStacksCommit;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Require stacks to have one of the labels
     */
    readonly labels?: outputs.GetStacksLabel[];
    /**
     * Require stacks to be locked
     */
    readonly locked?: outputs.GetStacksLocked;
    /**
     * Require stacks to have one of the names
     */
    readonly name?: outputs.GetStacksName;
    /**
     * Require stacks to be in one of the project roots
     */
    readonly projectRoot?: outputs.GetStacksProjectRoot;
    /**
     * Require stacks to be in one of the repositories
     */
    readonly repository?: outputs.GetStacksRepository;
    /**
     * List of stacks matching the predicates
     */
    readonly stacks: outputs.GetStacksStack[];
    /**
     * Require stacks to have one of the states
     */
    readonly state?: outputs.GetStacksState;
    /**
     * Require stacks to use one of the IaC vendors
     */
    readonly vendor?: outputs.GetStacksVendor;
    /**
     * Require stacks to use one of the worker pools
     */
    readonly workerPool?: outputs.GetStacksWorkerPool;
}
/**
 * `spacelift.getStacks` represents all the stacks in the Spacelift account visible to the API user, matching predicates.
 */
export function getStacksOutput(args?: GetStacksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStacksResult> {
    return pulumi.output(args).apply((a: any) => getStacks(a, opts))
}

/**
 * A collection of arguments for invoking getStacks.
 */
export interface GetStacksOutputArgs {
    administrative?: pulumi.Input<inputs.GetStacksAdministrativeArgs>;
    branch?: pulumi.Input<inputs.GetStacksBranchArgs>;
    /**
     * Require stacks to be on one of the commits
     */
    commit?: pulumi.Input<inputs.GetStacksCommitArgs>;
    labels?: pulumi.Input<pulumi.Input<inputs.GetStacksLabelArgs>[]>;
    /**
     * Require stacks to be locked
     */
    locked?: pulumi.Input<inputs.GetStacksLockedArgs>;
    name?: pulumi.Input<inputs.GetStacksNameArgs>;
    projectRoot?: pulumi.Input<inputs.GetStacksProjectRootArgs>;
    repository?: pulumi.Input<inputs.GetStacksRepositoryArgs>;
    /**
     * Require stacks to have one of the states
     */
    state?: pulumi.Input<inputs.GetStacksStateArgs>;
    /**
     * Require stacks to use one of the IaC vendors
     */
    vendor?: pulumi.Input<inputs.GetStacksVendorArgs>;
    /**
     * Require stacks to use one of the worker pools
     */
    workerPool?: pulumi.Input<inputs.GetStacksWorkerPoolArgs>;
}
