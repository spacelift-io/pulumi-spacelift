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
    'GetContextAttachmentResult',
    'AwaitableGetContextAttachmentResult',
    'get_context_attachment',
    'get_context_attachment_output',
]

@pulumi.output_type
class GetContextAttachmentResult:
    """
    A collection of values returned by getContextAttachment.
    """
    def __init__(__self__, context_id=None, id=None, module_id=None, priority=None, stack_id=None):
        if context_id and not isinstance(context_id, str):
            raise TypeError("Expected argument 'context_id' to be a str")
        pulumi.set(__self__, "context_id", context_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)

    @property
    @pulumi.getter(name="contextId")
    def context_id(self) -> str:
        """
        ID of the attached context
        """
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
        """
        ID of the attached module
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Priority of the context attachment. All the contexts attached to a stack are sorted by priority (lowest first), though values don't need to be unique. This ordering establishes precedence rules between contexts should there be a conflict and multiple contexts define the same value.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        """
        ID of the attached stack
        """
        return pulumi.get(self, "stack_id")


class AwaitableGetContextAttachmentResult(GetContextAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContextAttachmentResult(
            context_id=self.context_id,
            id=self.id,
            module_id=self.module_id,
            priority=self.priority,
            stack_id=self.stack_id)


def get_context_attachment(context_id: Optional[str] = None,
                           module_id: Optional[str] = None,
                           stack_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContextAttachmentResult:
    """
    `ContextAttachment` represents a Spacelift attachment of a single context to a single stack or module, with a predefined priority.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    apps_k8s_ie = spacelift.get_context_attachment(context_id="prod-k8s-ie",
        stack_id="apps-cluster")
    kafka_k8s_ie = spacelift.get_context_attachment(context_id="prod-k8s-ie",
        module_id="terraform-aws-kafka")
    ```


    :param str context_id: ID of the attached context
    :param str module_id: ID of the attached module
    :param str stack_id: ID of the attached stack
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    __args__['moduleId'] = module_id
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getContextAttachment:getContextAttachment', __args__, opts=opts, typ=GetContextAttachmentResult).value

    return AwaitableGetContextAttachmentResult(
        context_id=pulumi.get(__ret__, 'context_id'),
        id=pulumi.get(__ret__, 'id'),
        module_id=pulumi.get(__ret__, 'module_id'),
        priority=pulumi.get(__ret__, 'priority'),
        stack_id=pulumi.get(__ret__, 'stack_id'))


@_utilities.lift_output_func(get_context_attachment)
def get_context_attachment_output(context_id: Optional[pulumi.Input[str]] = None,
                                  module_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  stack_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContextAttachmentResult]:
    """
    `ContextAttachment` represents a Spacelift attachment of a single context to a single stack or module, with a predefined priority.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    apps_k8s_ie = spacelift.get_context_attachment(context_id="prod-k8s-ie",
        stack_id="apps-cluster")
    kafka_k8s_ie = spacelift.get_context_attachment(context_id="prod-k8s-ie",
        module_id="terraform-aws-kafka")
    ```


    :param str context_id: ID of the attached context
    :param str module_id: ID of the attached module
    :param str stack_id: ID of the attached stack
    """
    ...
