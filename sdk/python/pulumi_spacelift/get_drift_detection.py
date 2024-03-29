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
    'GetDriftDetectionResult',
    'AwaitableGetDriftDetectionResult',
    'get_drift_detection',
    'get_drift_detection_output',
]

@pulumi.output_type
class GetDriftDetectionResult:
    """
    A collection of values returned by getDriftDetection.
    """
    def __init__(__self__, id=None, ignore_state=None, reconcile=None, schedules=None, stack_id=None, timezone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_state and not isinstance(ignore_state, bool):
            raise TypeError("Expected argument 'ignore_state' to be a bool")
        pulumi.set(__self__, "ignore_state", ignore_state)
        if reconcile and not isinstance(reconcile, bool):
            raise TypeError("Expected argument 'reconcile' to be a bool")
        pulumi.set(__self__, "reconcile", reconcile)
        if schedules and not isinstance(schedules, list):
            raise TypeError("Expected argument 'schedules' to be a list")
        pulumi.set(__self__, "schedules", schedules)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if timezone and not isinstance(timezone, str):
            raise TypeError("Expected argument 'timezone' to be a str")
        pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreState")
    def ignore_state(self) -> Optional[bool]:
        """
        Controls whether drift detection should be performed on a stack in any final state instead of just 'Finished'.
        """
        return pulumi.get(self, "ignore_state")

    @property
    @pulumi.getter
    def reconcile(self) -> bool:
        """
        Whether a tracked run should be triggered when drift is detected.
        """
        return pulumi.get(self, "reconcile")

    @property
    @pulumi.getter
    def schedules(self) -> Sequence[str]:
        """
        List of cron schedule expressions based on which drift detection should be triggered.
        """
        return pulumi.get(self, "schedules")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        ID of the stack for which to set up drift detection
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def timezone(self) -> str:
        """
        Timezone in which the schedule is expressed
        """
        return pulumi.get(self, "timezone")


class AwaitableGetDriftDetectionResult(GetDriftDetectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDriftDetectionResult(
            id=self.id,
            ignore_state=self.ignore_state,
            reconcile=self.reconcile,
            schedules=self.schedules,
            stack_id=self.stack_id,
            timezone=self.timezone)


def get_drift_detection(ignore_state: Optional[bool] = None,
                        stack_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDriftDetectionResult:
    """
    `DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    core_infra_production_drift_detection = spacelift.get_drift_detection(stack_id="core-infra-production")
    ```


    :param bool ignore_state: Controls whether drift detection should be performed on a stack in any final state instead of just 'Finished'.
    :param str stack_id: ID of the stack for which to set up drift detection
    """
    __args__ = dict()
    __args__['ignoreState'] = ignore_state
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getDriftDetection:getDriftDetection', __args__, opts=opts, typ=GetDriftDetectionResult).value

    return AwaitableGetDriftDetectionResult(
        id=pulumi.get(__ret__, 'id'),
        ignore_state=pulumi.get(__ret__, 'ignore_state'),
        reconcile=pulumi.get(__ret__, 'reconcile'),
        schedules=pulumi.get(__ret__, 'schedules'),
        stack_id=pulumi.get(__ret__, 'stack_id'),
        timezone=pulumi.get(__ret__, 'timezone'))


@_utilities.lift_output_func(get_drift_detection)
def get_drift_detection_output(ignore_state: Optional[pulumi.Input[Optional[bool]]] = None,
                               stack_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDriftDetectionResult]:
    """
    `DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    core_infra_production_drift_detection = spacelift.get_drift_detection(stack_id="core-infra-production")
    ```


    :param bool ignore_state: Controls whether drift detection should be performed on a stack in any final state instead of just 'Finished'.
    :param str stack_id: ID of the stack for which to set up drift detection
    """
    ...
