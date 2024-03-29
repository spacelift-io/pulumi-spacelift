# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AwsIntegrationArgs', 'AwsIntegration']

@pulumi.input_type
class AwsIntegrationArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AwsIntegration resource.
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[int] duration_seconds: Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: Labels to set on the integration
        :param pulumi.Input[str] name: The friendly name of the integration
        :param pulumi.Input[str] space_id: ID (slug) of the space the integration is in
        """
        AwsIntegrationArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            role_arn=role_arn,
            duration_seconds=duration_seconds,
            external_id=external_id,
            generate_credentials_in_worker=generate_credentials_in_worker,
            labels=labels,
            name=name,
            space_id=space_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             role_arn: Optional[pulumi.Input[str]] = None,
             duration_seconds: Optional[pulumi.Input[int]] = None,
             external_id: Optional[pulumi.Input[str]] = None,
             generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
             labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             space_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if role_arn is None and 'roleArn' in kwargs:
            role_arn = kwargs['roleArn']
        if role_arn is None:
            raise TypeError("Missing 'role_arn' argument")
        if duration_seconds is None and 'durationSeconds' in kwargs:
            duration_seconds = kwargs['durationSeconds']
        if external_id is None and 'externalId' in kwargs:
            external_id = kwargs['externalId']
        if generate_credentials_in_worker is None and 'generateCredentialsInWorker' in kwargs:
            generate_credentials_in_worker = kwargs['generateCredentialsInWorker']
        if space_id is None and 'spaceId' in kwargs:
            space_id = kwargs['spaceId']

        _setter("role_arn", role_arn)
        if duration_seconds is not None:
            _setter("duration_seconds", duration_seconds)
        if external_id is not None:
            _setter("external_id", external_id)
        if generate_credentials_in_worker is not None:
            _setter("generate_credentials_in_worker", generate_credentials_in_worker)
        if labels is not None:
            _setter("labels", labels)
        if name is not None:
            _setter("name", name)
        if space_id is not None:
            _setter("space_id", space_id)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        ARN of the AWS IAM role to attach
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        """
        return pulumi.get(self, "duration_seconds")

    @duration_seconds.setter
    def duration_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_seconds", value)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[pulumi.Input[str]]:
        """
        Custom external ID (works only for private workers).
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter(name="generateCredentialsInWorker")
    def generate_credentials_in_worker(self) -> Optional[pulumi.Input[bool]]:
        """
        Generate AWS credentials in the private worker. Defaults to `false`.
        """
        return pulumi.get(self, "generate_credentials_in_worker")

    @generate_credentials_in_worker.setter
    def generate_credentials_in_worker(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "generate_credentials_in_worker", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Labels to set on the integration
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The friendly name of the integration
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID (slug) of the space the integration is in
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


@pulumi.input_type
class _AwsIntegrationState:
    def __init__(__self__, *,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AwsIntegration resources.
        :param pulumi.Input[int] duration_seconds: Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: Labels to set on the integration
        :param pulumi.Input[str] name: The friendly name of the integration
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[str] space_id: ID (slug) of the space the integration is in
        """
        _AwsIntegrationState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            duration_seconds=duration_seconds,
            external_id=external_id,
            generate_credentials_in_worker=generate_credentials_in_worker,
            labels=labels,
            name=name,
            role_arn=role_arn,
            space_id=space_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             duration_seconds: Optional[pulumi.Input[int]] = None,
             external_id: Optional[pulumi.Input[str]] = None,
             generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
             labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             role_arn: Optional[pulumi.Input[str]] = None,
             space_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if duration_seconds is None and 'durationSeconds' in kwargs:
            duration_seconds = kwargs['durationSeconds']
        if external_id is None and 'externalId' in kwargs:
            external_id = kwargs['externalId']
        if generate_credentials_in_worker is None and 'generateCredentialsInWorker' in kwargs:
            generate_credentials_in_worker = kwargs['generateCredentialsInWorker']
        if role_arn is None and 'roleArn' in kwargs:
            role_arn = kwargs['roleArn']
        if space_id is None and 'spaceId' in kwargs:
            space_id = kwargs['spaceId']

        if duration_seconds is not None:
            _setter("duration_seconds", duration_seconds)
        if external_id is not None:
            _setter("external_id", external_id)
        if generate_credentials_in_worker is not None:
            _setter("generate_credentials_in_worker", generate_credentials_in_worker)
        if labels is not None:
            _setter("labels", labels)
        if name is not None:
            _setter("name", name)
        if role_arn is not None:
            _setter("role_arn", role_arn)
        if space_id is not None:
            _setter("space_id", space_id)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        """
        return pulumi.get(self, "duration_seconds")

    @duration_seconds.setter
    def duration_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_seconds", value)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[pulumi.Input[str]]:
        """
        Custom external ID (works only for private workers).
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter(name="generateCredentialsInWorker")
    def generate_credentials_in_worker(self) -> Optional[pulumi.Input[bool]]:
        """
        Generate AWS credentials in the private worker. Defaults to `false`.
        """
        return pulumi.get(self, "generate_credentials_in_worker")

    @generate_credentials_in_worker.setter
    def generate_credentials_in_worker(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "generate_credentials_in_worker", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Labels to set on the integration
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The friendly name of the integration
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the AWS IAM role to attach
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID (slug) of the space the integration is in
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


class AwsIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `AwsIntegration` represents an integration with an AWS account. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect.

        Note: when assuming credentials for **shared workers**, Spacelift will use `$accountName@$integrationID@$stackID@$suffix` or `$accountName@$integrationID@$moduleID@$suffix` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole),$suffix will be `read` or `write`.

        ## Import

        ```sh
         $ pulumi import spacelift:index/awsIntegration:AwsIntegration read_write_integration $INTEGRATION_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration_seconds: Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: Labels to set on the integration
        :param pulumi.Input[str] name: The friendly name of the integration
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[str] space_id: ID (slug) of the space the integration is in
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsIntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AwsIntegration` represents an integration with an AWS account. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect.

        Note: when assuming credentials for **shared workers**, Spacelift will use `$accountName@$integrationID@$stackID@$suffix` or `$accountName@$integrationID@$moduleID@$suffix` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole),$suffix will be `read` or `write`.

        ## Import

        ```sh
         $ pulumi import spacelift:index/awsIntegration:AwsIntegration read_write_integration $INTEGRATION_ID
        ```

        :param str resource_name: The name of the resource.
        :param AwsIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            AwsIntegrationArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsIntegrationArgs.__new__(AwsIntegrationArgs)

            __props__.__dict__["duration_seconds"] = duration_seconds
            __props__.__dict__["external_id"] = external_id
            __props__.__dict__["generate_credentials_in_worker"] = generate_credentials_in_worker
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["space_id"] = space_id
        super(AwsIntegration, __self__).__init__(
            'spacelift:index/awsIntegration:AwsIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            duration_seconds: Optional[pulumi.Input[int]] = None,
            external_id: Optional[pulumi.Input[str]] = None,
            generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None) -> 'AwsIntegration':
        """
        Get an existing AwsIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration_seconds: Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: Labels to set on the integration
        :param pulumi.Input[str] name: The friendly name of the integration
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[str] space_id: ID (slug) of the space the integration is in
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsIntegrationState.__new__(_AwsIntegrationState)

        __props__.__dict__["duration_seconds"] = duration_seconds
        __props__.__dict__["external_id"] = external_id
        __props__.__dict__["generate_credentials_in_worker"] = generate_credentials_in_worker
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["space_id"] = space_id
        return AwsIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        Duration in seconds for which the assumed role credentials should be valid. Defaults to `900`.
        """
        return pulumi.get(self, "duration_seconds")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[Optional[str]]:
        """
        Custom external ID (works only for private workers).
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="generateCredentialsInWorker")
    def generate_credentials_in_worker(self) -> pulumi.Output[Optional[bool]]:
        """
        Generate AWS credentials in the private worker. Defaults to `false`.
        """
        return pulumi.get(self, "generate_credentials_in_worker")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Labels to set on the integration
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The friendly name of the integration
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        ARN of the AWS IAM role to attach
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        ID (slug) of the space the integration is in
        """
        return pulumi.get(self, "space_id")

