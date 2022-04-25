# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['DriftDetection']


class DriftDetection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 reconcile: Optional[pulumi.Input[bool]] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        core_infra_production = spacelift.Stack("core-infra-production",
            name="Core Infrastructure (production)",
            branch="master",
            repository="core-infra")
        core_infra_production_drift_detection = spacelift.DriftDetection("core-infra-production-drift-detection",
            reconcile=True,
            stack_id=core_infra_production.id,
            schedules=["*/15 * * * *"])
        # Every 15 minutes
        ```

        <!-- schema generated by tfplugindocs -->
        ## Schema

        ### Required

        - **schedule** (List of String) List of cron schedule expressions based on which drift detection should be triggered.
        - **stack_id** (String) ID of the stack for which to set up drift detection

        ### Optional

        - **id** (String) The ID of this resource.
        - **reconcile** (Boolean) Whether a tracked run should be triggered when drift is detected.

        ## Import

        Import is supported using the following syntax

        ```sh
         $ pulumi import spacelift:index/driftDetection:DriftDetection core-infra-production-drift-detection stack/$STACK_ID
        ```

        ```sh
         $ pulumi import spacelift:index/driftDetection:DriftDetection core-infra-production-drift-detection module/$MODULE_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] reconcile: Whether a tracked run should be triggered when drift is detected.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] schedules: List of cron schedule expressions based on which drift detection should be triggered.
        :param pulumi.Input[str] stack_id: ID of the stack for which to set up drift detection
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['reconcile'] = reconcile
            if schedules is None and not opts.urn:
                raise TypeError("Missing required property 'schedules'")
            __props__['schedules'] = schedules
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__['stack_id'] = stack_id
        super(DriftDetection, __self__).__init__(
            'spacelift:index/driftDetection:DriftDetection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            reconcile: Optional[pulumi.Input[bool]] = None,
            schedules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            stack_id: Optional[pulumi.Input[str]] = None) -> 'DriftDetection':
        """
        Get an existing DriftDetection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] reconcile: Whether a tracked run should be triggered when drift is detected.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] schedules: List of cron schedule expressions based on which drift detection should be triggered.
        :param pulumi.Input[str] stack_id: ID of the stack for which to set up drift detection
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["reconcile"] = reconcile
        __props__["schedules"] = schedules
        __props__["stack_id"] = stack_id
        return DriftDetection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def reconcile(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether a tracked run should be triggered when drift is detected.
        """
        return pulumi.get(self, "reconcile")

    @property
    @pulumi.getter
    def schedules(self) -> pulumi.Output[Sequence[str]]:
        """
        List of cron schedule expressions based on which drift detection should be triggered.
        """
        return pulumi.get(self, "schedules")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        ID of the stack for which to set up drift detection
        """
        return pulumi.get(self, "stack_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
