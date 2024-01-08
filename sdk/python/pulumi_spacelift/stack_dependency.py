# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StackDependencyArgs', 'StackDependency']

@pulumi.input_type
class StackDependencyArgs:
    def __init__(__self__, *,
                 depends_on_stack_id: pulumi.Input[str],
                 stack_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a StackDependency resource.
        :param pulumi.Input[str] depends_on_stack_id: immutable ID (slug) of stack to depend on.
        :param pulumi.Input[str] stack_id: immutable ID (slug) of stack which has a dependency.
        """
        StackDependencyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            depends_on_stack_id=depends_on_stack_id,
            stack_id=stack_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             depends_on_stack_id: Optional[pulumi.Input[str]] = None,
             stack_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if depends_on_stack_id is None and 'dependsOnStackId' in kwargs:
            depends_on_stack_id = kwargs['dependsOnStackId']
        if depends_on_stack_id is None:
            raise TypeError("Missing 'depends_on_stack_id' argument")
        if stack_id is None and 'stackId' in kwargs:
            stack_id = kwargs['stackId']
        if stack_id is None:
            raise TypeError("Missing 'stack_id' argument")

        _setter("depends_on_stack_id", depends_on_stack_id)
        _setter("stack_id", stack_id)

    @property
    @pulumi.getter(name="dependsOnStackId")
    def depends_on_stack_id(self) -> pulumi.Input[str]:
        """
        immutable ID (slug) of stack to depend on.
        """
        return pulumi.get(self, "depends_on_stack_id")

    @depends_on_stack_id.setter
    def depends_on_stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "depends_on_stack_id", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[str]:
        """
        immutable ID (slug) of stack which has a dependency.
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_id", value)


@pulumi.input_type
class _StackDependencyState:
    def __init__(__self__, *,
                 depends_on_stack_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StackDependency resources.
        :param pulumi.Input[str] depends_on_stack_id: immutable ID (slug) of stack to depend on.
        :param pulumi.Input[str] stack_id: immutable ID (slug) of stack which has a dependency.
        """
        _StackDependencyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            depends_on_stack_id=depends_on_stack_id,
            stack_id=stack_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             depends_on_stack_id: Optional[pulumi.Input[str]] = None,
             stack_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if depends_on_stack_id is None and 'dependsOnStackId' in kwargs:
            depends_on_stack_id = kwargs['dependsOnStackId']
        if stack_id is None and 'stackId' in kwargs:
            stack_id = kwargs['stackId']

        if depends_on_stack_id is not None:
            _setter("depends_on_stack_id", depends_on_stack_id)
        if stack_id is not None:
            _setter("stack_id", stack_id)

    @property
    @pulumi.getter(name="dependsOnStackId")
    def depends_on_stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        immutable ID (slug) of stack to depend on.
        """
        return pulumi.get(self, "depends_on_stack_id")

    @depends_on_stack_id.setter
    def depends_on_stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "depends_on_stack_id", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        immutable ID (slug) of stack which has a dependency.
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)


class StackDependency(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 depends_on_stack_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `StackDependency` represents a Spacelift **stack dependency** - a dependency between two stacks. When one stack depends on another, the tracked runs of the stack will not start until the dependent stack is successfully finished. Additionally, changes to the dependency will trigger the dependent.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        infra = spacelift.Stack("infra",
            branch="master",
            repository="core-infra")
        app = spacelift.Stack("app",
            branch="master",
            repository="app")
        test = spacelift.StackDependency("test",
            stack_id=app.id,
            depends_on_stack_id=infra.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] depends_on_stack_id: immutable ID (slug) of stack to depend on.
        :param pulumi.Input[str] stack_id: immutable ID (slug) of stack which has a dependency.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StackDependencyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `StackDependency` represents a Spacelift **stack dependency** - a dependency between two stacks. When one stack depends on another, the tracked runs of the stack will not start until the dependent stack is successfully finished. Additionally, changes to the dependency will trigger the dependent.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        infra = spacelift.Stack("infra",
            branch="master",
            repository="core-infra")
        app = spacelift.Stack("app",
            branch="master",
            repository="app")
        test = spacelift.StackDependency("test",
            stack_id=app.id,
            depends_on_stack_id=infra.id)
        ```

        :param str resource_name: The name of the resource.
        :param StackDependencyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StackDependencyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            StackDependencyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 depends_on_stack_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StackDependencyArgs.__new__(StackDependencyArgs)

            if depends_on_stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'depends_on_stack_id'")
            __props__.__dict__["depends_on_stack_id"] = depends_on_stack_id
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__.__dict__["stack_id"] = stack_id
        super(StackDependency, __self__).__init__(
            'spacelift:index/stackDependency:StackDependency',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            depends_on_stack_id: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None) -> 'StackDependency':
        """
        Get an existing StackDependency resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] depends_on_stack_id: immutable ID (slug) of stack to depend on.
        :param pulumi.Input[str] stack_id: immutable ID (slug) of stack which has a dependency.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StackDependencyState.__new__(_StackDependencyState)

        __props__.__dict__["depends_on_stack_id"] = depends_on_stack_id
        __props__.__dict__["stack_id"] = stack_id
        return StackDependency(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dependsOnStackId")
    def depends_on_stack_id(self) -> pulumi.Output[str]:
        """
        immutable ID (slug) of stack to depend on.
        """
        return pulumi.get(self, "depends_on_stack_id")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        immutable ID (slug) of stack which has a dependency.
        """
        return pulumi.get(self, "stack_id")
