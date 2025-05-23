# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'PackageConfigArgs',
    'PackageConfigArgsDict',
    'PackageArgs',
    'PackageArgsDict',
]

MYPY = False

if not MYPY:
    class PackageConfigArgsDict(TypedDict):
        additional: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        override: NotRequired[pulumi.Input[Sequence[pulumi.Input['PackageArgsDict']]]]
elif False:
    PackageConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PackageConfigArgs:
    def __init__(__self__, *,
                 additional: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 override: Optional[pulumi.Input[Sequence[pulumi.Input['PackageArgs']]]] = None):
        if additional is not None:
            pulumi.set(__self__, "additional", additional)
        if override is not None:
            pulumi.set(__self__, "override", override)

    @property
    @pulumi.getter
    def additional(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "additional")

    @additional.setter
    def additional(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "additional", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PackageArgs']]]]:
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PackageArgs']]]]):
        pulumi.set(self, "override", value)


if not MYPY:
    class PackageArgsDict(TypedDict):
        name: pulumi.Input[builtins.str]
        path: NotRequired[pulumi.Input[builtins.str]]
        target_release: NotRequired[pulumi.Input[builtins.str]]
        version: NotRequired[pulumi.Input[builtins.str]]
elif False:
    PackageArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PackageArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 target_release: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None):
        pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if target_release is not None:
            pulumi.set(__self__, "target_release", target_release)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="targetRelease")
    def target_release(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "target_release")

    @target_release.setter
    def target_release(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "target_release", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


