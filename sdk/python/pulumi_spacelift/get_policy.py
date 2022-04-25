# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetPolicyResult',
    'AwaitableGetPolicyResult',
    'get_policy',
]

@pulumi.output_type
class GetPolicyResult:
    """
    A collection of values returned by getPolicy.
    """
    def __init__(__self__, body=None, id=None, labels=None, name=None, policy_id=None, type=None):
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        pulumi.set(__self__, "policy_id", policy_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def body(self) -> str:
        return pulumi.get(self, "body")

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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetPolicyResult(GetPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyResult(
            body=self.body,
            id=self.id,
            labels=self.labels,
            name=self.name,
            policy_id=self.policy_id,
            type=self.type)


def get_policy(policy_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyResult:
    """
    `Policy` represents a Spacelift **policy** - a collection of customer-defined rules that are applied by Spacelift at one of the decision points within the application.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    policy = spacelift.get_policy(policy_id=spacelift_policy["policy"]["id"])
    pulumi.export("policyBody", policy.body)
    ```

    <!-- schema generated by tfplugindocs -->
    ## Schema

    ### Required

    - **policy_id** (String) immutable ID (slug) of the policy

    ### Optional

    - **id** (String) The ID of this resource.

    ### Read-Only

    - **body** (String) body of the policy
    - **labels** (Set of String)
    - **name** (String) name of the policy
    - **type** (String) type of the policy
    """
    __args__ = dict()
    __args__['policyId'] = policy_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('spacelift:index/getPolicy:getPolicy', __args__, opts=opts, typ=GetPolicyResult).value

    return AwaitableGetPolicyResult(
        body=__ret__.body,
        id=__ret__.id,
        labels=__ret__.labels,
        name=__ret__.name,
        policy_id=__ret__.policy_id,
        type=__ret__.type)
