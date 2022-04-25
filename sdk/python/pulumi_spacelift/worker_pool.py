# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['WorkerPool']


class WorkerPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `WorkerPool` represents a worker pool assigned to the Spacelift account.

        ## Schema

        ### Required

        - **name** (String) name of the worker pool

        ### Optional

        - **csr** (String, Sensitive) certificate signing request in base64
        - **description** (String) description of the worker pool
        - **id** (String) The ID of this resource.
        - **labels** (Set of String)

        ### Read-Only

        - **config** (String, Sensitive) credentials necessary to connect WorkerPool's workers to the control plane
        - **private_key** (String, Sensitive) private key in base64

        ## Import

        Import is supported using the following syntax

        ```sh
         $ pulumi import spacelift:index/workerPool:WorkerPool k8s-core $WORKER_POOL_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] csr: certificate signing request in base64
        :param pulumi.Input[str] description: description of the worker pool
        :param pulumi.Input[str] name: name of the worker pool
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

            __props__['csr'] = csr
            __props__['description'] = description
            __props__['labels'] = labels
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['config'] = None
            __props__['private_key'] = None
        super(WorkerPool, __self__).__init__(
            'spacelift:index/workerPool:WorkerPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            csr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None) -> 'WorkerPool':
        """
        Get an existing WorkerPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: credentials necessary to connect WorkerPool's workers to the control plane
        :param pulumi.Input[str] csr: certificate signing request in base64
        :param pulumi.Input[str] description: description of the worker pool
        :param pulumi.Input[str] name: name of the worker pool
        :param pulumi.Input[str] private_key: private key in base64
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config"] = config
        __props__["csr"] = csr
        __props__["description"] = description
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["private_key"] = private_key
        return WorkerPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        credentials necessary to connect WorkerPool's workers to the control plane
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Output[str]:
        """
        certificate signing request in base64
        """
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        description of the worker pool
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        name of the worker pool
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        private key in base64
        """
        return pulumi.get(self, "private_key")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

