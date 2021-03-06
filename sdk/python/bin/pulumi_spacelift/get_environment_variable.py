# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetEnvironmentVariableResult',
    'AwaitableGetEnvironmentVariableResult',
    'get_environment_variable',
]

@pulumi.output_type
class GetEnvironmentVariableResult:
    """
    A collection of values returned by getEnvironmentVariable.
    """
    def __init__(__self__, checksum=None, context_id=None, id=None, module_id=None, name=None, stack_id=None, value=None, write_only=None):
        if checksum and not isinstance(checksum, str):
            raise TypeError("Expected argument 'checksum' to be a str")
        pulumi.set(__self__, "checksum", checksum)
        if context_id and not isinstance(context_id, str):
            raise TypeError("Expected argument 'context_id' to be a str")
        pulumi.set(__self__, "context_id", context_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)
        if write_only and not isinstance(write_only, bool):
            raise TypeError("Expected argument 'write_only' to be a bool")
        pulumi.set(__self__, "write_only", write_only)

    @property
    @pulumi.getter
    def checksum(self) -> str:
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter(name="contextId")
    def context_id(self) -> Optional[str]:
        return pulumi.get(self, "context_id")

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
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="writeOnly")
    def write_only(self) -> bool:
        return pulumi.get(self, "write_only")


class AwaitableGetEnvironmentVariableResult(GetEnvironmentVariableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentVariableResult(
            checksum=self.checksum,
            context_id=self.context_id,
            id=self.id,
            module_id=self.module_id,
            name=self.name,
            stack_id=self.stack_id,
            value=self.value,
            write_only=self.write_only)


def get_environment_variable(context_id: Optional[str] = None,
                             module_id: Optional[str] = None,
                             name: Optional[str] = None,
                             stack_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentVariableResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    __args__['moduleId'] = module_id
    __args__['name'] = name
    __args__['stackId'] = stack_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getEnvironmentVariable:getEnvironmentVariable', __args__, opts=opts, typ=GetEnvironmentVariableResult).value

    return AwaitableGetEnvironmentVariableResult(
        checksum=__ret__.checksum,
        context_id=__ret__.context_id,
        id=__ret__.id,
        module_id=__ret__.module_id,
        name=__ret__.name,
        stack_id=__ret__.stack_id,
        value=__ret__.value,
        write_only=__ret__.write_only)
