# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StackGcpServiceAccountArgs', 'StackGcpServiceAccount']

@pulumi.input_type
class StackGcpServiceAccountArgs:
    def __init__(__self__, *,
                 token_scopes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 module_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StackGcpServiceAccount resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of scopes that will be requested when generating temporary GCP service account credentials
        :param pulumi.Input[str] module_id: ID of the module which uses GCP service account credentials
        :param pulumi.Input[str] stack_id: ID of the stack which uses GCP service account credentials
        """
        StackGcpServiceAccountArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            token_scopes=token_scopes,
            module_id=module_id,
            stack_id=stack_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             module_id: Optional[pulumi.Input[str]] = None,
             stack_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if token_scopes is None and 'tokenScopes' in kwargs:
            token_scopes = kwargs['tokenScopes']
        if token_scopes is None:
            raise TypeError("Missing 'token_scopes' argument")
        if module_id is None and 'moduleId' in kwargs:
            module_id = kwargs['moduleId']
        if stack_id is None and 'stackId' in kwargs:
            stack_id = kwargs['stackId']

        _setter("token_scopes", token_scopes)
        if module_id is not None:
            _setter("module_id", module_id)
        if stack_id is not None:
            _setter("stack_id", stack_id)

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of scopes that will be requested when generating temporary GCP service account credentials
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "token_scopes", value)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module which uses GCP service account credentials
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack which uses GCP service account credentials
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)


@pulumi.input_type
class _StackGcpServiceAccountState:
    def __init__(__self__, *,
                 module_id: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering StackGcpServiceAccount resources.
        :param pulumi.Input[str] module_id: ID of the module which uses GCP service account credentials
        :param pulumi.Input[str] service_account_email: Email address of the GCP service account dedicated for this stack
        :param pulumi.Input[str] stack_id: ID of the stack which uses GCP service account credentials
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of scopes that will be requested when generating temporary GCP service account credentials
        """
        _StackGcpServiceAccountState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            module_id=module_id,
            service_account_email=service_account_email,
            stack_id=stack_id,
            token_scopes=token_scopes,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             module_id: Optional[pulumi.Input[str]] = None,
             service_account_email: Optional[pulumi.Input[str]] = None,
             stack_id: Optional[pulumi.Input[str]] = None,
             token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if module_id is None and 'moduleId' in kwargs:
            module_id = kwargs['moduleId']
        if service_account_email is None and 'serviceAccountEmail' in kwargs:
            service_account_email = kwargs['serviceAccountEmail']
        if stack_id is None and 'stackId' in kwargs:
            stack_id = kwargs['stackId']
        if token_scopes is None and 'tokenScopes' in kwargs:
            token_scopes = kwargs['tokenScopes']

        if module_id is not None:
            _setter("module_id", module_id)
        if service_account_email is not None:
            _setter("service_account_email", service_account_email)
        if stack_id is not None:
            _setter("stack_id", stack_id)
        if token_scopes is not None:
            _setter("token_scopes", token_scopes)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module which uses GCP service account credentials
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        Email address of the GCP service account dedicated for this stack
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the stack which uses GCP service account credentials
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of scopes that will be requested when generating temporary GCP service account credentials
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "token_scopes", value)


class StackGcpServiceAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_google as google
        import pulumi_spacelift as spacelift

        k8s_core_stack = spacelift.Stack("k8s-coreStack",
            branch="master",
            repository="core-infra")
        k8s_core_stack_gcp_service_account = spacelift.StackGcpServiceAccount("k8s-coreStackGcpServiceAccount",
            stack_id=k8s_core_stack.id,
            token_scopes=[
                "https://www.googleapis.com/auth/compute",
                "https://www.googleapis.com/auth/cloud-platform",
                "https://www.googleapis.com/auth/devstorage.full_control",
            ])
        k8s_coregoogle_project = google.index.Google_project("k8s-coregoogle_project",
            name=Kubernetes code,
            project_id=unicorn-k8s-core,
            org_id=var.gcp_organization_id)
        k8s_coregoogle_project_iam_member = google.index.Google_project_iam_member("k8s-coregoogle_project_iam_member",
            project=k8s_coregoogle_project.id,
            role=roles/owner,
            member=fserviceAccount:{k8s_core_stack_gcp_service_account.service_account_email})
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] module_id: ID of the module which uses GCP service account credentials
        :param pulumi.Input[str] stack_id: ID of the stack which uses GCP service account credentials
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of scopes that will be requested when generating temporary GCP service account credentials
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StackGcpServiceAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_google as google
        import pulumi_spacelift as spacelift

        k8s_core_stack = spacelift.Stack("k8s-coreStack",
            branch="master",
            repository="core-infra")
        k8s_core_stack_gcp_service_account = spacelift.StackGcpServiceAccount("k8s-coreStackGcpServiceAccount",
            stack_id=k8s_core_stack.id,
            token_scopes=[
                "https://www.googleapis.com/auth/compute",
                "https://www.googleapis.com/auth/cloud-platform",
                "https://www.googleapis.com/auth/devstorage.full_control",
            ])
        k8s_coregoogle_project = google.index.Google_project("k8s-coregoogle_project",
            name=Kubernetes code,
            project_id=unicorn-k8s-core,
            org_id=var.gcp_organization_id)
        k8s_coregoogle_project_iam_member = google.index.Google_project_iam_member("k8s-coregoogle_project_iam_member",
            project=k8s_coregoogle_project.id,
            role=roles/owner,
            member=fserviceAccount:{k8s_core_stack_gcp_service_account.service_account_email})
        ```

        :param str resource_name: The name of the resource.
        :param StackGcpServiceAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StackGcpServiceAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            StackGcpServiceAccountArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StackGcpServiceAccountArgs.__new__(StackGcpServiceAccountArgs)

            __props__.__dict__["module_id"] = module_id
            __props__.__dict__["stack_id"] = stack_id
            if token_scopes is None and not opts.urn:
                raise TypeError("Missing required property 'token_scopes'")
            __props__.__dict__["token_scopes"] = token_scopes
            __props__.__dict__["service_account_email"] = None
        super(StackGcpServiceAccount, __self__).__init__(
            'spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            module_id: Optional[pulumi.Input[str]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'StackGcpServiceAccount':
        """
        Get an existing StackGcpServiceAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] module_id: ID of the module which uses GCP service account credentials
        :param pulumi.Input[str] service_account_email: Email address of the GCP service account dedicated for this stack
        :param pulumi.Input[str] stack_id: ID of the stack which uses GCP service account credentials
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of scopes that will be requested when generating temporary GCP service account credentials
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StackGcpServiceAccountState.__new__(_StackGcpServiceAccountState)

        __props__.__dict__["module_id"] = module_id
        __props__.__dict__["service_account_email"] = service_account_email
        __props__.__dict__["stack_id"] = stack_id
        __props__.__dict__["token_scopes"] = token_scopes
        return StackGcpServiceAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the module which uses GCP service account credentials
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        Email address of the GCP service account dedicated for this stack
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the stack which uses GCP service account credentials
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> pulumi.Output[Sequence[str]]:
        """
        List of scopes that will be requested when generating temporary GCP service account credentials
        """
        return pulumi.get(self, "token_scopes")

