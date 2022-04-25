// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.EnvironmentVariable` defines an environment variable on the context (`spacelift.Context`), stack (`spacelift.Stack`) or a module (`spacelift.Module`), thereby allowing to pass and share various secrets and configuration with Spacelift stacks.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * // For a context
 * const ireland_kubeconfig = pulumi.output(spacelift.getEnvironmentVariable({
 *     contextId: "prod-k8s-ie",
 *     name: "KUBECONFIG",
 * }, { async: true }));
 * // For a module
 * const module_kubeconfig = pulumi.output(spacelift.getEnvironmentVariable({
 *     moduleId: "k8s-module",
 *     name: "KUBECONFIG",
 * }, { async: true }));
 * // For a stack
 * const core_kubeconfig = pulumi.output(spacelift.getEnvironmentVariable({
 *     name: "KUBECONFIG",
 *     stackId: "k8s-core",
 * }, { async: true }));
 * ```
 *
 * <!-- schema generated by tfplugindocs -->
 * ## Schema
 *
 * ### Required
 *
 * - **name** (String) name of the environment variable
 *
 * ### Optional
 *
 * - **context_id** (String) ID of the context on which the environment variable is defined
 * - **id** (String) The ID of this resource.
 * - **module_id** (String) ID of the module on which the environment variable is defined
 * - **stack_id** (String) ID of the stack on which the environment variable is defined
 *
 * ### Read-Only
 *
 * - **checksum** (String) SHA-256 checksum of the value
 * - **value** (String, Sensitive) value of the environment variable
 * - **write_only** (Boolean) indicates whether the value can be read back outside a Run
 */
export function getEnvironmentVariable(args: GetEnvironmentVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentVariableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getEnvironmentVariable:getEnvironmentVariable", {
        "contextId": args.contextId,
        "moduleId": args.moduleId,
        "name": args.name,
        "stackId": args.stackId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnvironmentVariable.
 */
export interface GetEnvironmentVariableArgs {
    readonly contextId?: string;
    readonly moduleId?: string;
    readonly name: string;
    readonly stackId?: string;
}

/**
 * A collection of values returned by getEnvironmentVariable.
 */
export interface GetEnvironmentVariableResult {
    readonly checksum: string;
    readonly contextId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly name: string;
    readonly stackId?: string;
    readonly value: string;
    readonly writeOnly: boolean;
}
