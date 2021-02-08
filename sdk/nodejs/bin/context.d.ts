import * as pulumi from "@pulumi/pulumi";
export declare class Context extends pulumi.CustomResource {
    /**
     * Get an existing Context resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContextState, opts?: pulumi.CustomResourceOptions): Context;
    /**
     * Returns true if the given object is an instance of Context.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Context;
    /**
     * Free-form context description for users
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * Name of the context - should be unique in one account
     */
    readonly name: pulumi.Output<string>;
    /**
     * Create a Context resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContextArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Context resources.
 */
export interface ContextState {
    /**
     * Free-form context description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the context - should be unique in one account
     */
    readonly name?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Context resource.
 */
export interface ContextArgs {
    /**
     * Free-form context description for users
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the context - should be unique in one account
     */
    readonly name: pulumi.Input<string>;
}
