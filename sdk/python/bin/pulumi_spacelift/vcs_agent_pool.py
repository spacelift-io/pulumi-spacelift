# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['VCSAgentPool']


class VCSAgentPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `VCSAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        ghe = spacelift.VCSAgentPool("ghe",
            description="VCS agent pool for our internal GitHub Enterpris",
            name="ghe")
        ```

        <!-- schema generated by tfplugindocs -->
        ## Schema

        ### Required

        - **name** (String) Name of the VCS agent pool, must be unique within an account

        ### Optional

        - **description** (String) Free-form VCS agent pool description for users
        - **id** (String) The ID of this resource.

        ### Read-Only

        - **config** (String, Sensitive) VCS agent pool configuration, encoded using base64

        ## Import

        Import is supported using the following syntax

        ```sh
         $ pulumi import spacelift:index/vCSAgentPool:VCSAgentPool ghe $VCS_AGENT_POOL_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Free-form VCS agent pool description for users
        :param pulumi.Input[str] name: Name of the VCS agent pool, must be unique within an account
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

            __props__['description'] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['config'] = None
        super(VCSAgentPool, __self__).__init__(
            'spacelift:index/vCSAgentPool:VCSAgentPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'VCSAgentPool':
        """
        Get an existing VCSAgentPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: VCS agent pool configuration, encoded using base64
        :param pulumi.Input[str] description: Free-form VCS agent pool description for users
        :param pulumi.Input[str] name: Name of the VCS agent pool, must be unique within an account
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config"] = config
        __props__["description"] = description
        __props__["name"] = name
        return VCSAgentPool(resource_name, opts=opts, __props__=__props__)

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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
