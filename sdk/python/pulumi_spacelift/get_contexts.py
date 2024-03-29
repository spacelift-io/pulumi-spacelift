# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetContextsResult',
    'AwaitableGetContextsResult',
    'get_contexts',
    'get_contexts_output',
]

@pulumi.output_type
class GetContextsResult:
    """
    A collection of values returned by getContexts.
    """
    def __init__(__self__, contexts=None, id=None, labels=None):
        if contexts and not isinstance(contexts, list):
            raise TypeError("Expected argument 'contexts' to be a list")
        pulumi.set(__self__, "contexts", contexts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter
    def contexts(self) -> Sequence['outputs.GetContextsContextResult']:
        return pulumi.get(self, "contexts")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence['outputs.GetContextsLabelResult']]:
        """
        Require contexts to have one of the labels
        """
        return pulumi.get(self, "labels")


class AwaitableGetContextsResult(GetContextsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContextsResult(
            contexts=self.contexts,
            id=self.id,
            labels=self.labels)


def get_contexts(labels: Optional[Sequence[pulumi.InputType['GetContextsLabelArgs']]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContextsResult:
    """
    `get_contexts` represents all the contexts in the Spacelift account visible to the API user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    contexts = spacelift.get_contexts(labels=[
        spacelift.GetContextsLabelArgs(
            any_ofs=[
                "foo",
                "bar",
            ],
        ),
        spacelift.GetContextsLabelArgs(
            any_ofs=["baz"],
        ),
    ])
    ```
    """
    __args__ = dict()
    __args__['labels'] = labels
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getContexts:getContexts', __args__, opts=opts, typ=GetContextsResult).value

    return AwaitableGetContextsResult(
        contexts=pulumi.get(__ret__, 'contexts'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'))


@_utilities.lift_output_func(get_contexts)
def get_contexts_output(labels: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetContextsLabelArgs']]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContextsResult]:
    """
    `get_contexts` represents all the contexts in the Spacelift account visible to the API user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    contexts = spacelift.get_contexts(labels=[
        spacelift.GetContextsLabelArgs(
            any_ofs=[
                "foo",
                "bar",
            ],
        ),
        spacelift.GetContextsLabelArgs(
            any_ofs=["baz"],
        ),
    ])
    ```
    """
    ...
