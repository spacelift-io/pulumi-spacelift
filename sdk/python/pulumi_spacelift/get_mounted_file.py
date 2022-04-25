# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetMountedFileResult',
    'AwaitableGetMountedFileResult',
    'get_mounted_file',
]

@pulumi.output_type
class GetMountedFileResult:
    """
    A collection of values returned by getMountedFile.
    """
    def __init__(__self__, checksum=None, content=None, context_id=None, id=None, module_id=None, relative_path=None, stack_id=None, write_only=None):
        if checksum and not isinstance(checksum, str):
            raise TypeError("Expected argument 'checksum' to be a str")
        pulumi.set(__self__, "checksum", checksum)
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        pulumi.set(__self__, "content", content)
        if context_id and not isinstance(context_id, str):
            raise TypeError("Expected argument 'context_id' to be a str")
        pulumi.set(__self__, "context_id", context_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if relative_path and not isinstance(relative_path, str):
            raise TypeError("Expected argument 'relative_path' to be a str")
        pulumi.set(__self__, "relative_path", relative_path)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if write_only and not isinstance(write_only, bool):
            raise TypeError("Expected argument 'write_only' to be a bool")
        pulumi.set(__self__, "write_only", write_only)

    @property
    @pulumi.getter
    def checksum(self) -> str:
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter
    def content(self) -> str:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contextId")
    def context_id(self) -> Optional[str]:
        return pulumi.get(self, "context_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[str]:
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter(name="relativePath")
    def relative_path(self) -> str:
        return pulumi.get(self, "relative_path")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[str]:
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="writeOnly")
    def write_only(self) -> bool:
        return pulumi.get(self, "write_only")


class AwaitableGetMountedFileResult(GetMountedFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMountedFileResult(
            checksum=self.checksum,
            content=self.content,
            context_id=self.context_id,
            id=self.id,
            module_id=self.module_id,
            relative_path=self.relative_path,
            stack_id=self.stack_id,
            write_only=self.write_only)


def get_mounted_file(context_id: Optional[str] = None,
                     module_id: Optional[str] = None,
                     relative_path: Optional[str] = None,
                     stack_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMountedFileResult:
    """
    `MountedFile` represents a file mounted in each Run's workspace that is part of a configuration of a context (`Context`), stack (`Stack`) or a module (`Module`). In principle, it's very similar to an environment variable (`EnvironmentVariable`) except that the value is written to the filesystem rather than passed to the environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    ireland_kubeconfig = spacelift.get_mounted_file(context_id="prod-k8s-ie",
        relative_path="kubeconfig")
    module_kubeconfig = spacelift.get_mounted_file(module_id="k8s-module",
        relative_path="kubeconfig")
    core_kubeconfig = spacelift.get_mounted_file(relative_path="kubeconfig",
        stack_id="k8s-core")
    ```

    <!-- schema generated by tfplugindocs -->
    ## Schema

    ### Required

    - **relative_path** (String) relative path to the mounted file

    ### Optional

    - **context_id** (String) ID of the context where the mounted file is stored
    - **id** (String) The ID of this resource.
    - **module_id** (String) ID of the module where the mounted file is stored
    - **stack_id** (String) ID of the stack where the mounted file is stored

    ### Read-Only

    - **checksum** (String) SHA-256 checksum of the value
    - **content** (String, Sensitive) content of the mounted file encoded using Base-64
    - **write_only** (Boolean) indicates whether the value can be read back outside a Run
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    __args__['moduleId'] = module_id
    __args__['relativePath'] = relative_path
    __args__['stackId'] = stack_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getMountedFile:getMountedFile', __args__, opts=opts, typ=GetMountedFileResult).value

    return AwaitableGetMountedFileResult(
        checksum=__ret__.checksum,
        content=__ret__.content,
        context_id=__ret__.context_id,
        id=__ret__.id,
        module_id=__ret__.module_id,
        relative_path=__ret__.relative_path,
        stack_id=__ret__.stack_id,
        write_only=__ret__.write_only)
