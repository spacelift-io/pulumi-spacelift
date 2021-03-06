// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class PolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyAttachmentState, opts?: pulumi.CustomResourceOptions): PolicyAttachment {
        return new PolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/policyAttachment:PolicyAttachment';

    /**
     * Returns true if the given object is an instance of PolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAttachment.__pulumiType;
    }

    /**
     * JSON-encoded custom input to be passed to the evaluated document at the "attachment" key
     */
    public readonly customInput!: pulumi.Output<string | undefined>;
    /**
     * ID of the module to attach the policy to
     */
    public readonly moduleId!: pulumi.Output<string | undefined>;
    /**
     * ID of the policy to attach
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * ID of the stack to attach the policy to
     */
    public readonly stackId!: pulumi.Output<string | undefined>;

    /**
     * Create a PolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyAttachmentArgs | PolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyAttachmentState | undefined;
            inputs["customInput"] = state ? state.customInput : undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
        } else {
            const args = argsOrState as PolicyAttachmentArgs | undefined;
            if ((!args || args.policyId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyId'");
            }
            inputs["customInput"] = args ? args.customInput : undefined;
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PolicyAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyAttachment resources.
 */
export interface PolicyAttachmentState {
    /**
     * JSON-encoded custom input to be passed to the evaluated document at the "attachment" key
     */
    readonly customInput?: pulumi.Input<string>;
    /**
     * ID of the module to attach the policy to
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ID of the policy to attach
     */
    readonly policyId?: pulumi.Input<string>;
    /**
     * ID of the stack to attach the policy to
     */
    readonly stackId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyAttachment resource.
 */
export interface PolicyAttachmentArgs {
    /**
     * JSON-encoded custom input to be passed to the evaluated document at the "attachment" key
     */
    readonly customInput?: pulumi.Input<string>;
    /**
     * ID of the module to attach the policy to
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * ID of the policy to attach
     */
    readonly policyId: pulumi.Input<string>;
    /**
     * ID of the stack to attach the policy to
     */
    readonly stackId?: pulumi.Input<string>;
}
