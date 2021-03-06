# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .aws_role import *
from .context import *
from .context_attachment import *
from .environment_variable import *
from .gcp_service_account import *
from .get_aws_role import *
from .get_context import *
from .get_context_attachment import *
from .get_current_stack import *
from .get_environment_variable import *
from .get_gcp_service_account import *
from .get_ips import *
from .get_module import *
from .get_mounted_file import *
from .get_policy import *
from .get_stack import *
from .get_stack_aws_role import *
from .get_stack_gcp_service_account import *
from .get_webhook import *
from .get_worker_pool import *
from .get_worker_pools import *
from .module import *
from .mounted_file import *
from .policy import *
from .policy_attachment import *
from .provider import *
from .stack import *
from .stack_aws_role import *
from .stack_gcp_service_account import *
from .webhook import *
from .worker_pool import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "spacelift:index/awsRole:AwsRole":
                return AwsRole(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/context:Context":
                return Context(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/contextAttachment:ContextAttachment":
                return ContextAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/environmentVariable:EnvironmentVariable":
                return EnvironmentVariable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/gcpServiceAccount:GcpServiceAccount":
                return GcpServiceAccount(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/module:Module":
                return Module(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/mountedFile:MountedFile":
                return MountedFile(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/policy:Policy":
                return Policy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/policyAttachment:PolicyAttachment":
                return PolicyAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/stack:Stack":
                return Stack(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/stackAwsRole:StackAwsRole":
                return StackAwsRole(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount":
                return StackGcpServiceAccount(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/webhook:Webhook":
                return Webhook(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "spacelift:index/workerPool:WorkerPool":
                return WorkerPool(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("spacelift", "index/awsRole", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/context", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/contextAttachment", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/environmentVariable", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/gcpServiceAccount", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/module", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/mountedFile", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/policy", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/policyAttachment", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/stack", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/stackAwsRole", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/stackGcpServiceAccount", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/webhook", _module_instance)
    pulumi.runtime.register_resource_module("spacelift", "index/workerPool", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:spacelift":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("spacelift", Package())

_register_module()
