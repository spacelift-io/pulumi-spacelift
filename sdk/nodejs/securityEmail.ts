// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.SecurityEmail` represents an email address that receives notifications about security issues in Spacelift.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const example = new spacelift.SecurityEmail("example", {email: "user@example.com"});
 * ```
 */
export class SecurityEmail extends pulumi.CustomResource {
    /**
     * Get an existing SecurityEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityEmailState, opts?: pulumi.CustomResourceOptions): SecurityEmail {
        return new SecurityEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/securityEmail:SecurityEmail';

    /**
     * Returns true if the given object is an instance of SecurityEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityEmail.__pulumiType;
    }

    /**
     * Email address to which the security notifications are sent
     */
    public readonly email!: pulumi.Output<string>;

    /**
     * Create a SecurityEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityEmailArgs | SecurityEmailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityEmailState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
        } else {
            const args = argsOrState as SecurityEmailArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityEmail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityEmail resources.
 */
export interface SecurityEmailState {
    /**
     * Email address to which the security notifications are sent
     */
    email?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityEmail resource.
 */
export interface SecurityEmailArgs {
    /**
     * Email address to which the security notifications are sent
     */
    email: pulumi.Input<string>;
}
