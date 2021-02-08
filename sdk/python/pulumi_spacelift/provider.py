# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key_endpoint: Optional[pulumi.Input[str]] = None,
                 api_key_id: Optional[pulumi.Input[str]] = None,
                 api_key_secret: Optional[pulumi.Input[str]] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the spacelift package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key_endpoint: Endpoint to use when authenticating with an API key outside of Spacelift
        :param pulumi.Input[str] api_key_id: ID of the API key to use when executing outside of Spacelift
        :param pulumi.Input[str] api_key_secret: API key secret to use when executing outside of Spacelift
        :param pulumi.Input[str] api_token: Spacelift token generated by a run, only useful from within Spacelift
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

            if api_key_endpoint is None:
                api_key_endpoint = _utilities.get_env('SPACELIFT_API_KEY_ENDPOINT')
            __props__['api_key_endpoint'] = api_key_endpoint
            if api_key_id is None:
                api_key_id = _utilities.get_env('SPACELIFT_API_KEY_ID')
            __props__['api_key_id'] = api_key_id
            if api_key_secret is None:
                api_key_secret = _utilities.get_env('SPACELIFT_API_KEY_SECRET')
            __props__['api_key_secret'] = api_key_secret
            if api_token is None:
                api_token = _utilities.get_env('SPACELIFT_API_TOKEN')
            __props__['api_token'] = api_token
        super(Provider, __self__).__init__(
            'spacelift',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

