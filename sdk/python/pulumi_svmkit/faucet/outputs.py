# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'FaucetFlags',
]

@pulumi.output_type
class FaucetFlags(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowIPs":
            suggest = "allow_ips"
        elif key == "perRequestCap":
            suggest = "per_request_cap"
        elif key == "perTimeCap":
            suggest = "per_time_cap"
        elif key == "sliceSeconds":
            suggest = "slice_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FaucetFlags. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FaucetFlags.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FaucetFlags.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_ips: Optional[Sequence[str]] = None,
                 per_request_cap: Optional[int] = None,
                 per_time_cap: Optional[int] = None,
                 slice_seconds: Optional[int] = None):
        if allow_ips is not None:
            pulumi.set(__self__, "allow_ips", allow_ips)
        if per_request_cap is not None:
            pulumi.set(__self__, "per_request_cap", per_request_cap)
        if per_time_cap is not None:
            pulumi.set(__self__, "per_time_cap", per_time_cap)
        if slice_seconds is not None:
            pulumi.set(__self__, "slice_seconds", slice_seconds)

    @property
    @pulumi.getter(name="allowIPs")
    def allow_ips(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "allow_ips")

    @property
    @pulumi.getter(name="perRequestCap")
    def per_request_cap(self) -> Optional[int]:
        return pulumi.get(self, "per_request_cap")

    @property
    @pulumi.getter(name="perTimeCap")
    def per_time_cap(self) -> Optional[int]:
        return pulumi.get(self, "per_time_cap")

    @property
    @pulumi.getter(name="sliceSeconds")
    def slice_seconds(self) -> Optional[int]:
        return pulumi.get(self, "slice_seconds")


