import * as pulumi from "@pulumi/pulumi";
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
 * }, { async: true }));
 * ```
 *
 * <!-- schema generated by tfplugindocs -->
 * ## Schema
 *
 * ### Required
 *
 * - **context_id** (String) immutable ID (slug) of the context
 *
 * ### Optional
 *
 * - **id** (String) The ID of this resource.
 *
 * ### Read-Only
 *
 * - **description** (String) free-form context description for users
 * - **labels** (Set of String)
 * - **name** (String) name of the context
 */
export declare function getContext(args: GetContextArgs, opts?: pulumi.InvokeOptions): Promise<GetContextResult>;
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
    readonly labels: string[];
    readonly name: string;
}
