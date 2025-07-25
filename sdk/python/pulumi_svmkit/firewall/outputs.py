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
from ._enums import *

__all__ = [
    'FirewallParams',
]

@pulumi.output_type
class FirewallParams(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowPorts":
            suggest = "allow_ports"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FirewallParams. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FirewallParams.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FirewallParams.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_ports: Sequence[builtins.str]):
        pulumi.set(__self__, "allow_ports", allow_ports)

    @property
    @pulumi.getter(name="allowPorts")
    def allow_ports(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "allow_ports")


