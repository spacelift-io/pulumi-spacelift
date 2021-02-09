# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Module']


class Module(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrative: Optional[pulumi.Input[bool]] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gitlab: Optional[pulumi.Input[pulumi.InputType['ModuleGitlabArgs']]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 shared_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 worker_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Module resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] administrative: Indicates whether this module can manage others
        :param pulumi.Input[str] branch: GitHub branch to apply changes to
        :param pulumi.Input[str] description: Free-form module description for users
        :param pulumi.Input[str] repository: Name of the repository, without the owner part
        :param pulumi.Input[Sequence[pulumi.Input[str]]] shared_accounts: List of the accounts (subdomains) which should have access to the Module
        :param pulumi.Input[str] worker_pool_id: ID of the worker pool to use
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

            __props__['administrative'] = administrative
            if branch is None and not opts.urn:
                raise TypeError("Missing required property 'branch'")
            __props__['branch'] = branch
            __props__['description'] = description
            __props__['gitlab'] = gitlab
            __props__['labels'] = labels
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__['repository'] = repository
            __props__['shared_accounts'] = shared_accounts
            __props__['worker_pool_id'] = worker_pool_id
            __props__['aws_assume_role_policy_statement'] = None
        super(Module, __self__).__init__(
            'spacelift:index/module:Module',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            administrative: Optional[pulumi.Input[bool]] = None,
            aws_assume_role_policy_statement: Optional[pulumi.Input[str]] = None,
            branch: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            gitlab: Optional[pulumi.Input[pulumi.InputType['ModuleGitlabArgs']]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            shared_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            worker_pool_id: Optional[pulumi.Input[str]] = None) -> 'Module':
        """
        Get an existing Module resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] administrative: Indicates whether this module can manage others
        :param pulumi.Input[str] aws_assume_role_policy_statement: AWS IAM assume role policy statement setting up trust relationship
        :param pulumi.Input[str] branch: GitHub branch to apply changes to
        :param pulumi.Input[str] description: Free-form module description for users
        :param pulumi.Input[str] repository: Name of the repository, without the owner part
        :param pulumi.Input[Sequence[pulumi.Input[str]]] shared_accounts: List of the accounts (subdomains) which should have access to the Module
        :param pulumi.Input[str] worker_pool_id: ID of the worker pool to use
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["administrative"] = administrative
        __props__["aws_assume_role_policy_statement"] = aws_assume_role_policy_statement
        __props__["branch"] = branch
        __props__["description"] = description
        __props__["gitlab"] = gitlab
        __props__["labels"] = labels
        __props__["repository"] = repository
        __props__["shared_accounts"] = shared_accounts
        __props__["worker_pool_id"] = worker_pool_id
        return Module(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def administrative(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this module can manage others
        """
        return pulumi.get(self, "administrative")

    @property
    @pulumi.getter(name="awsAssumeRolePolicyStatement")
    def aws_assume_role_policy_statement(self) -> pulumi.Output[str]:
        """
        AWS IAM assume role policy statement setting up trust relationship
        """
        return pulumi.get(self, "aws_assume_role_policy_statement")

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[str]:
        """
        GitHub branch to apply changes to
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Free-form module description for users
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def gitlab(self) -> pulumi.Output[Optional['outputs.ModuleGitlab']]:
        return pulumi.get(self, "gitlab")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        Name of the repository, without the owner part
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="sharedAccounts")
    def shared_accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of the accounts (subdomains) which should have access to the Module
        """
        return pulumi.get(self, "shared_accounts")

    @property
    @pulumi.getter(name="workerPoolId")
    def worker_pool_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the worker pool to use
        """
        return pulumi.get(self, "worker_pool_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
