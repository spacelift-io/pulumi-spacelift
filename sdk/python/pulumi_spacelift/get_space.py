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
    'GetSpaceResult',
    'AwaitableGetSpaceResult',
    'get_space',
    'get_space_output',
]

@pulumi.output_type
class GetSpaceResult:
    """
    A collection of values returned by getSpace.
    """
    def __init__(__self__, description=None, id=None, inherit_entities=None, labels=None, name=None, parent_space_id=None, space_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inherit_entities and not isinstance(inherit_entities, bool):
            raise TypeError("Expected argument 'inherit_entities' to be a bool")
        pulumi.set(__self__, "inherit_entities", inherit_entities)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_space_id and not isinstance(parent_space_id, str):
            raise TypeError("Expected argument 'parent_space_id' to be a str")
        pulumi.set(__self__, "parent_space_id", parent_space_id)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        free-form space description for users
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inheritEntities")
    def inherit_entities(self) -> bool:
        """
        indication whether access to this space inherits read access to entities from the parent space
        """
        return pulumi.get(self, "inherit_entities")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        """
        list of labels describing a space
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        name of the space
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentSpaceId")
    def parent_space_id(self) -> str:
        """
        immutable ID (slug) of parent space
        """
        return pulumi.get(self, "parent_space_id")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        immutable ID (slug) of the space
        """
        return pulumi.get(self, "space_id")


class AwaitableGetSpaceResult(GetSpaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSpaceResult(
            description=self.description,
            id=self.id,
            inherit_entities=self.inherit_entities,
            labels=self.labels,
            name=self.name,
            parent_space_id=self.parent_space_id,
            space_id=self.space_id)


def get_space(space_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSpaceResult:
    """
    `Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    space = spacelift.get_space(space_id=spacelift_space["space"]["id"])
    pulumi.export("spaceDescription", space.description)
    ```


    :param str space_id: immutable ID (slug) of the space
    """
    __args__ = dict()
    __args__['spaceId'] = space_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getSpace:getSpace', __args__, opts=opts, typ=GetSpaceResult).value

    return AwaitableGetSpaceResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        inherit_entities=pulumi.get(__ret__, 'inherit_entities'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        parent_space_id=pulumi.get(__ret__, 'parent_space_id'),
        space_id=pulumi.get(__ret__, 'space_id'))


@_utilities.lift_output_func(get_space)
def get_space_output(space_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSpaceResult]:
    """
    `Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    space = spacelift.get_space(space_id=spacelift_space["space"]["id"])
    pulumi.export("spaceDescription", space.description)
    ```


    :param str space_id: immutable ID (slug) of the space
    """
    ...
