import * as pulumi from "@pulumi/pulumi";
/**
 * `spacelift.MountedFile` represents a file mounted in each Run's workspace that is part of a configuration of a context (`spacelift.Context`), stack (`spacelift.Stack`) or a module (`spacelift.Module`). In principle, it's very similar to an environment variable (`spacelift.EnvironmentVariable`) except that the value is written to the filesystem rather than passed to the environment.
 *
 * ## Schema
 *
 * ### Required
 *
 * - **content** (String, Sensitive) Content of the mounted file encoded using Base-64
 * - **relative_path** (String) Relative path to the mounted file, without the /mnt/workspace/ prefix
 *
 * ### Optional
 *
 * - **context_id** (String) ID of the context on which the mounted file is defined
 * - **id** (String) The ID of this resource.
 * - **module_id** (String) ID of the module on which the mounted file is defined
 * - **stack_id** (String) ID of the stack on which the mounted file is defined
 * - **write_only** (Boolean) Indicates whether the content can be read back outside a Run
 *
 * ### Read-Only
 *
 * - **checksum** (String) SHA-256 checksum of the value
 *
 * ## Import
 *
 * Import is supported using the following syntax
 *
 * ```sh
 *  $ pulumi import spacelift:index/mountedFile:MountedFile ireland-kubeconfig context/$CONTEXT_ID/$MOUNTED_FILE_ID
 * ```
 *
 * ```sh
 *  $ pulumi import spacelift:index/mountedFile:MountedFile module-kubeconfig module/$MODULE_ID/$MOUNTED_FILE_ID
 * ```
 *
 * ```sh
 *  $ pulumi import spacelift:index/mountedFile:MountedFile core-kubeconfig stack/$STACK_ID/$MOUNTED_FILE_ID
 * ```
 */
export declare class MountedFile extends pulumi.CustomResource {
    /**
     * Get an existing MountedFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MountedFileState, opts?: pulumi.CustomResourceOptions): MountedFile;
    /**
     * Returns true if the given object is an instance of MountedFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is MountedFile;
    /**
     * SHA-256 checksum of the value
     */
    readonly checksum: pulumi.Output<string>;
    /**
     * Content of the mounted file encoded using Base-64
     */
    readonly content: pulumi.Output<string>;
    /**
     * ID of the context on which the mounted file is defined
     */
    readonly contextId: pulumi.Output<string | undefined>;
    /**
     * ID of the module on which the mounted file is defined
     */
    readonly moduleId: pulumi.Output<string | undefined>;
    /**
     * Relative path to the mounted file, without the /mnt/workspace/ prefix
     */
    readonly relativePath: pulumi.Output<string>;
    /**
     * ID of the stack on which the mounted file is defined
     */
    readonly stackId: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the content can be read back outside a Run
     */
    readonly writeOnly: pulumi.Output<boolean | undefined>;
    /**
     * Create a MountedFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MountedFileArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering MountedFile resources.
 */
export interface MountedFileState {
    /**
     * SHA-256 checksum of the value
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * Content of the mounted file encoded using Base-64
     */
    readonly content?: pulumi.Input<string>;
    /**
     * ID of the context on which the mounted file is defined
     */
    readonly contextId?: pulumi.Input<string>;
    /**
     * ID of the module on which the mounted file is defined
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Relative path to the mounted file, without the /mnt/workspace/ prefix
     */
    readonly relativePath?: pulumi.Input<string>;
    /**
     * ID of the stack on which the mounted file is defined
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * Indicates whether the content can be read back outside a Run
     */
    readonly writeOnly?: pulumi.Input<boolean>;
}
/**
 * The set of arguments for constructing a MountedFile resource.
 */
export interface MountedFileArgs {
    /**
     * Content of the mounted file encoded using Base-64
     */
    readonly content: pulumi.Input<string>;
    /**
     * ID of the context on which the mounted file is defined
     */
    readonly contextId?: pulumi.Input<string>;
    /**
     * ID of the module on which the mounted file is defined
     */
    readonly moduleId?: pulumi.Input<string>;
    /**
     * Relative path to the mounted file, without the /mnt/workspace/ prefix
     */
    readonly relativePath: pulumi.Input<string>;
    /**
     * ID of the stack on which the mounted file is defined
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * Indicates whether the content can be read back outside a Run
     */
    readonly writeOnly?: pulumi.Input<boolean>;
}
