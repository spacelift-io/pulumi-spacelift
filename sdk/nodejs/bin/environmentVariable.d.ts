import * as pulumi from "@pulumi/pulumi";
export declare class EnvironmentVariable extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentVariableState, opts?: pulumi.CustomResourceOptions): EnvironmentVariable;
    /**
     * Returns true if the given object is an instance of EnvironmentVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is EnvironmentVariable;
    /**
     * SHA-256 checksum of the value
     */
    readonly checksum: pulumi.Output<string>;
    /**
     * ID of the context on which the environment variable is defined
     */
    readonly contextId: pulumi.Output<string | undefined>;
    /**
     * ID of the module on which the environment variable is defined
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * Name of the environment variable
     */
    readonly name: pulumi.Output<string>;
    /**
     * ID of the stack on which the environment variable is defined
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * Value of the environment variable
     */
    readonly value: pulumi.Output<string>;
    /**
     * Indicates whether the value can be read back outside a Run
     */
    readonly writeOnly: pulumi.Output<boolean | undefined>;
    /**
     * Create a EnvironmentVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentVariableArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering EnvironmentVariable resources.
 */
export interface EnvironmentVariableState {
    /**
     * SHA-256 checksum of the value
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * ID of the context on which the environment variable is defined
     */
    readonly contextId?: pulumi.Input<string>;
    /**
     * ID of the module on which the environment variable is defined
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Name of the environment variable
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the stack on which the environment variable is defined
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * Value of the environment variable
     */
    readonly value?: pulumi.Input<string>;
    /**
     * Indicates whether the value can be read back outside a Run
     */
    readonly writeOnly?: pulumi.Input<boolean>;
}
/**
 * The set of arguments for constructing a EnvironmentVariable resource.
 */
export interface EnvironmentVariableArgs {
    /**
     * ID of the context on which the environment variable is defined
     */
    readonly contextId?: pulumi.Input<string>;
    /**
     * ID of the module on which the environment variable is defined
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Name of the environment variable
     */
    readonly name: pulumi.Input<string>;
    /**
     * ID of the stack on which the environment variable is defined
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * Value of the environment variable
     */
    readonly value: pulumi.Input<string>;
    /**
     * Indicates whether the value can be read back outside a Run
     */
    readonly writeOnly?: pulumi.Input<boolean>;
}
