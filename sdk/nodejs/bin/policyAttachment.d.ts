import * as pulumi from "@pulumi/pulumi";
export declare class PolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyAttachmentState, opts?: pulumi.CustomResourceOptions): PolicyAttachment;
    /**
     * Returns true if the given object is an instance of PolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is PolicyAttachment;
    /**
     * JSON-encoded custom input to be passed to the evaluated document at the "attachment" key
     */
    readonly customInput: pulumi.Output<string | undefined>;
    /**
     * ID of the module to attach the policy to
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * ID of the policy to attach
     */
    readonly policyId: pulumi.Output<string>;
    /**
     * ID of the stack to attach the policy to
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * Create a PolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions);
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
