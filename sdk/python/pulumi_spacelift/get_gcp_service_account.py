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
    'GetGcpServiceAccountResult',
    'AwaitableGetGcpServiceAccountResult',
    'get_gcp_service_account',
    'get_gcp_service_account_output',
]

@pulumi.output_type
class GetGcpServiceAccountResult:
    """
    A collection of values returned by getGcpServiceAccount.
    """
    def __init__(__self__, id=None, module_id=None, service_account_email=None, stack_id=None, token_scopes=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if service_account_email and not isinstance(service_account_email, str):
            raise TypeError("Expected argument 'service_account_email' to be a str")
        pulumi.set(__self__, "service_account_email", service_account_email)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if token_scopes and not isinstance(token_scopes, list):
            raise TypeError("Expected argument 'token_scopes' to be a list")
        pulumi.set(__self__, "token_scopes", token_scopes)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[str]:
        """
        ID of the stack which uses GCP service account credentials
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> str:
        """
        email address of the GCP service account dedicated for this stack
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        """
        ID of the stack which uses GCP service account credentials
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Sequence[str]:
        """
        list of Google API scopes
        """
        return pulumi.get(self, "token_scopes")


class AwaitableGetGcpServiceAccountResult(GetGcpServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGcpServiceAccountResult(
            id=self.id,
            module_id=self.module_id,
            service_account_email=self.service_account_email,
            stack_id=self.stack_id,
            token_scopes=self.token_scopes)


def get_gcp_service_account(module_id: Optional[str] = None,
                            stack_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGcpServiceAccountResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    k8s_module = spacelift.get_gcp_service_account(module_id="k8s-module")
    k8s_core = spacelift.get_gcp_service_account(stack_id="k8s-core")
    ```


    :param str module_id: ID of the stack which uses GCP service account credentials
    :param str stack_id: ID of the stack which uses GCP service account credentials
    """
    __args__ = dict()
    __args__['moduleId'] = module_id
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getGcpServiceAccount:getGcpServiceAccount', __args__, opts=opts, typ=GetGcpServiceAccountResult).value

    return AwaitableGetGcpServiceAccountResult(
        id=__ret__.id,
        module_id=__ret__.module_id,
        service_account_email=__ret__.service_account_email,
        stack_id=__ret__.stack_id,
        token_scopes=__ret__.token_scopes)


@_utilities.lift_output_func(get_gcp_service_account)
def get_gcp_service_account_output(module_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   stack_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGcpServiceAccountResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    k8s_module = spacelift.get_gcp_service_account(module_id="k8s-module")
    k8s_core = spacelift.get_gcp_service_account(stack_id="k8s-core")
    ```


    :param str module_id: ID of the stack which uses GCP service account credentials
    :param str stack_id: ID of the stack which uses GCP service account credentials
    """
    ...
