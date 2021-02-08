import * as pulumi from "@pulumi/pulumi";
export declare class StackGcpServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing StackGcpServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackGcpServiceAccountState, opts?: pulumi.CustomResourceOptions): StackGcpServiceAccount;
    /**
     * Returns true if the given object is an instance of StackGcpServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is StackGcpServiceAccount;
    /**
     * ID of the module which uses GCP service account credentials
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * Email address of the GCP service account dedicated for this stack
     */
    readonly serviceAccountEmail: pulumi.Output<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    readonly tokenScopes: pulumi.Output<string[]>;
    /**
     * Create a StackGcpServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackGcpServiceAccountArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering StackGcpServiceAccount resources.
 */
export interface StackGcpServiceAccountState {
    /**
     * ID of the module which uses GCP service account credentials
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Email address of the GCP service account dedicated for this stack
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    readonly tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}
/**
 * The set of arguments for constructing a StackGcpServiceAccount resource.
 */
export interface StackGcpServiceAccountArgs {
    /**
     * ID of the module which uses GCP service account credentials
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    readonly tokenScopes: pulumi.Input<pulumi.Input<string>[]>;
}
