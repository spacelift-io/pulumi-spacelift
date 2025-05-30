// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as google from "@pulumi/google";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const k8s_coreStack = new spacelift.Stack("k8s-coreStack", {
 *     branch: "master",
 *     repository: "core-infra",
 * });
 * const k8s_coreStackGcpServiceAccount = new spacelift.StackGcpServiceAccount("k8s-coreStackGcpServiceAccount", {
 *     stackId: k8s_coreStack.id,
 *     tokenScopes: [
 *         "https://www.googleapis.com/auth/compute",
 *         "https://www.googleapis.com/auth/cloud-platform",
 *         "https://www.googleapis.com/auth/devstorage.full_control",
 *     ],
 * });
 * const k8s_coregoogle_project = new google.index.Google_project("k8s-coregoogle_project", {
 *     name: "Kubernetes code",
 *     projectId: "unicorn-k8s-core",
 *     orgId: _var.gcp_organization_id,
 * });
 * const k8s_coregoogle_project_iam_member = new google.index.Google_project_iam_member("k8s-coregoogle_project_iam_member", {
 *     project: k8s_coregoogle_project.id,
 *     role: "roles/owner",
 *     member: `serviceAccount:${k8s_coreStackGcpServiceAccount.serviceAccountEmail}`,
 * });
 * ```
 */
export class StackGcpServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing StackGcpServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackGcpServiceAccountState, opts?: pulumi.CustomResourceOptions): StackGcpServiceAccount {
        return new StackGcpServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount';

    /**
     * Returns true if the given object is an instance of StackGcpServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StackGcpServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StackGcpServiceAccount.__pulumiType;
    }

    /**
     * ID of the module which uses GCP service account credentials
     */
    public readonly moduleId!: pulumi.Output<string | undefined>;
    /**
     * Email address of the GCP service account dedicated for this stack
     */
    public /*out*/ readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    public readonly stackId!: pulumi.Output<string | undefined>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    public readonly tokenScopes!: pulumi.Output<string[]>;

    /**
     * Create a StackGcpServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackGcpServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackGcpServiceAccountArgs | StackGcpServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackGcpServiceAccountState | undefined;
            resourceInputs["moduleId"] = state ? state.moduleId : undefined;
            resourceInputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
            resourceInputs["tokenScopes"] = state ? state.tokenScopes : undefined;
        } else {
            const args = argsOrState as StackGcpServiceAccountArgs | undefined;
            if ((!args || args.tokenScopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tokenScopes'");
            }
            resourceInputs["moduleId"] = args ? args.moduleId : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["tokenScopes"] = args ? args.tokenScopes : undefined;
            resourceInputs["serviceAccountEmail"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StackGcpServiceAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StackGcpServiceAccount resources.
 */
export interface StackGcpServiceAccountState {
    /**
     * ID of the module which uses GCP service account credentials
     */
    moduleId?: pulumi.Input<string>;
    /**
     * Email address of the GCP service account dedicated for this stack
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    stackId?: pulumi.Input<string>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a StackGcpServiceAccount resource.
 */
export interface StackGcpServiceAccountArgs {
    /**
     * ID of the module which uses GCP service account credentials
     */
    moduleId?: pulumi.Input<string>;
    /**
     * ID of the stack which uses GCP service account credentials
     */
    stackId?: pulumi.Input<string>;
    /**
     * List of scopes that will be requested when generating temporary GCP service account credentials
     */
    tokenScopes: pulumi.Input<pulumi.Input<string>[]>;
}
