// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `spacelift.Webhook` represents a webhook endpoint to which Spacelift sends the POST request about run state changes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@pulumi/spacelift";
 *
 * const webhook = spacelift.getWebhook({
 *     webhookId: spacelift_webhook.webhook.id,
 * });
 * ```
 *
 * <!-- schema generated by tfplugindocs -->
 * ## Schema
 *
 * ### Required
 *
 * - **webhook_id** (String) ID of the webhook
 *
 * ### Optional
 *
 * - **id** (String) The ID of this resource.
 * - **module_id** (String) ID of the stack which triggers the webhooks
 * - **stack_id** (String) ID of the stack which triggers the webhooks
 *
 * ### Read-Only
 *
 * - **enabled** (Boolean) enables or disables sending webhooks
 * - **endpoint** (String) endpoint to send the POST request to
 * - **secret** (String) secret used to sign each POST request so you're able to verify that the request comes from us
 */
export function getWebhook(args: GetWebhookArgs, opts?: pulumi.InvokeOptions): Promise<GetWebhookResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("spacelift:index/getWebhook:getWebhook", {
        "moduleId": args.moduleId,
        "stackId": args.stackId,
        "webhookId": args.webhookId,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebhook.
 */
export interface GetWebhookArgs {
    readonly moduleId?: string;
    readonly stackId?: string;
    readonly webhookId: string;
}

/**
 * A collection of values returned by getWebhook.
 */
export interface GetWebhookResult {
    readonly enabled: boolean;
    readonly endpoint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly secret: string;
    readonly stackId?: string;
    readonly webhookId: string;
}
