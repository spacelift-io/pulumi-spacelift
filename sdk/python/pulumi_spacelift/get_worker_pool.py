# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetWorkerPoolResult',
    'AwaitableGetWorkerPoolResult',
    'get_worker_pool',
    'get_worker_pool_output',
]

@pulumi.output_type
class GetWorkerPoolResult:
    """
    A collection of values returned by getWorkerPool.
    """
    def __init__(__self__, config=None, description=None, id=None, labels=None, name=None, space_id=None, worker_pool_id=None):
        if config and not isinstance(config, str):
            raise TypeError("Expected argument 'config' to be a str")
        pulumi.set(__self__, "config", config)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)
        if worker_pool_id and not isinstance(worker_pool_id, str):
            raise TypeError("Expected argument 'worker_pool_id' to be a str")
        pulumi.set(__self__, "worker_pool_id", worker_pool_id)

    @property
    @pulumi.getter
    def config(self) -> str:
        """
        credentials necessary to connect WorkerPool's workers to the control plane
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        description of the worker pool
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        name of the worker pool
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        ID (slug) of the space the worker pool is in
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter(name="workerPoolId")
    def worker_pool_id(self) -> str:
        """
        ID of the worker pool
        """
        return pulumi.get(self, "worker_pool_id")


class AwaitableGetWorkerPoolResult(GetWorkerPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkerPoolResult(
            config=self.config,
            description=self.description,
            id=self.id,
            labels=self.labels,
            name=self.name,
            space_id=self.space_id,
            worker_pool_id=self.worker_pool_id)


def get_worker_pool(worker_pool_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkerPoolResult:
    """
    `WorkerPool` represents a worker pool assigned to the Spacelift account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    k8s_core = spacelift.get_worker_pool(worker_pool_id="k8s-core")
    ```


    :param str worker_pool_id: ID of the worker pool
    """
    __args__ = dict()
    __args__['workerPoolId'] = worker_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getWorkerPool:getWorkerPool', __args__, opts=opts, typ=GetWorkerPoolResult).value

    return AwaitableGetWorkerPoolResult(
        config=__ret__.config,
        description=__ret__.description,
        id=__ret__.id,
        labels=__ret__.labels,
        name=__ret__.name,
        space_id=__ret__.space_id,
        worker_pool_id=__ret__.worker_pool_id)


@_utilities.lift_output_func(get_worker_pool)
def get_worker_pool_output(worker_pool_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkerPoolResult]:
    """
    `WorkerPool` represents a worker pool assigned to the Spacelift account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    k8s_core = spacelift.get_worker_pool(worker_pool_id="k8s-core")
    ```


    :param str worker_pool_id: ID of the worker pool
    """
    ...
