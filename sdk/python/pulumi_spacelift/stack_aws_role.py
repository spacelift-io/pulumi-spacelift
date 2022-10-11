# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StackAwsRoleArgs', 'StackAwsRole']

@pulumi.input_type
class StackAwsRoleArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StackAwsRole resource.
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[int] duration_seconds: AWS IAM role session duration in seconds
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[str] module_id: ID of the module which assumes the AWS IAM role
        :param pulumi.Input[str] stack_id: ID of the stack which assumes the AWS IAM role
        """
        pulumi.set(__self__, "role_arn", role_arn)
        if duration_seconds is not None:
            pulumi.set(__self__, "duration_seconds", duration_seconds)
        if external_id is not None:
            pulumi.set(__self__, "external_id", external_id)
        if generate_credentials_in_worker is not None:
            pulumi.set(__self__, "generate_credentials_in_worker", generate_credentials_in_worker)
        if module_id is not None:
            pulumi.set(__self__, "module_id", module_id)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)

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
        AWS IAM role session duration in seconds
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
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module which assumes the AWS IAM role
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack which assumes the AWS IAM role
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)


@pulumi.input_type
class _StackAwsRoleState:
    def __init__(__self__, *,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StackAwsRole resources.
        :param pulumi.Input[int] duration_seconds: AWS IAM role session duration in seconds
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[str] module_id: ID of the module which assumes the AWS IAM role
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[str] stack_id: ID of the stack which assumes the AWS IAM role
        """
        if duration_seconds is not None:
            pulumi.set(__self__, "duration_seconds", duration_seconds)
        if external_id is not None:
            pulumi.set(__self__, "external_id", external_id)
        if generate_credentials_in_worker is not None:
            pulumi.set(__self__, "generate_credentials_in_worker", generate_credentials_in_worker)
        if module_id is not None:
            pulumi.set(__self__, "module_id", module_id)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        AWS IAM role session duration in seconds
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
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module which assumes the AWS IAM role
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

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
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack which assumes the AWS IAM role
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)


class StackAwsRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        > **Note:** `StackAwsRole` is deprecated. Please use `AwsRole` instead. The functionality is identical.

        **NOTE:** while this resource continues to work, we have replaced it with the `AwsIntegration` resource. The new resource allows integrations to be shared by multiple stacks/modules and also supports separate read vs write roles. Please use the `AwsIntegration` resource instead.

        `StackAwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.

        If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).

        Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration_seconds: AWS IAM role session duration in seconds
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[str] module_id: ID of the module which assumes the AWS IAM role
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[str] stack_id: ID of the stack which assumes the AWS IAM role
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StackAwsRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > **Note:** `StackAwsRole` is deprecated. Please use `AwsRole` instead. The functionality is identical.

        **NOTE:** while this resource continues to work, we have replaced it with the `AwsIntegration` resource. The new resource allows integrations to be shared by multiple stacks/modules and also supports separate read vs write roles. Please use the `AwsIntegration` resource instead.

        `StackAwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.

        If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).

        Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).

        :param str resource_name: The name of the resource.
        :param StackAwsRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StackAwsRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 generate_credentials_in_worker: Optional[pulumi.Input[bool]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StackAwsRoleArgs.__new__(StackAwsRoleArgs)

            __props__.__dict__["duration_seconds"] = duration_seconds
            __props__.__dict__["external_id"] = external_id
            __props__.__dict__["generate_credentials_in_worker"] = generate_credentials_in_worker
            __props__.__dict__["module_id"] = module_id
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["stack_id"] = stack_id
        super(StackAwsRole, __self__).__init__(
            'spacelift:index/stackAwsRole:StackAwsRole',
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
            module_id: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None) -> 'StackAwsRole':
        """
        Get an existing StackAwsRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration_seconds: AWS IAM role session duration in seconds
        :param pulumi.Input[str] external_id: Custom external ID (works only for private workers).
        :param pulumi.Input[bool] generate_credentials_in_worker: Generate AWS credentials in the private worker. Defaults to `false`.
        :param pulumi.Input[str] module_id: ID of the module which assumes the AWS IAM role
        :param pulumi.Input[str] role_arn: ARN of the AWS IAM role to attach
        :param pulumi.Input[str] stack_id: ID of the stack which assumes the AWS IAM role
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StackAwsRoleState.__new__(_StackAwsRoleState)

        __props__.__dict__["duration_seconds"] = duration_seconds
        __props__.__dict__["external_id"] = external_id
        __props__.__dict__["generate_credentials_in_worker"] = generate_credentials_in_worker
        __props__.__dict__["module_id"] = module_id
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["stack_id"] = stack_id
        return StackAwsRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> pulumi.Output[int]:
        """
        AWS IAM role session duration in seconds
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
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the module which assumes the AWS IAM role
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        ARN of the AWS IAM role to attach
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the stack which assumes the AWS IAM role
        """
        return pulumi.get(self, "stack_id")

