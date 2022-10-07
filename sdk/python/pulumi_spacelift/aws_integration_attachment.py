# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AwsIntegrationAttachmentArgs', 'AwsIntegrationAttachment']

@pulumi.input_type
class AwsIntegrationAttachmentArgs:
    def __init__(__self__, *,
                 integration_id: pulumi.Input[str],
                 module_id: Optional[pulumi.Input[str]] = None,
                 read: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AwsIntegrationAttachment resource.
        :param pulumi.Input[str] integration_id: ID of the integration to attach
        :param pulumi.Input[str] module_id: ID of the module to attach the integration to
        :param pulumi.Input[bool] read: Indicates whether this attachment is used for read operations. Defaults to `true`.
        :param pulumi.Input[str] stack_id: ID of the stack to attach the integration to
        :param pulumi.Input[bool] write: Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        pulumi.set(__self__, "integration_id", integration_id)
        if module_id is not None:
            pulumi.set(__self__, "module_id", module_id)
        if read is not None:
            pulumi.set(__self__, "read", read)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)
        if write is not None:
            pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> pulumi.Input[str]:
        """
        ID of the integration to attach
        """
        return pulumi.get(self, "integration_id")

    @integration_id.setter
    def integration_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration_id", value)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module to attach the integration to
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter
    def read(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this attachment is used for read operations. Defaults to `true`.
        """
        return pulumi.get(self, "read")

    @read.setter
    def read(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack to attach the integration to
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter
    def write(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        return pulumi.get(self, "write")

    @write.setter
    def write(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write", value)


@pulumi.input_type
class _AwsIntegrationAttachmentState:
    def __init__(__self__, *,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 integration_id: Optional[pulumi.Input[str]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 read: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AwsIntegrationAttachment resources.
        :param pulumi.Input[str] attachment_id: Internal ID of the attachment entity
        :param pulumi.Input[str] integration_id: ID of the integration to attach
        :param pulumi.Input[str] module_id: ID of the module to attach the integration to
        :param pulumi.Input[bool] read: Indicates whether this attachment is used for read operations. Defaults to `true`.
        :param pulumi.Input[str] stack_id: ID of the stack to attach the integration to
        :param pulumi.Input[bool] write: Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        if attachment_id is not None:
            pulumi.set(__self__, "attachment_id", attachment_id)
        if integration_id is not None:
            pulumi.set(__self__, "integration_id", integration_id)
        if module_id is not None:
            pulumi.set(__self__, "module_id", module_id)
        if read is not None:
            pulumi.set(__self__, "read", read)
        if stack_id is not None:
            pulumi.set(__self__, "stack_id", stack_id)
        if write is not None:
            pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        Internal ID of the attachment entity
        """
        return pulumi.get(self, "attachment_id")

    @attachment_id.setter
    def attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attachment_id", value)

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the integration to attach
        """
        return pulumi.get(self, "integration_id")

    @integration_id.setter
    def integration_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_id", value)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module to attach the integration to
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter
    def read(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this attachment is used for read operations. Defaults to `true`.
        """
        return pulumi.get(self, "read")

    @read.setter
    def read(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack to attach the integration to
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter
    def write(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        return pulumi.get(self, "write")

    @write.setter
    def write(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write", value)


class AwsIntegrationAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 integration_id: Optional[pulumi.Input[str]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 read: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        **Note:** This resource is experimental. Please continue to use `AwsRole`.

        `AwsIntegrationAttachment` represents the attachment between a reusable AWS integration and a single stack or module.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        # For a stack
        this_aws_integration_attachment = spacelift.AwsIntegrationAttachment("thisAwsIntegrationAttachment",
            integration_id=spacelift_aws_integration["this"]["id"],
            stack_id="my-stack-id",
            read=True,
            write=True,
            opts=pulumi.ResourceOptions(depends_on=[aws_iam_role["this"]]))
        # For a module
        this_index_aws_integration_attachment_aws_integration_attachment = spacelift.AwsIntegrationAttachment("thisIndex/awsIntegrationAttachmentAwsIntegrationAttachment",
            integration_id=spacelift_aws_integration["this"]["id"],
            module_id="my-module-id",
            read=True,
            write=True,
            opts=pulumi.ResourceOptions(depends_on=[aws_iam_role["this"]]))
        ```

        ## Import

        ```sh
         $ pulumi import spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment read_write_my_stack $INTEGRATION_ID/$PROJECT_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_id: ID of the integration to attach
        :param pulumi.Input[str] module_id: ID of the module to attach the integration to
        :param pulumi.Input[bool] read: Indicates whether this attachment is used for read operations. Defaults to `true`.
        :param pulumi.Input[str] stack_id: ID of the stack to attach the integration to
        :param pulumi.Input[bool] write: Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsIntegrationAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Note:** This resource is experimental. Please continue to use `AwsRole`.

        `AwsIntegrationAttachment` represents the attachment between a reusable AWS integration and a single stack or module.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        # For a stack
        this_aws_integration_attachment = spacelift.AwsIntegrationAttachment("thisAwsIntegrationAttachment",
            integration_id=spacelift_aws_integration["this"]["id"],
            stack_id="my-stack-id",
            read=True,
            write=True,
            opts=pulumi.ResourceOptions(depends_on=[aws_iam_role["this"]]))
        # For a module
        this_index_aws_integration_attachment_aws_integration_attachment = spacelift.AwsIntegrationAttachment("thisIndex/awsIntegrationAttachmentAwsIntegrationAttachment",
            integration_id=spacelift_aws_integration["this"]["id"],
            module_id="my-module-id",
            read=True,
            write=True,
            opts=pulumi.ResourceOptions(depends_on=[aws_iam_role["this"]]))
        ```

        ## Import

        ```sh
         $ pulumi import spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment read_write_my_stack $INTEGRATION_ID/$PROJECT_ID
        ```

        :param str resource_name: The name of the resource.
        :param AwsIntegrationAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsIntegrationAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 integration_id: Optional[pulumi.Input[str]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 read: Optional[pulumi.Input[bool]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsIntegrationAttachmentArgs.__new__(AwsIntegrationAttachmentArgs)

            if integration_id is None and not opts.urn:
                raise TypeError("Missing required property 'integration_id'")
            __props__.__dict__["integration_id"] = integration_id
            __props__.__dict__["module_id"] = module_id
            __props__.__dict__["read"] = read
            __props__.__dict__["stack_id"] = stack_id
            __props__.__dict__["write"] = write
            __props__.__dict__["attachment_id"] = None
        super(AwsIntegrationAttachment, __self__).__init__(
            'spacelift:index/awsIntegrationAttachment:AwsIntegrationAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachment_id: Optional[pulumi.Input[str]] = None,
            integration_id: Optional[pulumi.Input[str]] = None,
            module_id: Optional[pulumi.Input[str]] = None,
            read: Optional[pulumi.Input[bool]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            write: Optional[pulumi.Input[bool]] = None) -> 'AwsIntegrationAttachment':
        """
        Get an existing AwsIntegrationAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_id: Internal ID of the attachment entity
        :param pulumi.Input[str] integration_id: ID of the integration to attach
        :param pulumi.Input[str] module_id: ID of the module to attach the integration to
        :param pulumi.Input[bool] read: Indicates whether this attachment is used for read operations. Defaults to `true`.
        :param pulumi.Input[str] stack_id: ID of the stack to attach the integration to
        :param pulumi.Input[bool] write: Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsIntegrationAttachmentState.__new__(_AwsIntegrationAttachmentState)

        __props__.__dict__["attachment_id"] = attachment_id
        __props__.__dict__["integration_id"] = integration_id
        __props__.__dict__["module_id"] = module_id
        __props__.__dict__["read"] = read
        __props__.__dict__["stack_id"] = stack_id
        __props__.__dict__["write"] = write
        return AwsIntegrationAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> pulumi.Output[str]:
        """
        Internal ID of the attachment entity
        """
        return pulumi.get(self, "attachment_id")

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> pulumi.Output[str]:
        """
        ID of the integration to attach
        """
        return pulumi.get(self, "integration_id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the module to attach the integration to
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def read(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this attachment is used for read operations. Defaults to `true`.
        """
        return pulumi.get(self, "read")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the stack to attach the integration to
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def write(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this attachment is used for write operations. Defaults to `true`.
        """
        return pulumi.get(self, "write")

