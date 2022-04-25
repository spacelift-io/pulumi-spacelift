# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['EnvironmentVariable']


class EnvironmentVariable(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context_id: Optional[pulumi.Input[str]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 write_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `EnvironmentVariable` defines an environment variable on the context (`Context`), stack (`Stack`) or a module (`Module`), thereby allowing to pass and share various secrets and configuration with Spacelift stacks.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        # For a context
        ireland_kubeconfig = spacelift.EnvironmentVariable("ireland-kubeconfig",
            context_id="prod-k8s-ie",
            name="KUBECONFIG",
            value="/project/spacelift/kubeconfig",
            write_only=False)
        # For a module
        module_kubeconfig = spacelift.EnvironmentVariable("module-kubeconfig",
            module_id="k8s-module",
            name="KUBECONFIG",
            value="/project/spacelift/kubeconfig",
            write_only=False)
        # For a stack
        core_kubeconfig = spacelift.EnvironmentVariable("core-kubeconfig",
            name="KUBECONFIG",
            stack_id="k8s-core",
            value="/project/spacelift/kubeconfig",
            write_only=False)
        ```

        <!-- schema generated by tfplugindocs -->
        ## Schema

        ### Required

        - **name** (String) Name of the environment variable
        - **value** (String, Sensitive) Value of the environment variable

        ### Optional

        - **context_id** (String) ID of the context on which the environment variable is defined
        - **id** (String) The ID of this resource.
        - **module_id** (String) ID of the module on which the environment variable is defined
        - **stack_id** (String) ID of the stack on which the environment variable is defined
        - **write_only** (Boolean) Indicates whether the value can be read back outside a Run

        ### Read-Only

        - **checksum** (String) SHA-256 checksum of the value

        ## Import

        Import is supported using the following syntax

        ```sh
         $ pulumi import spacelift:index/environmentVariable:EnvironmentVariable ireland-kubeconfig context/$CONTEXT_ID/$ENVIRONMENT_VARIABLE_NAME
        ```

        ```sh
         $ pulumi import spacelift:index/environmentVariable:EnvironmentVariable module-kubeconfig module/$MODULE_ID/$ENVIRONMENT_VARIABLE_NAME
        ```

        ```sh
         $ pulumi import spacelift:index/environmentVariable:EnvironmentVariable core-kubeconfig stack/$STACK_ID/$ENVIRONMENT_VARIABLE_NAME
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context_id: ID of the context on which the environment variable is defined
        :param pulumi.Input[str] module_id: ID of the module on which the environment variable is defined
        :param pulumi.Input[str] name: Name of the environment variable
        :param pulumi.Input[str] stack_id: ID of the stack on which the environment variable is defined
        :param pulumi.Input[str] value: Value of the environment variable
        :param pulumi.Input[bool] write_only: Indicates whether the value can be read back outside a Run
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

            __props__['context_id'] = context_id
            __props__['module_id'] = module_id
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['stack_id'] = stack_id
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__['value'] = value
            __props__['write_only'] = write_only
            __props__['checksum'] = None
        super(EnvironmentVariable, __self__).__init__(
            'spacelift:index/environmentVariable:EnvironmentVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            checksum: Optional[pulumi.Input[str]] = None,
            context_id: Optional[pulumi.Input[str]] = None,
            module_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            write_only: Optional[pulumi.Input[bool]] = None) -> 'EnvironmentVariable':
        """
        Get an existing EnvironmentVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] checksum: SHA-256 checksum of the value
        :param pulumi.Input[str] context_id: ID of the context on which the environment variable is defined
        :param pulumi.Input[str] module_id: ID of the module on which the environment variable is defined
        :param pulumi.Input[str] name: Name of the environment variable
        :param pulumi.Input[str] stack_id: ID of the stack on which the environment variable is defined
        :param pulumi.Input[str] value: Value of the environment variable
        :param pulumi.Input[bool] write_only: Indicates whether the value can be read back outside a Run
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["checksum"] = checksum
        __props__["context_id"] = context_id
        __props__["module_id"] = module_id
        __props__["name"] = name
        __props__["stack_id"] = stack_id
        __props__["value"] = value
        __props__["write_only"] = write_only
        return EnvironmentVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def checksum(self) -> pulumi.Output[str]:
        """
        SHA-256 checksum of the value
        """
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter(name="contextId")
    def context_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the context on which the environment variable is defined
        """
        return pulumi.get(self, "context_id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the module on which the environment variable is defined
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the environment variable
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the stack on which the environment variable is defined
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Value of the environment variable
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="writeOnly")
    def write_only(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the value can be read back outside a Run
        """
        return pulumi.get(self, "write_only")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

