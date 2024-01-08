# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetGitlabWebhookEndpointResult',
    'AwaitableGetGitlabWebhookEndpointResult',
    'get_gitlab_webhook_endpoint',
    'get_gitlab_webhook_endpoint_output',
]

@pulumi.output_type
class GetGitlabWebhookEndpointResult:
    """
    A collection of values returned by getGitlabWebhookEndpoint.
    """
    def __init__(__self__, id=None, webhook_endpoint=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if webhook_endpoint and not isinstance(webhook_endpoint, str):
            raise TypeError("Expected argument 'webhook_endpoint' to be a str")
        pulumi.set(__self__, "webhook_endpoint", webhook_endpoint)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="webhookEndpoint")
    def webhook_endpoint(self) -> str:
        """
        Gitlab webhook endpoint address
        """
        return pulumi.get(self, "webhook_endpoint")


class AwaitableGetGitlabWebhookEndpointResult(GetGitlabWebhookEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGitlabWebhookEndpointResult(
            id=self.id,
            webhook_endpoint=self.webhook_endpoint)


def get_gitlab_webhook_endpoint(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGitlabWebhookEndpointResult:
    """
    `get_gitlab_webhook_endpoint` returns details about Gitlab webhook endpoint

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    this = spacelift.get_gitlab_webhook_endpoint()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getGitlabWebhookEndpoint:getGitlabWebhookEndpoint', __args__, opts=opts, typ=GetGitlabWebhookEndpointResult).value

    return AwaitableGetGitlabWebhookEndpointResult(
        id=pulumi.get(__ret__, 'id'),
        webhook_endpoint=pulumi.get(__ret__, 'webhook_endpoint'))


@_utilities.lift_output_func(get_gitlab_webhook_endpoint)
def get_gitlab_webhook_endpoint_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGitlabWebhookEndpointResult]:
    """
    `get_gitlab_webhook_endpoint` returns details about Gitlab webhook endpoint

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    this = spacelift.get_gitlab_webhook_endpoint()
    ```
    """
    ...