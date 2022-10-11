# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StackDestructorArgs', 'StackDestructor']

@pulumi.input_type
class StackDestructorArgs:
    def __init__(__self__, *,
                 stack_id: pulumi.Input[str],
                 deactivated: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a StackDestructor resource.
        :param pulumi.Input[str] stack_id: ID of the stack to delete and destroy on destruction
        :param pulumi.Input[bool] deactivated: If set to true, destruction won't delete the stack
        """
        pulumi.set(__self__, "stack_id", stack_id)
        if deactivated is not None:
            pulumi.set(__self__, "deactivated", deactivated)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[str]:
        """
        ID of the stack to delete and destroy on destruction
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter
    def deactivated(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, destruction won't delete the stack
        """
        return pulumi.get(self, "deactivated")

    @deactivated.setter
    def deactivated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deactivated", value)


@pulumi.input_type
class _StackDestructorState:
    def __init__(__self__, *,
                 deactivated: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StackDestructor resources.
        :param pulumi.Input[bool] deactivated: If set to true, destruction won't delete the stack
        :param pulumi.Input[str] stack_id: ID of the stack to delete and destroy on destruction
        """
        if deactivated is not None:
            pulumi.set(__self__, "deactivated", deactivated)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)

    @property
    @pulumi.getter
    def deactivated(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, destruction won't delete the stack
        """
        return pulumi.get(self, "deactivated")

    @deactivated.setter
    def deactivated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deactivated", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack to delete and destroy on destruction
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)


class StackDestructor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deactivated: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `StackDestructor` is used to destroy the resources of a Stack before deleting it. `depends_on` should be used to make sure that all necessary resources (environment variables, roles, integrations, etc.) are still in place when the destruction run is executed. **Note:** Destroying this resource will delete the resources in the stack. If this resource needs to be deleted and the resources in the stacks are to be preserved, ensure that the `deactivated` attribute is set to `true`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        k8s_core_stack = spacelift.Stack("k8s-coreStack")
        # ...
        credentials = spacelift.EnvironmentVariable("credentials")
        # ...
        k8s_core_stack_destructor = spacelift.StackDestructor("k8s-coreStackDestructor", stack_id=k8s_core_stack.id,
        opts=pulumi.ResourceOptions(depends_on=[credentials]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deactivated: If set to true, destruction won't delete the stack
        :param pulumi.Input[str] stack_id: ID of the stack to delete and destroy on destruction
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StackDestructorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `StackDestructor` is used to destroy the resources of a Stack before deleting it. `depends_on` should be used to make sure that all necessary resources (environment variables, roles, integrations, etc.) are still in place when the destruction run is executed. **Note:** Destroying this resource will delete the resources in the stack. If this resource needs to be deleted and the resources in the stacks are to be preserved, ensure that the `deactivated` attribute is set to `true`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        k8s_core_stack = spacelift.Stack("k8s-coreStack")
        # ...
        credentials = spacelift.EnvironmentVariable("credentials")
        # ...
        k8s_core_stack_destructor = spacelift.StackDestructor("k8s-coreStackDestructor", stack_id=k8s_core_stack.id,
        opts=pulumi.ResourceOptions(depends_on=[credentials]))
        ```

        :param str resource_name: The name of the resource.
        :param StackDestructorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StackDestructorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deactivated: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StackDestructorArgs.__new__(StackDestructorArgs)

            __props__.__dict__["deactivated"] = deactivated
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
        super(StackDestructor, __self__).__init__(
            'spacelift:index/stackDestructor:StackDestructor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deactivated: Optional[pulumi.Input[bool]] = None,
            stack_id: Optional[pulumi.Input[str]] = None) -> 'StackDestructor':
        """
        Get an existing StackDestructor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deactivated: If set to true, destruction won't delete the stack
        :param pulumi.Input[str] stack_id: ID of the stack to delete and destroy on destruction
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StackDestructorState.__new__(_StackDestructorState)

        __props__.__dict__["deactivated"] = deactivated
        __props__.__dict__["stack_id"] = stack_id
        return StackDestructor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def deactivated(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to true, destruction won't delete the stack
        """
        return pulumi.get(self, "deactivated")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        ID of the stack to delete and destroy on destruction
        """
        return pulumi.get(self, "stack_id")

