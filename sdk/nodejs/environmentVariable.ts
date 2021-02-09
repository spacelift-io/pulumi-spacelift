// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EnvironmentVariable extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentVariableState, opts?: pulumi.CustomResourceOptions): EnvironmentVariable {
        return new EnvironmentVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/environmentVariable:EnvironmentVariable';

    /**
     * Returns true if the given object is an instance of EnvironmentVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentVariable.__pulumiType;
    }

    /**
     * SHA-256 checksum of the value
     */
    public /*out*/ readonly checksum!: pulumi.Output<string>;
    /**
     * ID of the context on which the environment variable is defined
     */
    public readonly contextId!: pulumi.Output<string | undefined>;
    /**
     * ID of the module on which the environment variable is defined
     */
    public readonly moduleId!: pulumi.Output<string | undefined>;
    /**
     * Name of the environment variable
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the stack on which the environment variable is defined
     */
    public readonly stackId!: pulumi.Output<string | undefined>;
    /**
     * Value of the environment variable
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * Indicates whether the value can be read back outside a Run
     */
    public readonly writeOnly!: pulumi.Output<boolean | undefined>;

    /**
     * Create a EnvironmentVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentVariableArgs | EnvironmentVariableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EnvironmentVariableState | undefined;
            inputs["checksum"] = state ? state.checksum : undefined;
            inputs["contextId"] = state ? state.contextId : undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["value"] = state ? state.value : undefined;
            inputs["writeOnly"] = state ? state.writeOnly : undefined;
        } else {
            const args = argsOrState as EnvironmentVariableArgs | undefined;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.value === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'value'");
            }
            inputs["contextId"] = args ? args.contextId : undefined;
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["writeOnly"] = args ? args.writeOnly : undefined;
            inputs["checksum"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EnvironmentVariable.__pulumiType, name, inputs, opts);
    }
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