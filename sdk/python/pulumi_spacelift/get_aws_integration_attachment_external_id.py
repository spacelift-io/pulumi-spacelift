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
    'GetAwsIntegrationAttachmentExternalIdResult',
    'AwaitableGetAwsIntegrationAttachmentExternalIdResult',
    'get_aws_integration_attachment_external_id',
    'get_aws_integration_attachment_external_id_output',
]

@pulumi.output_type
class GetAwsIntegrationAttachmentExternalIdResult:
    """
    A collection of values returned by getAwsIntegrationAttachmentExternalId.
    """
    def __init__(__self__, assume_role_policy_statement=None, external_id=None, id=None, integration_id=None, module_id=None, read=None, stack_id=None, write=None):
        if assume_role_policy_statement and not isinstance(assume_role_policy_statement, str):
            raise TypeError("Expected argument 'assume_role_policy_statement' to be a str")
        pulumi.set(__self__, "assume_role_policy_statement", assume_role_policy_statement)
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        pulumi.set(__self__, "external_id", external_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if integration_id and not isinstance(integration_id, str):
            raise TypeError("Expected argument 'integration_id' to be a str")
        pulumi.set(__self__, "integration_id", integration_id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if read and not isinstance(read, bool):
            raise TypeError("Expected argument 'read' to be a bool")
        pulumi.set(__self__, "read", read)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if write and not isinstance(write, bool):
            raise TypeError("Expected argument 'write' to be a bool")
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter(name="assumeRolePolicyStatement")
    def assume_role_policy_statement(self) -> str:
        """
        An assume role policy statement that can be attached to your role to allow Spacelift to assume it
        """
        return pulumi.get(self, "assume_role_policy_statement")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        """
        The external ID that will be used during role assumption
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> str:
        """
        immutable ID (slug) of the AWS integration
        """
        return pulumi.get(self, "integration_id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[str]:
        """
        immutable ID (slug) of the module
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def read(self) -> Optional[bool]:
        """
        whether the integration will be used for read operations
        """
        return pulumi.get(self, "read")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        """
        immutable ID (slug) of the stack
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def write(self) -> Optional[bool]:
        """
        whether the integration will be used for write operations
        """
        return pulumi.get(self, "write")


class AwaitableGetAwsIntegrationAttachmentExternalIdResult(GetAwsIntegrationAttachmentExternalIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsIntegrationAttachmentExternalIdResult(
            assume_role_policy_statement=self.assume_role_policy_statement,
            external_id=self.external_id,
            id=self.id,
            integration_id=self.integration_id,
            module_id=self.module_id,
            read=self.read,
            stack_id=self.stack_id,
            write=self.write)


def get_aws_integration_attachment_external_id(integration_id: Optional[str] = None,
                                               module_id: Optional[str] = None,
                                               read: Optional[bool] = None,
                                               stack_id: Optional[str] = None,
                                               write: Optional[bool] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsIntegrationAttachmentExternalIdResult:
    """
    `get_aws_integration_attachment_external_id` is used to generate the external ID that would be used for role assumption when an AWS integration is attached to a stack or module.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    my_stack = spacelift.get_aws_integration_attachment_external_id(integration_id=spacelift_aws_integration["this"]["id"],
        stack_id="my-stack-id",
        read=True,
        write=True)
    my_module = spacelift.get_aws_integration_attachment_external_id(integration_id=spacelift_aws_integration["this"]["id"],
        module_id="my-module-id",
        read=True,
        write=True)
    ```


    :param str integration_id: immutable ID (slug) of the AWS integration
    :param str module_id: immutable ID (slug) of the module
    :param bool read: whether the integration will be used for read operations
    :param str stack_id: immutable ID (slug) of the stack
    :param bool write: whether the integration will be used for write operations
    """
    __args__ = dict()
    __args__['integrationId'] = integration_id
    __args__['moduleId'] = module_id
    __args__['read'] = read
    __args__['stackId'] = stack_id
    __args__['write'] = write
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getAwsIntegrationAttachmentExternalId:getAwsIntegrationAttachmentExternalId', __args__, opts=opts, typ=GetAwsIntegrationAttachmentExternalIdResult).value

    return AwaitableGetAwsIntegrationAttachmentExternalIdResult(
        assume_role_policy_statement=pulumi.get(__ret__, 'assume_role_policy_statement'),
        external_id=pulumi.get(__ret__, 'external_id'),
        id=pulumi.get(__ret__, 'id'),
        integration_id=pulumi.get(__ret__, 'integration_id'),
        module_id=pulumi.get(__ret__, 'module_id'),
        read=pulumi.get(__ret__, 'read'),
        stack_id=pulumi.get(__ret__, 'stack_id'),
        write=pulumi.get(__ret__, 'write'))


@_utilities.lift_output_func(get_aws_integration_attachment_external_id)
def get_aws_integration_attachment_external_id_output(integration_id: Optional[pulumi.Input[str]] = None,
                                                      module_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                      read: Optional[pulumi.Input[Optional[bool]]] = None,
                                                      stack_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                      write: Optional[pulumi.Input[Optional[bool]]] = None,
                                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAwsIntegrationAttachmentExternalIdResult]:
    """
    `get_aws_integration_attachment_external_id` is used to generate the external ID that would be used for role assumption when an AWS integration is attached to a stack or module.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    my_stack = spacelift.get_aws_integration_attachment_external_id(integration_id=spacelift_aws_integration["this"]["id"],
        stack_id="my-stack-id",
        read=True,
        write=True)
    my_module = spacelift.get_aws_integration_attachment_external_id(integration_id=spacelift_aws_integration["this"]["id"],
        module_id="my-module-id",
        read=True,
        write=True)
    ```


    :param str integration_id: immutable ID (slug) of the AWS integration
    :param str module_id: immutable ID (slug) of the module
    :param bool read: whether the integration will be used for read operations
    :param str stack_id: immutable ID (slug) of the stack
    :param bool write: whether the integration will be used for write operations
    """
    ...
