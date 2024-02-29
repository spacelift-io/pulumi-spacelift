# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['IdpGroupMappingArgs', 'IdpGroupMapping']

@pulumi.input_type
class IdpGroupMappingArgs:
    def __init__(__self__, *,
                 policies: pulumi.Input[Sequence[pulumi.Input['IdpGroupMappingPolicyArgs']]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IdpGroupMapping resource.
        :param pulumi.Input[str] name: Name of the user group - should be unique in one account
        """
        pulumi.set(__self__, "policies", policies)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Input[Sequence[pulumi.Input['IdpGroupMappingPolicyArgs']]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: pulumi.Input[Sequence[pulumi.Input['IdpGroupMappingPolicyArgs']]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user group - should be unique in one account
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IdpGroupMappingState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['IdpGroupMappingPolicyArgs']]]] = None):
        """
        Input properties used for looking up and filtering IdpGroupMapping resources.
        :param pulumi.Input[str] name: Name of the user group - should be unique in one account
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user group - should be unique in one account
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdpGroupMappingPolicyArgs']]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdpGroupMappingPolicyArgs']]]]):
        pulumi.set(self, "policies", value)


class IdpGroupMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdpGroupMappingPolicyArgs']]]]] = None,
                 __props__=None):
        """
        `IdpGroupMapping` represents a mapping (binding) between a user group (as provided by IdP) and a Spacelift User Management Policy. If you assign permissions (a Policy) to a user group, all users in the group will have those permissions unless the user's permissions are higher than the group's permissions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        test = spacelift.IdpGroupMapping("test", policies=[
            spacelift.IdpGroupMappingPolicyArgs(
                role="ADMIN",
                space_id="root",
            ),
            spacelift.IdpGroupMappingPolicyArgs(
                role="ADMIN",
                space_id="legacy",
            ),
        ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the user group - should be unique in one account
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdpGroupMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `IdpGroupMapping` represents a mapping (binding) between a user group (as provided by IdP) and a Spacelift User Management Policy. If you assign permissions (a Policy) to a user group, all users in the group will have those permissions unless the user's permissions are higher than the group's permissions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        test = spacelift.IdpGroupMapping("test", policies=[
            spacelift.IdpGroupMappingPolicyArgs(
                role="ADMIN",
                space_id="root",
            ),
            spacelift.IdpGroupMappingPolicyArgs(
                role="ADMIN",
                space_id="legacy",
            ),
        ])
        ```

        :param str resource_name: The name of the resource.
        :param IdpGroupMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdpGroupMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdpGroupMappingPolicyArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdpGroupMappingArgs.__new__(IdpGroupMappingArgs)

            __props__.__dict__["name"] = name
            if policies is None and not opts.urn:
                raise TypeError("Missing required property 'policies'")
            __props__.__dict__["policies"] = policies
        super(IdpGroupMapping, __self__).__init__(
            'spacelift:index/idpGroupMapping:IdpGroupMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdpGroupMappingPolicyArgs']]]]] = None) -> 'IdpGroupMapping':
        """
        Get an existing IdpGroupMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the user group - should be unique in one account
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdpGroupMappingState.__new__(_IdpGroupMappingState)

        __props__.__dict__["name"] = name
        __props__.__dict__["policies"] = policies
        return IdpGroupMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the user group - should be unique in one account
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence['outputs.IdpGroupMappingPolicy']]:
        return pulumi.get(self, "policies")

