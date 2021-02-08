import * as pulumi from "@pulumi/pulumi";
export declare function getMountedFile(args: GetMountedFileArgs, opts?: pulumi.InvokeOptions): Promise<GetMountedFileResult>;
/**
 * A collection of arguments for invoking getMountedFile.
 */
export interface GetMountedFileArgs {
    readonly contextId?: string;
    readonly moduleId?: string;
    readonly relativePath: string;
    readonly stackId?: string;
}
/**
 * A collection of values returned by getMountedFile.
 */
export interface GetMountedFileResult {
    readonly checksum: string;
    readonly content: string;
    readonly contextId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly moduleId?: string;
    readonly relativePath: string;
    readonly stackId?: string;
    readonly writeOnly: boolean;
}
