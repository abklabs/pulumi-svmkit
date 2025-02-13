# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CpuGovernor',
    'TunerVariant',
]


class CpuGovernor(str, Enum):
    PERFORMANCE = "performance"
    """
    The performance governor
    """
    POWERSAVE = "powersave"
    """
    The powersave governor
    """
    ONDEMAND = "ondemand"
    """
    The ondemand governor
    """
    CONSERVATIVE = "conservative"
    """
    The conservative governor
    """
    SCHEDUTIL = "schedutil"
    """
    The schedutil governor
    """
    USERSPACE = "userspace"
    """
    The userspace governor
    """


class TunerVariant(str, Enum):
    GENERIC = "generic"
    """
    The generic tuner
    """
