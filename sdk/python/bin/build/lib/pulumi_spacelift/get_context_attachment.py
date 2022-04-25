# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetContextAttachmentResult',
    'AwaitableGetContextAttachmentResult',
    'get_context_attachment',
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
    def priority(self) -> int:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
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

    <!-- schema generated by tfplugindocs -->
    ## Schema

    ### Required

    - **context_id** (String) ID of the attached context

    ### Optional

    - **id** (String) The ID of this resource.
    - **module_id** (String) ID of the attached module
    - **stack_id** (String) ID of the attached stack

    ### Read-Only

    - **priority** (Number) priority of the context attachment, used in case of conflicts
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    __args__['moduleId'] = module_id
    __args__['stackId'] = stack_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getContextAttachment:getContextAttachment', __args__, opts=opts, typ=GetContextAttachmentResult).value

    return AwaitableGetContextAttachmentResult(
        context_id=__ret__.context_id,
        id=__ret__.id,
        module_id=__ret__.module_id,
        priority=__ret__.priority,
        stack_id=__ret__.stack_id)
