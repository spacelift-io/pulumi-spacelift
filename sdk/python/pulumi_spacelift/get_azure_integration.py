# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAzureIntegrationResult',
    'AwaitableGetAzureIntegrationResult',
    'get_azure_integration',
    'get_azure_integration_output',
]

@pulumi.output_type
class GetAzureIntegrationResult:
    """
    A collection of values returned by getAzureIntegration.
    """
    def __init__(__self__, admin_consent_provided=None, admin_consent_url=None, application_id=None, default_subscription_id=None, display_name=None, id=None, integration_id=None, labels=None, name=None, space_id=None, tenant_id=None):
        if admin_consent_provided and not isinstance(admin_consent_provided, bool):
            raise TypeError("Expected argument 'admin_consent_provided' to be a bool")
        pulumi.set(__self__, "admin_consent_provided", admin_consent_provided)
        if admin_consent_url and not isinstance(admin_consent_url, str):
            raise TypeError("Expected argument 'admin_consent_url' to be a str")
        pulumi.set(__self__, "admin_consent_url", admin_consent_url)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if default_subscription_id and not isinstance(default_subscription_id, str):
            raise TypeError("Expected argument 'default_subscription_id' to be a str")
        pulumi.set(__self__, "default_subscription_id", default_subscription_id)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if integration_id and not isinstance(integration_id, str):
            raise TypeError("Expected argument 'integration_id' to be a str")
        pulumi.set(__self__, "integration_id", integration_id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="adminConsentProvided")
    def admin_consent_provided(self) -> bool:
        """
        Indicates whether admin consent has been performed for the AAD Application.
        """
        return pulumi.get(self, "admin_consent_provided")

    @property
    @pulumi.getter(name="adminConsentUrl")
    def admin_consent_url(self) -> str:
        """
        The URL to use to provide admin consent to the application in the customer's tenant
        """
        return pulumi.get(self, "admin_consent_url")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        """
        The applicationId of the Azure AD application used by the integration.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="defaultSubscriptionId")
    def default_subscription_id(self) -> str:
        """
        The default subscription ID to use, if one isn't specified at the stack/module level
        """
        return pulumi.get(self, "default_subscription_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for the application in Azure. This is automatically generated when the integration is created, and cannot be changed without deleting and recreating the integration.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> str:
        """
        Immutable ID of the integration. Either `integration_id` or `name` must be specified.
        """
        return pulumi.get(self, "integration_id")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        """
        Labels to set on the integration
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The friendly name of the integration. Either `integration_id` or `name` must be specified.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        ID (slug) of the space the integration is in
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The Azure AD tenant ID
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetAzureIntegrationResult(GetAzureIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAzureIntegrationResult(
            admin_consent_provided=self.admin_consent_provided,
            admin_consent_url=self.admin_consent_url,
            application_id=self.application_id,
            default_subscription_id=self.default_subscription_id,
            display_name=self.display_name,
            id=self.id,
            integration_id=self.integration_id,
            labels=self.labels,
            name=self.name,
            space_id=self.space_id,
            tenant_id=self.tenant_id)


def get_azure_integration(integration_id: Optional[str] = None,
                          name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAzureIntegrationResult:
    """
    `AzureIntegration` represents an integration with an Azure AD tenant. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect. Note that you will need to provide admin consent manually for the integration to work

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    example = spacelift.get_azure_integration(name="Production")
    ```


    :param str integration_id: Immutable ID of the integration. Either `integration_id` or `name` must be specified.
    :param str name: The friendly name of the integration. Either `integration_id` or `name` must be specified.
    """
    __args__ = dict()
    __args__['integrationId'] = integration_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getAzureIntegration:getAzureIntegration', __args__, opts=opts, typ=GetAzureIntegrationResult).value

    return AwaitableGetAzureIntegrationResult(
        admin_consent_provided=pulumi.get(__ret__, 'admin_consent_provided'),
        admin_consent_url=pulumi.get(__ret__, 'admin_consent_url'),
        application_id=pulumi.get(__ret__, 'application_id'),
        default_subscription_id=pulumi.get(__ret__, 'default_subscription_id'),
        display_name=pulumi.get(__ret__, 'display_name'),
        id=pulumi.get(__ret__, 'id'),
        integration_id=pulumi.get(__ret__, 'integration_id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        space_id=pulumi.get(__ret__, 'space_id'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_azure_integration)
def get_azure_integration_output(integration_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 name: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAzureIntegrationResult]:
    """
    `AzureIntegration` represents an integration with an Azure AD tenant. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect. Note that you will need to provide admin consent manually for the integration to work

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    example = spacelift.get_azure_integration(name="Production")
    ```


    :param str integration_id: Immutable ID of the integration. Either `integration_id` or `name` must be specified.
    :param str name: The friendly name of the integration. Either `integration_id` or `name` must be specified.
    """
    ...
