// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookState, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/webhook:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * enables or disables sending webhooks
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * endpoint to send the POST request to
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * ID of the module which triggers the webhooks
     */
    public readonly moduleId!: pulumi.Output<string | undefined>;
    /**
     * secret used to sign each POST request so you're able to verify that the request comes from us
     */
    public readonly secret!: pulumi.Output<string | undefined>;
    /**
     * ID of the stack which triggers the webhooks
     */
    public readonly stackId!: pulumi.Output<string | undefined>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookArgs | WebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebhookState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["moduleId"] = state ? state.moduleId : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
        } else {
            const args = argsOrState as WebhookArgs | undefined;
            if ((!args || args.endpoint === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'endpoint'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["moduleId"] = args ? args.moduleId : undefined;
            inputs["secret"] = args ? args.secret : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Webhook.__pulumiType, name, inputs, opts);
    }
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
