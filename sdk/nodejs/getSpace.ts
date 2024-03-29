// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const space = spacelift.getSpace({
 *     spaceId: spacelift_space.space.id,
 * });
 * export const spaceDescription = space.then(space => space.description);
 * ```
 */
export function getSpace(args: GetSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("spacelift:index/getSpace:getSpace", {
        "spaceId": args.spaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpace.
 */
export interface GetSpaceArgs {
    /**
     * immutable ID (slug) of the space
     */
    spaceId: string;
}

/**
 * A collection of values returned by getSpace.
 */
export interface GetSpaceResult {
    /**
     * free-form space description for users
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * indication whether access to this space inherits read access to entities from the parent space
     */
    readonly inheritEntities: boolean;
    /**
     * list of labels describing a space
     */
    readonly labels: string[];
    /**
     * name of the space
     */
    readonly name: string;
    /**
     * immutable ID (slug) of parent space
     */
    readonly parentSpaceId: string;
    /**
     * immutable ID (slug) of the space
     */
    readonly spaceId: string;
}
/**
 * `spacelift.Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const space = spacelift.getSpace({
 *     spaceId: spacelift_space.space.id,
 * });
 * export const spaceDescription = space.then(space => space.description);
 * ```
 */
export function getSpaceOutput(args: GetSpaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpaceResult> {
    return pulumi.output(args).apply((a: any) => getSpace(a, opts))
}

/**
 * A collection of arguments for invoking getSpace.
 */
export interface GetSpaceOutputArgs {
    /**
     * immutable ID (slug) of the space
     */
    spaceId: pulumi.Input<string>;
}
