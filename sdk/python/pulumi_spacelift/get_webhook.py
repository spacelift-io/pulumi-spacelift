# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetWebhookResult',
    'AwaitableGetWebhookResult',
    'get_webhook',
]

@pulumi.output_type
class GetWebhookResult:
    """
    A collection of values returned by getWebhook.
    """
    def __init__(__self__, enabled=None, endpoint=None, id=None, module_id=None, secret=None, stack_id=None, webhook_id=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if secret and not isinstance(secret, str):
            raise TypeError("Expected argument 'secret' to be a str")
        pulumi.set(__self__, "secret", secret)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if webhook_id and not isinstance(webhook_id, str):
            raise TypeError("Expected argument 'webhook_id' to be a str")
        pulumi.set(__self__, "webhook_id", webhook_id)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        return pulumi.get(self, "endpoint")

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
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def secret(self) -> str:
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="webhookId")
    def webhook_id(self) -> str:
        return pulumi.get(self, "webhook_id")


class AwaitableGetWebhookResult(GetWebhookResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebhookResult(
            enabled=self.enabled,
            endpoint=self.endpoint,
            id=self.id,
            module_id=self.module_id,
            secret=self.secret,
            stack_id=self.stack_id,
            webhook_id=self.webhook_id)


def get_webhook(module_id: Optional[str] = None,
                stack_id: Optional[str] = None,
                webhook_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebhookResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['moduleId'] = module_id
    __args__['stackId'] = stack_id
    __args__['webhookId'] = webhook_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getWebhook:getWebhook', __args__, opts=opts, typ=GetWebhookResult).value

    return AwaitableGetWebhookResult(
        enabled=__ret__.enabled,
        endpoint=__ret__.endpoint,
        id=__ret__.id,
        module_id=__ret__.module_id,
        secret=__ret__.secret,
        stack_id=__ret__.stack_id,
        webhook_id=__ret__.webhook_id)