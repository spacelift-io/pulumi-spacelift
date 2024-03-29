# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VcsAgentPoolArgs', 'VcsAgentPool']

@pulumi.input_type
class VcsAgentPoolArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VcsAgentPool resource.
        :param pulumi.Input[str] description: Free-form VCS agent pool description for users
        :param pulumi.Input[str] name: Name of the VCS agent pool, must be unique within an account
        """
        VcsAgentPoolArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-form VCS agent pool description for users
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the VCS agent pool, must be unique within an account
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _VcsAgentPoolState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VcsAgentPool resources.
        :param pulumi.Input[str] config: VCS agent pool configuration, encoded using base64
        :param pulumi.Input[str] description: Free-form VCS agent pool description for users
        :param pulumi.Input[str] name: Name of the VCS agent pool, must be unique within an account
        """
        _VcsAgentPoolState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            config=config,
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             config: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if config is not None:
            _setter("config", config)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        VCS agent pool configuration, encoded using base64
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-form VCS agent pool description for users
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the VCS agent pool, must be unique within an account
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class VcsAgentPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `VcsAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        ghe = spacelift.VcsAgentPool("ghe", description="VCS agent pool for our internal GitHub Enterprise")
        ```

        ## Import

        ```sh
         $ pulumi import spacelift:index/vcsAgentPool:VcsAgentPool ghe $VCS_AGENT_POOL_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Free-form VCS agent pool description for users
        :param pulumi.Input[str] name: Name of the VCS agent pool, must be unique within an account
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VcsAgentPoolArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `VcsAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        ghe = spacelift.VcsAgentPool("ghe", description="VCS agent pool for our internal GitHub Enterprise")
        ```

        ## Import

        ```sh
         $ pulumi import spacelift:index/vcsAgentPool:VcsAgentPool ghe $VCS_AGENT_POOL_ID
        ```

        :param str resource_name: The name of the resource.
        :param VcsAgentPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VcsAgentPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VcsAgentPoolArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VcsAgentPoolArgs.__new__(VcsAgentPoolArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["config"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["config"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(VcsAgentPool, __self__).__init__(
            'spacelift:index/vcsAgentPool:VcsAgentPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'VcsAgentPool':
        """
        Get an existing VcsAgentPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: VCS agent pool configuration, encoded using base64
        :param pulumi.Input[str] description: Free-form VCS agent pool description for users
        :param pulumi.Input[str] name: Name of the VCS agent pool, must be unique within an account
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VcsAgentPoolState.__new__(_VcsAgentPoolState)

        __props__.__dict__["config"] = config
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        return VcsAgentPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        VCS agent pool configuration, encoded using base64
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Free-form VCS agent pool description for users
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the VCS agent pool, must be unique within an account
        """
        return pulumi.get(self, "name")

