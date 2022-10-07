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
    'GetIPsResult',
    'AwaitableGetIPsResult',
    'get_ips',
]

@pulumi.output_type
class GetIPsResult:
    """
    A collection of values returned by getIPs.
    """
    def __init__(__self__, id=None, ips=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        """
        the list of spacelift.io outgoing IP addresses
        """
        return pulumi.get(self, "ips")


class AwaitableGetIPsResult(GetIPsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIPsResult(
            id=self.id,
            ips=self.ips)


def get_ips(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIPsResult:
    """
    `get_i_ps` returns the list of Spacelift's outgoing IP addresses, which you can use to whitelist connections coming from the Spacelift's "mothership".

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    ips = spacelift.get_ips()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getIPs:getIPs', __args__, opts=opts, typ=GetIPsResult).value

    return AwaitableGetIPsResult(
        id=__ret__.id,
        ips=__ret__.ips)
