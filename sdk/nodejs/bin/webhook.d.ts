import * as pulumi from "@pulumi/pulumi";
export declare class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookState, opts?: pulumi.CustomResourceOptions): Webhook;
    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Webhook;
    /**
     * enables or disables sending webhooks
     */
    readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * endpoint to send the POST request to
     */
    readonly endpoint: pulumi.Output<string>;
    /**
     * ID of the module which triggers the webhooks
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * secret used to sign each POST request so you're able to verify that the request comes from us
     */
    readonly secret: pulumi.Output<string | undefined>;
    /**
     * ID of the stack which triggers the webhooks
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Webhook resources.
 */
export interface WebhookState {
    /**
     * enables or disables sending webhooks
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * endpoint to send the POST request to
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * ID of the module which triggers the webhooks
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * secret used to sign each POST request so you're able to verify that the request comes from us
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * ID of the stack which triggers the webhooks
     */
    readonly stackId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    /**
     * enables or disables sending webhooks
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * endpoint to send the POST request to
     */
    readonly endpoint: pulumi.Input<string>;
    /**
     * ID of the module which triggers the webhooks
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * secret used to sign each POST request so you're able to verify that the request comes from us
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * ID of the stack which triggers the webhooks
     */
    readonly stackId?: pulumi.Input<string>;
}
