# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetStackAwsRoleResult',
    'AwaitableGetStackAwsRoleResult',
    'get_stack_aws_role',
]

@pulumi.output_type
class GetStackAwsRoleResult:
    """
    A collection of values returned by getStackAwsRole.
    """
    def __init__(__self__, id=None, module_id=None, role_arn=None, stack_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)

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
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        return pulumi.get(self, "stack_id")


class AwaitableGetStackAwsRoleResult(GetStackAwsRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStackAwsRoleResult(
            id=self.id,
            module_id=self.module_id,
            role_arn=self.role_arn,
            stack_id=self.stack_id)


def get_stack_aws_role(module_id: Optional[str] = None,
                       stack_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStackAwsRoleResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['moduleId'] = module_id
    __args__['stackId'] = stack_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getStackAwsRole:getStackAwsRole', __args__, opts=opts, typ=GetStackAwsRoleResult).value

    return AwaitableGetStackAwsRoleResult(
        id=__ret__.id,
        module_id=__ret__.module_id,
        role_arn=__ret__.role_arn,
        stack_id=__ret__.stack_id)