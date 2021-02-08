import * as pulumi from "@pulumi/pulumi";
export declare class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy;
    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Policy;
    /**
     * Body of the policy
     */
    readonly body: pulumi.Output<string>;
    /**
     * Name of the policy - should be unique in one account
     */
    readonly name: pulumi.Output<string>;
    /**
     * Body of the policy
     */
    readonly type: pulumi.Output<string>;
    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Body of the policy
     */
    readonly body?: pulumi.Input<string>;
    /**
     * Name of the policy - should be unique in one account
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Body of the policy
     */
    readonly type?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Body of the policy
     */
    readonly body: pulumi.Input<string>;
    /**
     * Name of the policy - should be unique in one account
     */
    readonly name: pulumi.Input<string>;
    /**
     * Body of the policy
     */
    readonly type: pulumi.Input<string>;
}
