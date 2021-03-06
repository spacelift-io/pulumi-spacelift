# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'ModuleGitlab',
    'StackCloudformation',
    'StackGitlab',
    'StackPulumi',
    'GetModuleGitlabResult',
    'GetStackCloudformationResult',
    'GetStackGitlabResult',
    'GetStackPulumiResult',
    'GetWorkerPoolsWorkerPoolResult',
]

@pulumi.output_type
class ModuleGitlab(dict):
    def __init__(__self__, *,
                 namespace: str):
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StackCloudformation(dict):
    def __init__(__self__, *,
                 entry_template_file: str,
                 region: str,
                 stack_name: str,
                 template_bucket: str):
        pulumi.set(__self__, "entry_template_file", entry_template_file)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "stack_name", stack_name)
        pulumi.set(__self__, "template_bucket", template_bucket)

    @property
    @pulumi.getter(name="entryTemplateFile")
    def entry_template_file(self) -> str:
        return pulumi.get(self, "entry_template_file")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> str:
        return pulumi.get(self, "stack_name")

    @property
    @pulumi.getter(name="templateBucket")
    def template_bucket(self) -> str:
        return pulumi.get(self, "template_bucket")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StackGitlab(dict):
    def __init__(__self__, *,
                 namespace: str):
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StackPulumi(dict):
    def __init__(__self__, *,
                 login_url: str,
                 stack_name: str):
        pulumi.set(__self__, "login_url", login_url)
        pulumi.set(__self__, "stack_name", stack_name)

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> str:
        return pulumi.get(self, "login_url")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> str:
        return pulumi.get(self, "stack_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetModuleGitlabResult(dict):
    def __init__(__self__, *,
                 namespace: str):
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")


@pulumi.output_type
class GetStackCloudformationResult(dict):
    def __init__(__self__, *,
                 entry_template_file: str,
                 region: str,
                 stack_name: str,
                 template_bucket: str):
        pulumi.set(__self__, "entry_template_file", entry_template_file)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "stack_name", stack_name)
        pulumi.set(__self__, "template_bucket", template_bucket)

    @property
    @pulumi.getter(name="entryTemplateFile")
    def entry_template_file(self) -> str:
        return pulumi.get(self, "entry_template_file")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> str:
        return pulumi.get(self, "stack_name")

    @property
    @pulumi.getter(name="templateBucket")
    def template_bucket(self) -> str:
        return pulumi.get(self, "template_bucket")


@pulumi.output_type
class GetStackGitlabResult(dict):
    def __init__(__self__, *,
                 namespace: str):
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")


@pulumi.output_type
class GetStackPulumiResult(dict):
    def __init__(__self__, *,
                 login_url: str,
                 stack_name: str):
        pulumi.set(__self__, "login_url", login_url)
        pulumi.set(__self__, "stack_name", stack_name)

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> str:
        return pulumi.get(self, "login_url")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> str:
        return pulumi.get(self, "stack_name")


@pulumi.output_type
class GetWorkerPoolsWorkerPoolResult(dict):
    def __init__(__self__, *,
                 config: str,
                 description: str,
                 name: str,
                 worker_pool_id: str):
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "worker_pool_id", worker_pool_id)

    @property
    @pulumi.getter
    def config(self) -> str:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="workerPoolId")
    def worker_pool_id(self) -> str:
        return pulumi.get(self, "worker_pool_id")


