# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetIpsResult',
    'AwaitableGetIpsResult',
    'get_ips',
]

@pulumi.output_type
class GetIpsResult:
    """
    A collection of values returned by getIps.
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
        return pulumi.get(self, "ips")


class AwaitableGetIpsResult(GetIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpsResult(
            id=self.id,
            ips=self.ips)


def get_ips(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpsResult:
    """
    `getIps` returns the list of Spacelift's outgoing IP addresses, which you can use to whitelist connections coming from the Spacelift's "mothership".

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    ips = spacelift.get_ips()
    ```

    <!-- schema generated by tfplugindocs -->
    ## Schema

    ### Optional

    - **id** (String) The ID of this resource.

    ### Read-Only

    - **ips** (Set of String) the list of spacelift.io outgoing IP addresses
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getIps:getIps', __args__, opts=opts, typ=GetIpsResult).value

    return AwaitableGetIpsResult(
        id=__ret__.id,
        ips=__ret__.ips)
