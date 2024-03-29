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
    'GetScheduledDeleteStackResult',
    'AwaitableGetScheduledDeleteStackResult',
    'get_scheduled_delete_stack',
    'get_scheduled_delete_stack_output',
]

@pulumi.output_type
class GetScheduledDeleteStackResult:
    """
    A collection of values returned by getScheduledDeleteStack.
    """
    def __init__(__self__, at=None, delete_resources=None, id=None, schedule_id=None, scheduled_delete_stack_id=None, stack_id=None):
        if at and not isinstance(at, int):
            raise TypeError("Expected argument 'at' to be a int")
        pulumi.set(__self__, "at", at)
        if delete_resources and not isinstance(delete_resources, bool):
            raise TypeError("Expected argument 'delete_resources' to be a bool")
        pulumi.set(__self__, "delete_resources", delete_resources)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if schedule_id and not isinstance(schedule_id, str):
            raise TypeError("Expected argument 'schedule_id' to be a str")
        pulumi.set(__self__, "schedule_id", schedule_id)
        if scheduled_delete_stack_id and not isinstance(scheduled_delete_stack_id, str):
            raise TypeError("Expected argument 'scheduled_delete_stack_id' to be a str")
        pulumi.set(__self__, "scheduled_delete_stack_id", scheduled_delete_stack_id)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)

    @property
    @pulumi.getter
    def at(self) -> int:
        """
        Timestamp (unix timestamp) at which time the scheduling should happen.
        """
        return pulumi.get(self, "at")

    @property
    @pulumi.getter(name="deleteResources")
    def delete_resources(self) -> bool:
        """
        Indicates whether the resources of the stack should be deleted.
        """
        return pulumi.get(self, "delete_resources")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> str:
        """
        ID of the schedule
        """
        return pulumi.get(self, "schedule_id")

    @property
    @pulumi.getter(name="scheduledDeleteStackId")
    def scheduled_delete_stack_id(self) -> str:
        """
        ID of the scheduled delete*stack (stack*id/schedule_id)
        """
        return pulumi.get(self, "scheduled_delete_stack_id")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        Stack ID of the scheduling config
        """
        return pulumi.get(self, "stack_id")


class AwaitableGetScheduledDeleteStackResult(GetScheduledDeleteStackResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScheduledDeleteStackResult(
            at=self.at,
            delete_resources=self.delete_resources,
            id=self.id,
            schedule_id=self.schedule_id,
            scheduled_delete_stack_id=self.scheduled_delete_stack_id,
            stack_id=self.stack_id)


def get_scheduled_delete_stack(scheduled_delete_stack_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScheduledDeleteStackResult:
    """
    `ScheduledDeleteTask` represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    ireland_kubeconfig_delete = spacelift.get_scheduled_delete_stack(scheduled_delete_stack_id="$STACK_ID/$SCHEDULED_DELETE_STACK_ID")
    ```


    :param str scheduled_delete_stack_id: ID of the scheduled delete*stack (stack*id/schedule_id)
    """
    __args__ = dict()
    __args__['scheduledDeleteStackId'] = scheduled_delete_stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getScheduledDeleteStack:getScheduledDeleteStack', __args__, opts=opts, typ=GetScheduledDeleteStackResult).value

    return AwaitableGetScheduledDeleteStackResult(
        at=pulumi.get(__ret__, 'at'),
        delete_resources=pulumi.get(__ret__, 'delete_resources'),
        id=pulumi.get(__ret__, 'id'),
        schedule_id=pulumi.get(__ret__, 'schedule_id'),
        scheduled_delete_stack_id=pulumi.get(__ret__, 'scheduled_delete_stack_id'),
        stack_id=pulumi.get(__ret__, 'stack_id'))


@_utilities.lift_output_func(get_scheduled_delete_stack)
def get_scheduled_delete_stack_output(scheduled_delete_stack_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScheduledDeleteStackResult]:
    """
    `ScheduledDeleteTask` represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    ireland_kubeconfig_delete = spacelift.get_scheduled_delete_stack(scheduled_delete_stack_id="$STACK_ID/$SCHEDULED_DELETE_STACK_ID")
    ```


    :param str scheduled_delete_stack_id: ID of the scheduled delete*stack (stack*id/schedule_id)
    """
    ...
