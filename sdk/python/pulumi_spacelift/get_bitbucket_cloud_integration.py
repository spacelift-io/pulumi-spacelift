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
    'GetBitbucketCloudIntegrationResult',
    'AwaitableGetBitbucketCloudIntegrationResult',
    'get_bitbucket_cloud_integration',
    'get_bitbucket_cloud_integration_output',
]

@pulumi.output_type
class GetBitbucketCloudIntegrationResult:
    """
    A collection of values returned by getBitbucketCloudIntegration.
    """
    def __init__(__self__, description=None, id=None, is_default=None, labels=None, name=None, space_id=None, username=None, webhook_url=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if webhook_url and not isinstance(webhook_url, str):
            raise TypeError("Expected argument 'webhook_url' to be a str")
        pulumi.set(__self__, "webhook_url", webhook_url)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Bitbucket Cloud integration description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Bitbucket Cloud integration id. If not provided, the default integration will be returned
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        """
        Bitbucket Cloud integration is default
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        """
        Bitbucket Cloud integration labels
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Bitbucket Cloud integration name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        Bitbucket Cloud integration space id
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Bitbucket Cloud username
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="webhookUrl")
    def webhook_url(self) -> str:
        """
        Bitbucket Cloud integration webhook URL
        """
        return pulumi.get(self, "webhook_url")


class AwaitableGetBitbucketCloudIntegrationResult(GetBitbucketCloudIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBitbucketCloudIntegrationResult(
            description=self.description,
            id=self.id,
            is_default=self.is_default,
            labels=self.labels,
            name=self.name,
            space_id=self.space_id,
            username=self.username,
            webhook_url=self.webhook_url)


def get_bitbucket_cloud_integration(id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBitbucketCloudIntegrationResult:
    """
    `get_bitbucket_cloud_integration` returns details about Bitbucket Cloud integration

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    bitbucket_cloud_integration = spacelift.get_bitbucket_cloud_integration()
    ```


    :param str id: Bitbucket Cloud integration id. If not provided, the default integration will be returned
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getBitbucketCloudIntegration:getBitbucketCloudIntegration', __args__, opts=opts, typ=GetBitbucketCloudIntegrationResult).value

    return AwaitableGetBitbucketCloudIntegrationResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        space_id=pulumi.get(__ret__, 'space_id'),
        username=pulumi.get(__ret__, 'username'),
        webhook_url=pulumi.get(__ret__, 'webhook_url'))


@_utilities.lift_output_func(get_bitbucket_cloud_integration)
def get_bitbucket_cloud_integration_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBitbucketCloudIntegrationResult]:
    """
    `get_bitbucket_cloud_integration` returns details about Bitbucket Cloud integration

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    bitbucket_cloud_integration = spacelift.get_bitbucket_cloud_integration()
    ```


    :param str id: Bitbucket Cloud integration id. If not provided, the default integration will be returned
    """
    ...
