// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getContext(args: GetContextArgs, opts?: pulumi.InvokeOptions): Promise<GetContextResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getContext:getContext", {
        "contextId": args.contextId,
    }, opts);
}

/**
 * A collection of arguments for invoking getContext.
 */
export interface GetContextArgs {
    readonly contextId: string;
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
    readonly name: string;
}
