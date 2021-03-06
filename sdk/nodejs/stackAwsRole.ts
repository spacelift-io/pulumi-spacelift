// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class StackAwsRole extends pulumi.CustomResource {
    /**
     * Get an existing StackAwsRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackAwsRoleState, opts?: pulumi.CustomResourceOptions): StackAwsRole {
        return new StackAwsRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/stackAwsRole:StackAwsRole';

    /**
     * Returns true if the given object is an instance of StackAwsRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StackAwsRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StackAwsRole.__pulumiType;
    }

    /**
     * ID of the module which assumes the AWS IAM role
     */
    public readonly moduleId!: pulumi.Output<string | undefined>;
    /**
     * ARN of the AWS IAM role to attach
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * ID of the stack which assumes the AWS IAM role
     */
    public readonly stackId!: pulumi.Output<string | undefined>;

    /**
     * Create a StackAwsRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackAwsRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackAwsRoleArgs | StackAwsRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StackAwsRoleState | undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
        } else {
            const args = argsOrState as StackAwsRoleArgs | undefined;
            if ((!args || args.roleArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StackAwsRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StackAwsRole resources.
 */
export interface StackAwsRoleState {
    /**
     * ID of the module which assumes the AWS IAM role
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ARN of the AWS IAM role to attach
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * ID of the stack which assumes the AWS IAM role
     */
    readonly stackId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StackAwsRole resource.
 */
export interface StackAwsRoleArgs {
    /**
     * ID of the module which assumes the AWS IAM role
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ARN of the AWS IAM role to attach
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * ID of the stack which assumes the AWS IAM role
     */
    readonly stackId?: pulumi.Input<string>;
}
