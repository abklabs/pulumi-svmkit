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
    'ExplorerFlagsArgs',
    'ExplorerFlagsArgsDict',
]

MYPY = False

if not MYPY:
    class ExplorerFlagsArgsDict(TypedDict):
        hostname: NotRequired[pulumi.Input[builtins.str]]
        keep_alive_timeout: NotRequired[pulumi.Input[builtins.int]]
        port: NotRequired[pulumi.Input[builtins.int]]
elif False:
    ExplorerFlagsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ExplorerFlagsArgs:
    def __init__(__self__, *,
                 hostname: Optional[pulumi.Input[builtins.str]] = None,
                 keep_alive_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None):
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if keep_alive_timeout is not None:
            pulumi.set(__self__, "keep_alive_timeout", keep_alive_timeout)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="keepAliveTimeout")
    def keep_alive_timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "keep_alive_timeout")

    @keep_alive_timeout.setter
    def keep_alive_timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "keep_alive_timeout", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)


