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
    'GetCurrentStackResult',
    'AwaitableGetCurrentStackResult',
    'get_current_stack',
    'get_current_stack_output',
]

@pulumi.output_type
class GetCurrentStackResult:
    """
    A collection of values returned by getCurrentStack.
    """
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetCurrentStackResult(GetCurrentStackResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCurrentStackResult(
            id=self.id)


def get_current_stack(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCurrentStackResult:
    """
    `get_current_stack` is a data source that provides information about the current administrative stack if the run is executed within Spacelift by a stack or module. This allows clever tricks like attaching contexts or policies to the stack that manages them.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    this = spacelift.get_current_stack()
    core_kubeconfig = spacelift.EnvironmentVariable("core-kubeconfig",
        stack_id=this.id,
        value="bacon")
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getCurrentStack:getCurrentStack', __args__, opts=opts, typ=GetCurrentStackResult).value

    return AwaitableGetCurrentStackResult(
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_current_stack)
def get_current_stack_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCurrentStackResult]:
    """
    `get_current_stack` is a data source that provides information about the current administrative stack if the run is executed within Spacelift by a stack or module. This allows clever tricks like attaching contexts or policies to the stack that manages them.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    this = spacelift.get_current_stack()
    core_kubeconfig = spacelift.EnvironmentVariable("core-kubeconfig",
        stack_id=this.id,
        value="bacon")
    ```
    """
    ...
