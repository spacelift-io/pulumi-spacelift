# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['ContextAttachment']


class ContextAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context_id: Optional[pulumi.Input[str]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ContextAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context_id: ID of the context to attach
        :param pulumi.Input[str] module_id: ID of the module to attach the context to
        :param pulumi.Input[int] priority: Priority of the context attachment, used in case of conflicts
        :param pulumi.Input[str] stack_id: ID of the stack to attach the context to
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

            if context_id is None and not opts.urn:
                raise TypeError("Missing required property 'context_id'")
            __props__['context_id'] = context_id
            __props__['module_id'] = module_id
            __props__['priority'] = priority
            __props__['stack_id'] = stack_id
        super(ContextAttachment, __self__).__init__(
            'spacelift:index/contextAttachment:ContextAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            context_id: Optional[pulumi.Input[str]] = None,
            module_id: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            stack_id: Optional[pulumi.Input[str]] = None) -> 'ContextAttachment':
        """
        Get an existing ContextAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context_id: ID of the context to attach
        :param pulumi.Input[str] module_id: ID of the module to attach the context to
        :param pulumi.Input[int] priority: Priority of the context attachment, used in case of conflicts
        :param pulumi.Input[str] stack_id: ID of the stack to attach the context to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["context_id"] = context_id
        __props__["module_id"] = module_id
        __props__["priority"] = priority
        __props__["stack_id"] = stack_id
        return ContextAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contextId")
    def context_id(self) -> pulumi.Output[str]:
        """
        ID of the context to attach
        """
        return pulumi.get(self, "context_id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the module to attach the context to
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        Priority of the context attachment, used in case of conflicts
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the stack to attach the context to
        """
        return pulumi.get(self, "stack_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
