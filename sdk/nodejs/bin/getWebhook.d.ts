import * as pulumi from "@pulumi/pulumi";
export declare function getWebhook(args: GetWebhookArgs, opts?: pulumi.InvokeOptions): Promise<GetWebhookResult>;
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
