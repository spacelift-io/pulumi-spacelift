// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getCurrentStack(opts?: pulumi.InvokeOptions): Promise<GetCurrentStackResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getCurrentStack:getCurrentStack", {
    }, opts);
}

/**
 * A collection of values returned by getCurrentStack.
 */
export interface GetCurrentStackResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
