# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetPoliciesResult',
    'AwaitableGetPoliciesResult',
    'get_policies',
    'get_policies_output',
]

@pulumi.output_type
class GetPoliciesResult:
    """
    A collection of values returned by getPolicies.
    """
    def __init__(__self__, id=None, labels=None, policies=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[str]]:
        """
        required labels to match
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetPoliciesPolicyResult']:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        required policy type
        """
        return pulumi.get(self, "type")


class AwaitableGetPoliciesResult(GetPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoliciesResult(
            id=self.id,
            labels=self.labels,
            policies=self.policies,
            type=self.type)


def get_policies(labels: Optional[Sequence[str]] = None,
                 type: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoliciesResult:
    """
    `get_policies` can find all policies that have certain labels.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    all = spacelift.get_policies()
    plan_autoattach = spacelift.get_policies(type="PLAN",
        labels=["autoattach"])
    pulumi.export("policyIds", [__item["id"] for __item in data["spacelift_policies"]["this"]["policies"]])
    ```
    """
    __args__ = dict()
    __args__['labels'] = labels
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getPolicies:getPolicies', __args__, opts=opts, typ=GetPoliciesResult).value

    return AwaitableGetPoliciesResult(
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        policies=pulumi.get(__ret__, 'policies'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_policies)
def get_policies_output(labels: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        type: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPoliciesResult]:
    """
    `get_policies` can find all policies that have certain labels.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    all = spacelift.get_policies()
    plan_autoattach = spacelift.get_policies(type="PLAN",
        labels=["autoattach"])
    pulumi.export("policyIds", [__item["id"] for __item in data["spacelift_policies"]["this"]["policies"]])
    ```
    """
    ...
