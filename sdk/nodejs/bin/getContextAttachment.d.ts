import * as pulumi from "@pulumi/pulumi";
export declare function getContextAttachment(args: GetContextAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetContextAttachmentResult>;
/**
 * A collection of arguments for invoking getContextAttachment.
 */
export interface GetContextAttachmentArgs {
    readonly contextId: string;
    readonly moduleId?: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getContextAttachment.
 */
export interface GetContextAttachmentResult {
    readonly contextId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly priority: number;
    readonly stackId?: string;
}
