# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StackDependencyReferenceArgs', 'StackDependencyReference']

@pulumi.input_type
class StackDependencyReferenceArgs:
    def __init__(__self__, *,
                 input_name: pulumi.Input[str],
                 output_name: pulumi.Input[str],
                 stack_dependency_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a StackDependencyReference resource.
        :param pulumi.Input[str] input_name: Name of the input of the stack dependency reference
        :param pulumi.Input[str] output_name: Name of the output of stack to depend on
        :param pulumi.Input[str] stack_dependency_id: Immutable ID of stack dependency
        """
        StackDependencyReferenceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            input_name=input_name,
            output_name=output_name,
            stack_dependency_id=stack_dependency_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             input_name: Optional[pulumi.Input[str]] = None,
             output_name: Optional[pulumi.Input[str]] = None,
             stack_dependency_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if input_name is None and 'inputName' in kwargs:
            input_name = kwargs['inputName']
        if input_name is None:
            raise TypeError("Missing 'input_name' argument")
        if output_name is None and 'outputName' in kwargs:
            output_name = kwargs['outputName']
        if output_name is None:
            raise TypeError("Missing 'output_name' argument")
        if stack_dependency_id is None and 'stackDependencyId' in kwargs:
            stack_dependency_id = kwargs['stackDependencyId']
        if stack_dependency_id is None:
            raise TypeError("Missing 'stack_dependency_id' argument")

        _setter("input_name", input_name)
        _setter("output_name", output_name)
        _setter("stack_dependency_id", stack_dependency_id)

    @property
    @pulumi.getter(name="inputName")
    def input_name(self) -> pulumi.Input[str]:
        """
        Name of the input of the stack dependency reference
        """
        return pulumi.get(self, "input_name")

    @input_name.setter
    def input_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "input_name", value)

    @property
    @pulumi.getter(name="outputName")
    def output_name(self) -> pulumi.Input[str]:
        """
        Name of the output of stack to depend on
        """
        return pulumi.get(self, "output_name")

    @output_name.setter
    def output_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "output_name", value)

    @property
    @pulumi.getter(name="stackDependencyId")
    def stack_dependency_id(self) -> pulumi.Input[str]:
        """
        Immutable ID of stack dependency
        """
        return pulumi.get(self, "stack_dependency_id")

    @stack_dependency_id.setter
    def stack_dependency_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_dependency_id", value)


@pulumi.input_type
class _StackDependencyReferenceState:
    def __init__(__self__, *,
                 input_name: Optional[pulumi.Input[str]] = None,
                 output_name: Optional[pulumi.Input[str]] = None,
                 stack_dependency_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StackDependencyReference resources.
        :param pulumi.Input[str] input_name: Name of the input of the stack dependency reference
        :param pulumi.Input[str] output_name: Name of the output of stack to depend on
        :param pulumi.Input[str] stack_dependency_id: Immutable ID of stack dependency
        """
        _StackDependencyReferenceState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            input_name=input_name,
            output_name=output_name,
            stack_dependency_id=stack_dependency_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             input_name: Optional[pulumi.Input[str]] = None,
             output_name: Optional[pulumi.Input[str]] = None,
             stack_dependency_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if input_name is None and 'inputName' in kwargs:
            input_name = kwargs['inputName']
        if output_name is None and 'outputName' in kwargs:
            output_name = kwargs['outputName']
        if stack_dependency_id is None and 'stackDependencyId' in kwargs:
            stack_dependency_id = kwargs['stackDependencyId']

        if input_name is not None:
            _setter("input_name", input_name)
        if output_name is not None:
            _setter("output_name", output_name)
        if stack_dependency_id is not None:
            _setter("stack_dependency_id", stack_dependency_id)

    @property
    @pulumi.getter(name="inputName")
    def input_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the input of the stack dependency reference
        """
        return pulumi.get(self, "input_name")

    @input_name.setter
    def input_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_name", value)

    @property
    @pulumi.getter(name="outputName")
    def output_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the output of stack to depend on
        """
        return pulumi.get(self, "output_name")

    @output_name.setter
    def output_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_name", value)

    @property
    @pulumi.getter(name="stackDependencyId")
    def stack_dependency_id(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable ID of stack dependency
        """
        return pulumi.get(self, "stack_dependency_id")

    @stack_dependency_id.setter
    def stack_dependency_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_dependency_id", value)


class StackDependencyReference(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_name: Optional[pulumi.Input[str]] = None,
                 output_name: Optional[pulumi.Input[str]] = None,
                 stack_dependency_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `StackDependencyReference` represents a Spacelift **stack dependency reference** - a reference matches a stack's output to another stack's input. It is similar to an environment variable (`EnvironmentVariable`), except that value is provided by another stack's output.

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
        test_stack_dependency = spacelift.StackDependency("testStackDependency",
            stack_id=app.id,
            depends_on_stack_id=infra.id)
        test_stack_dependency_reference = spacelift.StackDependencyReference("testStackDependencyReference",
            stack_dependency_id=test_stack_dependency.id,
            output_name="DB_CONNECTION_STRING",
            input_name="APP_DB_URL")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] input_name: Name of the input of the stack dependency reference
        :param pulumi.Input[str] output_name: Name of the output of stack to depend on
        :param pulumi.Input[str] stack_dependency_id: Immutable ID of stack dependency
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StackDependencyReferenceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `StackDependencyReference` represents a Spacelift **stack dependency reference** - a reference matches a stack's output to another stack's input. It is similar to an environment variable (`EnvironmentVariable`), except that value is provided by another stack's output.

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
        test_stack_dependency = spacelift.StackDependency("testStackDependency",
            stack_id=app.id,
            depends_on_stack_id=infra.id)
        test_stack_dependency_reference = spacelift.StackDependencyReference("testStackDependencyReference",
            stack_dependency_id=test_stack_dependency.id,
            output_name="DB_CONNECTION_STRING",
            input_name="APP_DB_URL")
        ```

        :param str resource_name: The name of the resource.
        :param StackDependencyReferenceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StackDependencyReferenceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            StackDependencyReferenceArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_name: Optional[pulumi.Input[str]] = None,
                 output_name: Optional[pulumi.Input[str]] = None,
                 stack_dependency_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StackDependencyReferenceArgs.__new__(StackDependencyReferenceArgs)

            if input_name is None and not opts.urn:
                raise TypeError("Missing required property 'input_name'")
            __props__.__dict__["input_name"] = input_name
            if output_name is None and not opts.urn:
                raise TypeError("Missing required property 'output_name'")
            __props__.__dict__["output_name"] = output_name
            if stack_dependency_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_dependency_id'")
            __props__.__dict__["stack_dependency_id"] = stack_dependency_id
        super(StackDependencyReference, __self__).__init__(
            'spacelift:index/stackDependencyReference:StackDependencyReference',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            input_name: Optional[pulumi.Input[str]] = None,
            output_name: Optional[pulumi.Input[str]] = None,
            stack_dependency_id: Optional[pulumi.Input[str]] = None) -> 'StackDependencyReference':
        """
        Get an existing StackDependencyReference resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] input_name: Name of the input of the stack dependency reference
        :param pulumi.Input[str] output_name: Name of the output of stack to depend on
        :param pulumi.Input[str] stack_dependency_id: Immutable ID of stack dependency
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StackDependencyReferenceState.__new__(_StackDependencyReferenceState)

        __props__.__dict__["input_name"] = input_name
        __props__.__dict__["output_name"] = output_name
        __props__.__dict__["stack_dependency_id"] = stack_dependency_id
        return StackDependencyReference(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="inputName")
    def input_name(self) -> pulumi.Output[str]:
        """
        Name of the input of the stack dependency reference
        """
        return pulumi.get(self, "input_name")

    @property
    @pulumi.getter(name="outputName")
    def output_name(self) -> pulumi.Output[str]:
        """
        Name of the output of stack to depend on
        """
        return pulumi.get(self, "output_name")

    @property
    @pulumi.getter(name="stackDependencyId")
    def stack_dependency_id(self) -> pulumi.Output[str]:
        """
        Immutable ID of stack dependency
        """
        return pulumi.get(self, "stack_dependency_id")
