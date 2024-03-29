// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `spacelift.AzureIntegration` represents an integration with an Azure AD tenant. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect. Note that you will need to provide admin consent manually for the integration to work
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as spacelift from "@spacelift-io/pulumi-spacelift";
 *
 * const example = new spacelift.AzureIntegration("example", {
 *     defaultSubscriptionId: "default-subscription-id",
 *     labels: [
 *         "one",
 *         "two",
 *     ],
 *     tenantId: "tenant-id",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import spacelift:index/azureIntegration:AzureIntegration example $INTEGRATION_ID
 * ```
 */
export class AzureIntegration extends pulumi.CustomResource {
    /**
     * Get an existing AzureIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureIntegrationState, opts?: pulumi.CustomResourceOptions): AzureIntegration {
        return new AzureIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'spacelift:index/azureIntegration:AzureIntegration';

    /**
     * Returns true if the given object is an instance of AzureIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureIntegration.__pulumiType;
    }

    /**
     * Indicates whether admin consent has been performed for the AAD Application.
     */
    public /*out*/ readonly adminConsentProvided!: pulumi.Output<boolean>;
    /**
     * The URL to use to provide admin consent to the application in the customer's tenant
     */
    public /*out*/ readonly adminConsentUrl!: pulumi.Output<string>;
    /**
     * The applicationId of the Azure AD application used by the integration.
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * The default subscription ID to use, if one isn't specified at the stack/module level
     */
    public readonly defaultSubscriptionId!: pulumi.Output<string | undefined>;
    /**
     * The display name for the application in Azure. This is automatically generated when the integration is created, and cannot be changed without deleting and recreating the integration.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * Labels to set on the integration
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The friendly name of the integration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID (slug) of the space the integration is in
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The Azure AD tenant ID
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AzureIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureIntegrationArgs | AzureIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureIntegrationState | undefined;
            resourceInputs["adminConsentProvided"] = state ? state.adminConsentProvided : undefined;
            resourceInputs["adminConsentUrl"] = state ? state.adminConsentUrl : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["defaultSubscriptionId"] = state ? state.defaultSubscriptionId : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AzureIntegrationArgs | undefined;
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["defaultSubscriptionId"] = args ? args.defaultSubscriptionId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["adminConsentProvided"] = undefined /*out*/;
            resourceInputs["adminConsentUrl"] = undefined /*out*/;
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureIntegration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureIntegration resources.
 */
export interface AzureIntegrationState {
    /**
     * Indicates whether admin consent has been performed for the AAD Application.
     */
    adminConsentProvided?: pulumi.Input<boolean>;
    /**
     * The URL to use to provide admin consent to the application in the customer's tenant
     */
    adminConsentUrl?: pulumi.Input<string>;
    /**
     * The applicationId of the Azure AD application used by the integration.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The default subscription ID to use, if one isn't specified at the stack/module level
     */
    defaultSubscriptionId?: pulumi.Input<string>;
    /**
     * The display name for the application in Azure. This is automatically generated when the integration is created, and cannot be changed without deleting and recreating the integration.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Labels to set on the integration
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The friendly name of the integration
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the integration is in
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The Azure AD tenant ID
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureIntegration resource.
 */
export interface AzureIntegrationArgs {
    /**
     * The default subscription ID to use, if one isn't specified at the stack/module level
     */
    defaultSubscriptionId?: pulumi.Input<string>;
    /**
     * Labels to set on the integration
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The friendly name of the integration
     */
    name?: pulumi.Input<string>;
    /**
     * ID (slug) of the space the integration is in
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The Azure AD tenant ID
     */
    tenantId: pulumi.Input<string>;
}
