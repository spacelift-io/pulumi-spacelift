# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAwsIntegrationResult',
    'AwaitableGetAwsIntegrationResult',
    'get_aws_integration',
    'get_aws_integration_output',
]

@pulumi.output_type
class GetAwsIntegrationResult:
    """
    A collection of values returned by getAwsIntegration.
    """
    def __init__(__self__, duration_seconds=None, external_id=None, generate_credentials_in_worker=None, id=None, integration_id=None, labels=None, name=None, role_arn=None, space_id=None):
        if duration_seconds and not isinstance(duration_seconds, int):
            raise TypeError("Expected argument 'duration_seconds' to be a int")
        pulumi.set(__self__, "duration_seconds", duration_seconds)
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        pulumi.set(__self__, "external_id", external_id)
        if generate_credentials_in_worker and not isinstance(generate_credentials_in_worker, bool):
            raise TypeError("Expected argument 'generate_credentials_in_worker' to be a bool")
        pulumi.set(__self__, "generate_credentials_in_worker", generate_credentials_in_worker)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if integration_id and not isinstance(integration_id, str):
            raise TypeError("Expected argument 'integration_id' to be a str")
        pulumi.set(__self__, "integration_id", integration_id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> int:
        """
        Duration in seconds for which the assumed role credentials should be valid
        """
        return pulumi.get(self, "duration_seconds")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        """
        Custom external ID (works only for private workers).
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="generateCredentialsInWorker")
    def generate_credentials_in_worker(self) -> bool:
        """
        Generate AWS credentials in the private worker
        """
        return pulumi.get(self, "generate_credentials_in_worker")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> str:
        """
        Immutable ID of the integration. Either `integration_id` or `name` must be specified.
        """
        return pulumi.get(self, "integration_id")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the AWS integration. Either `integration_id` or `name` must be specified.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        ARN of the AWS IAM role to attach
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        ID (slug) of the space the integration is in
        """
        return pulumi.get(self, "space_id")


class AwaitableGetAwsIntegrationResult(GetAwsIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsIntegrationResult(
            duration_seconds=self.duration_seconds,
            external_id=self.external_id,
            generate_credentials_in_worker=self.generate_credentials_in_worker,
            id=self.id,
            integration_id=self.integration_id,
            labels=self.labels,
            name=self.name,
            role_arn=self.role_arn,
            space_id=self.space_id)


def get_aws_integration(integration_id: Optional[str] = None,
                        name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsIntegrationResult:
    """
    `AwsIntegration` represents an integration with an AWS account. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect.

    Note: when assuming credentials for **shared workers**, Spacelift will use `$accountName@$integrationID@$stackID@suffix` or `$accountName@$integrationID@$moduleID@$suffix` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole),$suffix will be `read` or `write`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    example = spacelift.get_aws_integration(name="Production")
    ```


    :param str integration_id: Immutable ID of the integration. Either `integration_id` or `name` must be specified.
    :param str name: Name of the AWS integration. Either `integration_id` or `name` must be specified.
    """
    __args__ = dict()
    __args__['integrationId'] = integration_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('spacelift:index/getAwsIntegration:getAwsIntegration', __args__, opts=opts, typ=GetAwsIntegrationResult).value

    return AwaitableGetAwsIntegrationResult(
        duration_seconds=pulumi.get(__ret__, 'duration_seconds'),
        external_id=pulumi.get(__ret__, 'external_id'),
        generate_credentials_in_worker=pulumi.get(__ret__, 'generate_credentials_in_worker'),
        id=pulumi.get(__ret__, 'id'),
        integration_id=pulumi.get(__ret__, 'integration_id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        space_id=pulumi.get(__ret__, 'space_id'))


@_utilities.lift_output_func(get_aws_integration)
def get_aws_integration_output(integration_id: Optional[pulumi.Input[Optional[str]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAwsIntegrationResult]:
    """
    `AwsIntegration` represents an integration with an AWS account. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect.

    Note: when assuming credentials for **shared workers**, Spacelift will use `$accountName@$integrationID@$stackID@suffix` or `$accountName@$integrationID@$moduleID@$suffix` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole),$suffix will be `read` or `write`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_spacelift as spacelift

    example = spacelift.get_aws_integration(name="Production")
    ```


    :param str integration_id: Immutable ID of the integration. Either `integration_id` or `name` must be specified.
    :param str name: Name of the AWS integration. Either `integration_id` or `name` must be specified.
    """
    ...
