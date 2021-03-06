import * as pulumi from "@pulumi/pulumi";
export declare class ContextAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ContextAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContextAttachmentState, opts?: pulumi.CustomResourceOptions): ContextAttachment;
    /**
     * Returns true if the given object is an instance of ContextAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ContextAttachment;
    /**
     * ID of the context to attach
     */
    readonly contextId: pulumi.Output<string>;
    /**
     * ID of the module to attach the context to
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * Priority of the context attachment, used in case of conflicts
     */
    readonly priority: pulumi.Output<number | undefined>;
    /**
     * ID of the stack to attach the context to
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * Create a ContextAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContextAttachmentArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering ContextAttachment resources.
 */
export interface ContextAttachmentState {
    /**
     * ID of the context to attach
     */
    readonly contextId?: pulumi.Input<string>;
    /**
     * ID of the module to attach the context to
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Priority of the context attachment, used in case of conflicts
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * ID of the stack to attach the context to
     */
    readonly stackId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a ContextAttachment resource.
 */
export interface ContextAttachmentArgs {
    /**
     * ID of the context to attach
     */
    readonly contextId: pulumi.Input<string>;
    /**
     * ID of the module to attach the context to
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Priority of the context attachment, used in case of conflicts
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * ID of the stack to attach the context to
     */
    readonly stackId?: pulumi.Input<string>;
}
