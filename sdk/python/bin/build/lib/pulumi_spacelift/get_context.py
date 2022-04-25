# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetContextResult',
    'AwaitableGetContextResult',
    'get_context',
]

@pulumi.output_type
class GetContextResult:
    """
    A collection of values returned by getContext.
    """
    def __init__(__self__, context_id=None, description=None, id=None, labels=None, name=None):
        if context_id and not isinstance(context_id, str):
            raise TypeError("Expected argument 'context_id' to be a str")
        pulumi.set(__self__, "context_id", context_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="contextId")
    def context_id(self) -> str:
        return pulumi.get(self, "context_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetContextResult(GetContextResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContextResult(
            context_id=self.context_id,
            description=self.description,
            id=self.id,
            labels=self.labels,
            name=self.name)


def get_context(context_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContextResult:
    """
    `Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`Stack`) or modules (`Module`) using a context attachment (`ContextAttachment`)`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    prod_k8s_ie = spacelift.get_context(context_id="prod-k8s-ie")
    ```

    <!-- schema generated by tfplugindocs -->
    ## Schema

    ### Required

    - **context_id** (String) immutable ID (slug) of the context

    ### Optional

    - **id** (String) The ID of this resource.

    ### Read-Only

    - **description** (String) free-form context description for users
    - **labels** (Set of String)
    - **name** (String) name of the context
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getContext:getContext', __args__, opts=opts, typ=GetContextResult).value

    return AwaitableGetContextResult(
        context_id=__ret__.context_id,
        description=__ret__.description,
        id=__ret__.id,
        labels=__ret__.labels,
        name=__ret__.name)
