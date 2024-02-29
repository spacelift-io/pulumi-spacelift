# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ContextArgs', 'Context']

@pulumi.input_type
class ContextArgs:
    def __init__(__self__, *,
                 after_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_runs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Context resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_applies: List of after-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_destroys: List of after-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_inits: List of after-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_performs: List of after-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_plans: List of after-plan scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_runs: List of after-run scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_applies: List of before-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_destroys: List of before-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_inits: List of before-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_performs: List of before-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_plans: List of before-plan scripts
        :param pulumi.Input[str] description: Free-form context description for users
        :param pulumi.Input[str] name: Name of the context - should be unique in one account
        :param pulumi.Input[str] space_id: ID (slug) of the space the context is in
        """
        if after_applies is not None:
            pulumi.set(__self__, "after_applies", after_applies)
        if after_destroys is not None:
            pulumi.set(__self__, "after_destroys", after_destroys)
        if after_inits is not None:
            pulumi.set(__self__, "after_inits", after_inits)
        if after_performs is not None:
            pulumi.set(__self__, "after_performs", after_performs)
        if after_plans is not None:
            pulumi.set(__self__, "after_plans", after_plans)
        if after_runs is not None:
            pulumi.set(__self__, "after_runs", after_runs)
        if before_applies is not None:
            pulumi.set(__self__, "before_applies", before_applies)
        if before_destroys is not None:
            pulumi.set(__self__, "before_destroys", before_destroys)
        if before_inits is not None:
            pulumi.set(__self__, "before_inits", before_inits)
        if before_performs is not None:
            pulumi.set(__self__, "before_performs", before_performs)
        if before_plans is not None:
            pulumi.set(__self__, "before_plans", before_plans)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter(name="afterApplies")
    def after_applies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-apply scripts
        """
        return pulumi.get(self, "after_applies")

    @after_applies.setter
    def after_applies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_applies", value)

    @property
    @pulumi.getter(name="afterDestroys")
    def after_destroys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-destroy scripts
        """
        return pulumi.get(self, "after_destroys")

    @after_destroys.setter
    def after_destroys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_destroys", value)

    @property
    @pulumi.getter(name="afterInits")
    def after_inits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-init scripts
        """
        return pulumi.get(self, "after_inits")

    @after_inits.setter
    def after_inits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_inits", value)

    @property
    @pulumi.getter(name="afterPerforms")
    def after_performs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-perform scripts
        """
        return pulumi.get(self, "after_performs")

    @after_performs.setter
    def after_performs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_performs", value)

    @property
    @pulumi.getter(name="afterPlans")
    def after_plans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-plan scripts
        """
        return pulumi.get(self, "after_plans")

    @after_plans.setter
    def after_plans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_plans", value)

    @property
    @pulumi.getter(name="afterRuns")
    def after_runs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-run scripts
        """
        return pulumi.get(self, "after_runs")

    @after_runs.setter
    def after_runs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_runs", value)

    @property
    @pulumi.getter(name="beforeApplies")
    def before_applies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-apply scripts
        """
        return pulumi.get(self, "before_applies")

    @before_applies.setter
    def before_applies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_applies", value)

    @property
    @pulumi.getter(name="beforeDestroys")
    def before_destroys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-destroy scripts
        """
        return pulumi.get(self, "before_destroys")

    @before_destroys.setter
    def before_destroys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_destroys", value)

    @property
    @pulumi.getter(name="beforeInits")
    def before_inits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-init scripts
        """
        return pulumi.get(self, "before_inits")

    @before_inits.setter
    def before_inits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_inits", value)

    @property
    @pulumi.getter(name="beforePerforms")
    def before_performs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-perform scripts
        """
        return pulumi.get(self, "before_performs")

    @before_performs.setter
    def before_performs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_performs", value)

    @property
    @pulumi.getter(name="beforePlans")
    def before_plans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-plan scripts
        """
        return pulumi.get(self, "before_plans")

    @before_plans.setter
    def before_plans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_plans", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-form context description for users
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the context - should be unique in one account
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID (slug) of the space the context is in
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


@pulumi.input_type
class _ContextState:
    def __init__(__self__, *,
                 after_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_runs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Context resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_applies: List of after-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_destroys: List of after-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_inits: List of after-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_performs: List of after-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_plans: List of after-plan scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_runs: List of after-run scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_applies: List of before-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_destroys: List of before-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_inits: List of before-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_performs: List of before-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_plans: List of before-plan scripts
        :param pulumi.Input[str] description: Free-form context description for users
        :param pulumi.Input[str] name: Name of the context - should be unique in one account
        :param pulumi.Input[str] space_id: ID (slug) of the space the context is in
        """
        if after_applies is not None:
            pulumi.set(__self__, "after_applies", after_applies)
        if after_destroys is not None:
            pulumi.set(__self__, "after_destroys", after_destroys)
        if after_inits is not None:
            pulumi.set(__self__, "after_inits", after_inits)
        if after_performs is not None:
            pulumi.set(__self__, "after_performs", after_performs)
        if after_plans is not None:
            pulumi.set(__self__, "after_plans", after_plans)
        if after_runs is not None:
            pulumi.set(__self__, "after_runs", after_runs)
        if before_applies is not None:
            pulumi.set(__self__, "before_applies", before_applies)
        if before_destroys is not None:
            pulumi.set(__self__, "before_destroys", before_destroys)
        if before_inits is not None:
            pulumi.set(__self__, "before_inits", before_inits)
        if before_performs is not None:
            pulumi.set(__self__, "before_performs", before_performs)
        if before_plans is not None:
            pulumi.set(__self__, "before_plans", before_plans)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter(name="afterApplies")
    def after_applies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-apply scripts
        """
        return pulumi.get(self, "after_applies")

    @after_applies.setter
    def after_applies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_applies", value)

    @property
    @pulumi.getter(name="afterDestroys")
    def after_destroys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-destroy scripts
        """
        return pulumi.get(self, "after_destroys")

    @after_destroys.setter
    def after_destroys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_destroys", value)

    @property
    @pulumi.getter(name="afterInits")
    def after_inits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-init scripts
        """
        return pulumi.get(self, "after_inits")

    @after_inits.setter
    def after_inits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_inits", value)

    @property
    @pulumi.getter(name="afterPerforms")
    def after_performs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-perform scripts
        """
        return pulumi.get(self, "after_performs")

    @after_performs.setter
    def after_performs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_performs", value)

    @property
    @pulumi.getter(name="afterPlans")
    def after_plans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-plan scripts
        """
        return pulumi.get(self, "after_plans")

    @after_plans.setter
    def after_plans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_plans", value)

    @property
    @pulumi.getter(name="afterRuns")
    def after_runs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of after-run scripts
        """
        return pulumi.get(self, "after_runs")

    @after_runs.setter
    def after_runs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after_runs", value)

    @property
    @pulumi.getter(name="beforeApplies")
    def before_applies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-apply scripts
        """
        return pulumi.get(self, "before_applies")

    @before_applies.setter
    def before_applies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_applies", value)

    @property
    @pulumi.getter(name="beforeDestroys")
    def before_destroys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-destroy scripts
        """
        return pulumi.get(self, "before_destroys")

    @before_destroys.setter
    def before_destroys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_destroys", value)

    @property
    @pulumi.getter(name="beforeInits")
    def before_inits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-init scripts
        """
        return pulumi.get(self, "before_inits")

    @before_inits.setter
    def before_inits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_inits", value)

    @property
    @pulumi.getter(name="beforePerforms")
    def before_performs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-perform scripts
        """
        return pulumi.get(self, "before_performs")

    @before_performs.setter
    def before_performs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_performs", value)

    @property
    @pulumi.getter(name="beforePlans")
    def before_plans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of before-plan scripts
        """
        return pulumi.get(self, "before_plans")

    @before_plans.setter
    def before_plans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before_plans", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-form context description for users
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the context - should be unique in one account
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID (slug) of the space the context is in
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


class Context(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 after_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_runs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`Stack`) or modules (`Module`) using a context attachment (`ContextAttachment`)`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        prod_k8s_ie = spacelift.Context("prod-k8s-ie", description="Configuration details for the compute cluster in 🇮🇪")
        ```

        ## Import

        ```sh
         $ pulumi import spacelift:index/context:Context prod-k8s-ie $CONTEXT_ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_applies: List of after-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_destroys: List of after-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_inits: List of after-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_performs: List of after-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_plans: List of after-plan scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_runs: List of after-run scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_applies: List of before-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_destroys: List of before-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_inits: List of before-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_performs: List of before-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_plans: List of before-plan scripts
        :param pulumi.Input[str] description: Free-form context description for users
        :param pulumi.Input[str] name: Name of the context - should be unique in one account
        :param pulumi.Input[str] space_id: ID (slug) of the space the context is in
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContextArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `Context` represents a Spacelift **context** - a collection of configuration elements (either environment variables or mounted files) that can be administratively attached to multiple stacks (`Stack`) or modules (`Module`) using a context attachment (`ContextAttachment`)`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_spacelift as spacelift

        prod_k8s_ie = spacelift.Context("prod-k8s-ie", description="Configuration details for the compute cluster in 🇮🇪")
        ```

        ## Import

        ```sh
         $ pulumi import spacelift:index/context:Context prod-k8s-ie $CONTEXT_ID
        ```

        :param str resource_name: The name of the resource.
        :param ContextArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContextArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 after_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 after_runs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContextArgs.__new__(ContextArgs)

            __props__.__dict__["after_applies"] = after_applies
            __props__.__dict__["after_destroys"] = after_destroys
            __props__.__dict__["after_inits"] = after_inits
            __props__.__dict__["after_performs"] = after_performs
            __props__.__dict__["after_plans"] = after_plans
            __props__.__dict__["after_runs"] = after_runs
            __props__.__dict__["before_applies"] = before_applies
            __props__.__dict__["before_destroys"] = before_destroys
            __props__.__dict__["before_inits"] = before_inits
            __props__.__dict__["before_performs"] = before_performs
            __props__.__dict__["before_plans"] = before_plans
            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["space_id"] = space_id
        super(Context, __self__).__init__(
            'spacelift:index/context:Context',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            after_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            after_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            after_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            after_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            after_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            after_runs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            before_applies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            before_destroys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            before_inits: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            before_performs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            before_plans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None) -> 'Context':
        """
        Get an existing Context resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_applies: List of after-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_destroys: List of after-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_inits: List of after-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_performs: List of after-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_plans: List of after-plan scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after_runs: List of after-run scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_applies: List of before-apply scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_destroys: List of before-destroy scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_inits: List of before-init scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_performs: List of before-perform scripts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before_plans: List of before-plan scripts
        :param pulumi.Input[str] description: Free-form context description for users
        :param pulumi.Input[str] name: Name of the context - should be unique in one account
        :param pulumi.Input[str] space_id: ID (slug) of the space the context is in
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContextState.__new__(_ContextState)

        __props__.__dict__["after_applies"] = after_applies
        __props__.__dict__["after_destroys"] = after_destroys
        __props__.__dict__["after_inits"] = after_inits
        __props__.__dict__["after_performs"] = after_performs
        __props__.__dict__["after_plans"] = after_plans
        __props__.__dict__["after_runs"] = after_runs
        __props__.__dict__["before_applies"] = before_applies
        __props__.__dict__["before_destroys"] = before_destroys
        __props__.__dict__["before_inits"] = before_inits
        __props__.__dict__["before_performs"] = before_performs
        __props__.__dict__["before_plans"] = before_plans
        __props__.__dict__["description"] = description
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["space_id"] = space_id
        return Context(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="afterApplies")
    def after_applies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of after-apply scripts
        """
        return pulumi.get(self, "after_applies")

    @property
    @pulumi.getter(name="afterDestroys")
    def after_destroys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of after-destroy scripts
        """
        return pulumi.get(self, "after_destroys")

    @property
    @pulumi.getter(name="afterInits")
    def after_inits(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of after-init scripts
        """
        return pulumi.get(self, "after_inits")

    @property
    @pulumi.getter(name="afterPerforms")
    def after_performs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of after-perform scripts
        """
        return pulumi.get(self, "after_performs")

    @property
    @pulumi.getter(name="afterPlans")
    def after_plans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of after-plan scripts
        """
        return pulumi.get(self, "after_plans")

    @property
    @pulumi.getter(name="afterRuns")
    def after_runs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of after-run scripts
        """
        return pulumi.get(self, "after_runs")

    @property
    @pulumi.getter(name="beforeApplies")
    def before_applies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of before-apply scripts
        """
        return pulumi.get(self, "before_applies")

    @property
    @pulumi.getter(name="beforeDestroys")
    def before_destroys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of before-destroy scripts
        """
        return pulumi.get(self, "before_destroys")

    @property
    @pulumi.getter(name="beforeInits")
    def before_inits(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of before-init scripts
        """
        return pulumi.get(self, "before_inits")

    @property
    @pulumi.getter(name="beforePerforms")
    def before_performs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of before-perform scripts
        """
        return pulumi.get(self, "before_performs")

    @property
    @pulumi.getter(name="beforePlans")
    def before_plans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of before-plan scripts
        """
        return pulumi.get(self, "before_plans")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Free-form context description for users
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
        Name of the context - should be unique in one account
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        ID (slug) of the space the context is in
        """
        return pulumi.get(self, "space_id")

