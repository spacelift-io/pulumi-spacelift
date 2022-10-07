// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`spacelift.Stack`) or modules (`spacelift.Module`) using a context attachment (`spacelift.ContextAttachment`)`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const prod_k8s_ie = pulumi.output(spacelift.getContext({
 *     contextId: "prod-k8s-ie",
 * }));
 * ```
 */
export function getContext(args: GetContextArgs, opts?: pulumi.InvokeOptions): Promise<GetContextResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("spacelift:index/getContext:getContext", {
        "contextId": args.contextId,
    }, opts);
}

/**
 * A collection of arguments for invoking getContext.
 */
export interface GetContextArgs {
    contextId: string;
}

/**
 * A collection of values returned by getContext.
 */
export interface GetContextResult {
    readonly contextId: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: string[];
    readonly name: string;
    readonly spaceId: string;
}

export function getContextOutput(args: GetContextOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContextResult> {
    return pulumi.output(args).apply(a => getContext(a, opts))
}

/**
 * A collection of arguments for invoking getContext.
 */
export interface GetContextOutputArgs {
    contextId: pulumi.Input<string>;
}
