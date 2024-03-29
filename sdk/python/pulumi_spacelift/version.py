# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VersionArgs', 'Version']

@pulumi.input_type
class VersionArgs:
    def __init__(__self__, *,
                 module_id: pulumi.Input[str],
                 commit_sha: Optional[pulumi.Input[str]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 version_number: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Version resource.
        :param pulumi.Input[str] module_id: ID of the module on which the version creation is to be triggered.
        :param pulumi.Input[str] commit_sha: The commit SHA for which to trigger a version.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of the resource.
        :param pulumi.Input[str] version_number: A semantic version number to set for the triggered version, example: 0.11.2
        """
        VersionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            module_id=module_id,
            commit_sha=commit_sha,
            keepers=keepers,
            version_number=version_number,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             module_id: Optional[pulumi.Input[str]] = None,
             commit_sha: Optional[pulumi.Input[str]] = None,
             keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             version_number: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if module_id is None and 'moduleId' in kwargs:
            module_id = kwargs['moduleId']
        if module_id is None:
            raise TypeError("Missing 'module_id' argument")
        if commit_sha is None and 'commitSha' in kwargs:
            commit_sha = kwargs['commitSha']
        if version_number is None and 'versionNumber' in kwargs:
            version_number = kwargs['versionNumber']

        _setter("module_id", module_id)
        if commit_sha is not None:
            _setter("commit_sha", commit_sha)
        if keepers is not None:
            _setter("keepers", keepers)
        if version_number is not None:
            _setter("version_number", version_number)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Input[str]:
        """
        ID of the module on which the version creation is to be triggered.
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter(name="commitSha")
    def commit_sha(self) -> Optional[pulumi.Input[str]]:
        """
        The commit SHA for which to trigger a version.
        """
        return pulumi.get(self, "commit_sha")

    @commit_sha.setter
    def commit_sha(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_sha", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of the resource.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> Optional[pulumi.Input[str]]:
        """
        A semantic version number to set for the triggered version, example: 0.11.2
        """
        return pulumi.get(self, "version_number")

    @version_number.setter
    def version_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_number", value)


@pulumi.input_type
class _VersionState:
    def __init__(__self__, *,
                 commit_sha: Optional[pulumi.Input[str]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 version_number: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Version resources.
        :param pulumi.Input[str] commit_sha: The commit SHA for which to trigger a version.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of the resource.
        :param pulumi.Input[str] module_id: ID of the module on which the version creation is to be triggered.
        :param pulumi.Input[str] version_number: A semantic version number to set for the triggered version, example: 0.11.2
        """
        _VersionState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            commit_sha=commit_sha,
            keepers=keepers,
            module_id=module_id,
            version_number=version_number,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             commit_sha: Optional[pulumi.Input[str]] = None,
             keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             module_id: Optional[pulumi.Input[str]] = None,
             version_number: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if commit_sha is None and 'commitSha' in kwargs:
            commit_sha = kwargs['commitSha']
        if module_id is None and 'moduleId' in kwargs:
            module_id = kwargs['moduleId']
        if version_number is None and 'versionNumber' in kwargs:
            version_number = kwargs['versionNumber']

        if commit_sha is not None:
            _setter("commit_sha", commit_sha)
        if keepers is not None:
            _setter("keepers", keepers)
        if module_id is not None:
            _setter("module_id", module_id)
        if version_number is not None:
            _setter("version_number", version_number)

    @property
    @pulumi.getter(name="commitSha")
    def commit_sha(self) -> Optional[pulumi.Input[str]]:
        """
        The commit SHA for which to trigger a version.
        """
        return pulumi.get(self, "commit_sha")

    @commit_sha.setter
    def commit_sha(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_sha", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of the resource.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the module on which the version creation is to be triggered.
        """
        return pulumi.get(self, "module_id")

    @module_id.setter
    def module_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "module_id", value)

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> Optional[pulumi.Input[str]]:
        """
        A semantic version number to set for the triggered version, example: 0.11.2
        """
        return pulumi.get(self, "version_number")

    @version_number.setter
    def version_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_number", value)


class Version(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 commit_sha: Optional[pulumi.Input[str]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 version_number: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `Version` allows to programmatically trigger a version creation in response to arbitrary changes in the keepers section.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] commit_sha: The commit SHA for which to trigger a version.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of the resource.
        :param pulumi.Input[str] module_id: ID of the module on which the version creation is to be triggered.
        :param pulumi.Input[str] version_number: A semantic version number to set for the triggered version, example: 0.11.2
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `Version` allows to programmatically trigger a version creation in response to arbitrary changes in the keepers section.

        :param str resource_name: The name of the resource.
        :param VersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VersionArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 commit_sha: Optional[pulumi.Input[str]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 module_id: Optional[pulumi.Input[str]] = None,
                 version_number: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VersionArgs.__new__(VersionArgs)

            __props__.__dict__["commit_sha"] = commit_sha
            __props__.__dict__["keepers"] = keepers
            if module_id is None and not opts.urn:
                raise TypeError("Missing required property 'module_id'")
            __props__.__dict__["module_id"] = module_id
            __props__.__dict__["version_number"] = version_number
        super(Version, __self__).__init__(
            'spacelift:index/version:Version',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            commit_sha: Optional[pulumi.Input[str]] = None,
            keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            module_id: Optional[pulumi.Input[str]] = None,
            version_number: Optional[pulumi.Input[str]] = None) -> 'Version':
        """
        Get an existing Version resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] commit_sha: The commit SHA for which to trigger a version.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of the resource.
        :param pulumi.Input[str] module_id: ID of the module on which the version creation is to be triggered.
        :param pulumi.Input[str] version_number: A semantic version number to set for the triggered version, example: 0.11.2
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VersionState.__new__(_VersionState)

        __props__.__dict__["commit_sha"] = commit_sha
        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["module_id"] = module_id
        __props__.__dict__["version_number"] = version_number
        return Version(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="commitSha")
    def commit_sha(self) -> pulumi.Output[Optional[str]]:
        """
        The commit SHA for which to trigger a version.
        """
        return pulumi.get(self, "commit_sha")

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of the resource.
        """
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> pulumi.Output[str]:
        """
        ID of the module on which the version creation is to be triggered.
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> pulumi.Output[Optional[str]]:
        """
        A semantic version number to set for the triggered version, example: 0.11.2
        """
        return pulumi.get(self, "version_number")

