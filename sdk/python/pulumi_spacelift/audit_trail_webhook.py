# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AuditTrailWebhookArgs', 'AuditTrailWebhook']

@pulumi.input_type
class AuditTrailWebhookArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 endpoint: pulumi.Input[str],
                 secret: pulumi.Input[str],
                 include_runs: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AuditTrailWebhook resource.
        :param pulumi.Input[bool] enabled: `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        :param pulumi.Input[str] endpoint: `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        :param pulumi.Input[str] secret: `secret` is a secret that Spacelift will send with the request
        :param pulumi.Input[bool] include_runs: `include_runs` determines whether the webhook should include information about the run that triggered the event.
        """
        AuditTrailWebhookArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            enabled=enabled,
            endpoint=endpoint,
            secret=secret,
            include_runs=include_runs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             enabled: Optional[pulumi.Input[bool]] = None,
             endpoint: Optional[pulumi.Input[str]] = None,
             secret: Optional[pulumi.Input[str]] = None,
             include_runs: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if enabled is None:
            raise TypeError("Missing 'enabled' argument")
        if endpoint is None:
            raise TypeError("Missing 'endpoint' argument")
        if secret is None:
            raise TypeError("Missing 'secret' argument")
        if include_runs is None and 'includeRuns' in kwargs:
            include_runs = kwargs['includeRuns']

        _setter("enabled", enabled)
        _setter("endpoint", endpoint)
        _setter("secret", secret)
        if include_runs is not None:
            _setter("include_runs", include_runs)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Input[str]:
        """
        `secret` is a secret that Spacelift will send with the request
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter(name="includeRuns")
    def include_runs(self) -> Optional[pulumi.Input[bool]]:
        """
        `include_runs` determines whether the webhook should include information about the run that triggered the event.
        """
        return pulumi.get(self, "include_runs")

    @include_runs.setter
    def include_runs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_runs", value)


@pulumi.input_type
class _AuditTrailWebhookState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 include_runs: Optional[pulumi.Input[bool]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuditTrailWebhook resources.
        :param pulumi.Input[bool] enabled: `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        :param pulumi.Input[str] endpoint: `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        :param pulumi.Input[bool] include_runs: `include_runs` determines whether the webhook should include information about the run that triggered the event.
        :param pulumi.Input[str] secret: `secret` is a secret that Spacelift will send with the request
        """
        _AuditTrailWebhookState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            enabled=enabled,
            endpoint=endpoint,
            include_runs=include_runs,
            secret=secret,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             enabled: Optional[pulumi.Input[bool]] = None,
             endpoint: Optional[pulumi.Input[str]] = None,
             include_runs: Optional[pulumi.Input[bool]] = None,
             secret: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if include_runs is None and 'includeRuns' in kwargs:
            include_runs = kwargs['includeRuns']

        if enabled is not None:
            _setter("enabled", enabled)
        if endpoint is not None:
            _setter("endpoint", endpoint)
        if include_runs is not None:
            _setter("include_runs", include_runs)
        if secret is not None:
            _setter("secret", secret)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="includeRuns")
    def include_runs(self) -> Optional[pulumi.Input[bool]]:
        """
        `include_runs` determines whether the webhook should include information about the run that triggered the event.
        """
        return pulumi.get(self, "include_runs")

    @include_runs.setter
    def include_runs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_runs", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        `secret` is a secret that Spacelift will send with the request
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


class AuditTrailWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 include_runs: Optional[pulumi.Input[bool]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `AuditTrailWebhook` represents a webhook endpoint to which Spacelift sends POST requests about audit events.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        example = spacelift.AuditTrailWebhook("example",
            enabled=True,
            endpoint="https://example.com",
            secret="mysecretkey")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        :param pulumi.Input[str] endpoint: `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        :param pulumi.Input[bool] include_runs: `include_runs` determines whether the webhook should include information about the run that triggered the event.
        :param pulumi.Input[str] secret: `secret` is a secret that Spacelift will send with the request
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuditTrailWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AuditTrailWebhook` represents a webhook endpoint to which Spacelift sends POST requests about audit events.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        example = spacelift.AuditTrailWebhook("example",
            enabled=True,
            endpoint="https://example.com",
            secret="mysecretkey")
        ```

        :param str resource_name: The name of the resource.
        :param AuditTrailWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuditTrailWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            AuditTrailWebhookArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 include_runs: Optional[pulumi.Input[bool]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuditTrailWebhookArgs.__new__(AuditTrailWebhookArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["include_runs"] = include_runs
            if secret is None and not opts.urn:
                raise TypeError("Missing required property 'secret'")
            __props__.__dict__["secret"] = None if secret is None else pulumi.Output.secret(secret)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AuditTrailWebhook, __self__).__init__(
            'spacelift:index/auditTrailWebhook:AuditTrailWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            include_runs: Optional[pulumi.Input[bool]] = None,
            secret: Optional[pulumi.Input[str]] = None) -> 'AuditTrailWebhook':
        """
        Get an existing AuditTrailWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        :param pulumi.Input[str] endpoint: `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        :param pulumi.Input[bool] include_runs: `include_runs` determines whether the webhook should include information about the run that triggered the event.
        :param pulumi.Input[str] secret: `secret` is a secret that Spacelift will send with the request
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuditTrailWebhookState.__new__(_AuditTrailWebhookState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["include_runs"] = include_runs
        __props__.__dict__["secret"] = secret
        return AuditTrailWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        `endpoint` is the URL to which Spacelift will send POST requests about audit events.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="includeRuns")
    def include_runs(self) -> pulumi.Output[Optional[bool]]:
        """
        `include_runs` determines whether the webhook should include information about the run that triggered the event.
        """
        return pulumi.get(self, "include_runs")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        `secret` is a secret that Spacelift will send with the request
        """
        return pulumi.get(self, "secret")

