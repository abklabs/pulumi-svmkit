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
from ._enums import *

__all__ = [
    'TunerFsParamsArgs',
    'TunerFsParamsArgsDict',
    'TunerKernelParamsArgs',
    'TunerKernelParamsArgsDict',
    'TunerNetParamsArgs',
    'TunerNetParamsArgsDict',
    'TunerParamsArgs',
    'TunerParamsArgsDict',
    'TunerVmParamsArgs',
    'TunerVmParamsArgsDict',
]

MYPY = False

if not MYPY:
    class TunerFsParamsArgsDict(TypedDict):
        fs_nr_open: NotRequired[pulumi.Input[int]]
elif False:
    TunerFsParamsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TunerFsParamsArgs:
    def __init__(__self__, *,
                 fs_nr_open: Optional[pulumi.Input[int]] = None):
        if fs_nr_open is not None:
            pulumi.set(__self__, "fs_nr_open", fs_nr_open)

    @property
    @pulumi.getter(name="fsNrOpen")
    def fs_nr_open(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fs_nr_open")

    @fs_nr_open.setter
    def fs_nr_open(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fs_nr_open", value)


if not MYPY:
    class TunerKernelParamsArgsDict(TypedDict):
        kernel_hung_task_timeout_secs: NotRequired[pulumi.Input[int]]
        kernel_nmi_watchdog: NotRequired[pulumi.Input[int]]
        kernel_pid_max: NotRequired[pulumi.Input[int]]
        kernel_sched_min_granularity_ns: NotRequired[pulumi.Input[int]]
        kernel_sched_wakeup_granularity_ns: NotRequired[pulumi.Input[int]]
        kernel_timer_migration: NotRequired[pulumi.Input[int]]
elif False:
    TunerKernelParamsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TunerKernelParamsArgs:
    def __init__(__self__, *,
                 kernel_hung_task_timeout_secs: Optional[pulumi.Input[int]] = None,
                 kernel_nmi_watchdog: Optional[pulumi.Input[int]] = None,
                 kernel_pid_max: Optional[pulumi.Input[int]] = None,
                 kernel_sched_min_granularity_ns: Optional[pulumi.Input[int]] = None,
                 kernel_sched_wakeup_granularity_ns: Optional[pulumi.Input[int]] = None,
                 kernel_timer_migration: Optional[pulumi.Input[int]] = None):
        if kernel_hung_task_timeout_secs is not None:
            pulumi.set(__self__, "kernel_hung_task_timeout_secs", kernel_hung_task_timeout_secs)
        if kernel_nmi_watchdog is not None:
            pulumi.set(__self__, "kernel_nmi_watchdog", kernel_nmi_watchdog)
        if kernel_pid_max is not None:
            pulumi.set(__self__, "kernel_pid_max", kernel_pid_max)
        if kernel_sched_min_granularity_ns is not None:
            pulumi.set(__self__, "kernel_sched_min_granularity_ns", kernel_sched_min_granularity_ns)
        if kernel_sched_wakeup_granularity_ns is not None:
            pulumi.set(__self__, "kernel_sched_wakeup_granularity_ns", kernel_sched_wakeup_granularity_ns)
        if kernel_timer_migration is not None:
            pulumi.set(__self__, "kernel_timer_migration", kernel_timer_migration)

    @property
    @pulumi.getter(name="kernelHungTaskTimeoutSecs")
    def kernel_hung_task_timeout_secs(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "kernel_hung_task_timeout_secs")

    @kernel_hung_task_timeout_secs.setter
    def kernel_hung_task_timeout_secs(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kernel_hung_task_timeout_secs", value)

    @property
    @pulumi.getter(name="kernelNmiWatchdog")
    def kernel_nmi_watchdog(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "kernel_nmi_watchdog")

    @kernel_nmi_watchdog.setter
    def kernel_nmi_watchdog(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kernel_nmi_watchdog", value)

    @property
    @pulumi.getter(name="kernelPidMax")
    def kernel_pid_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "kernel_pid_max")

    @kernel_pid_max.setter
    def kernel_pid_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kernel_pid_max", value)

    @property
    @pulumi.getter(name="kernelSchedMinGranularityNs")
    def kernel_sched_min_granularity_ns(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "kernel_sched_min_granularity_ns")

    @kernel_sched_min_granularity_ns.setter
    def kernel_sched_min_granularity_ns(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kernel_sched_min_granularity_ns", value)

    @property
    @pulumi.getter(name="kernelSchedWakeupGranularityNs")
    def kernel_sched_wakeup_granularity_ns(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "kernel_sched_wakeup_granularity_ns")

    @kernel_sched_wakeup_granularity_ns.setter
    def kernel_sched_wakeup_granularity_ns(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kernel_sched_wakeup_granularity_ns", value)

    @property
    @pulumi.getter(name="kernelTimerMigration")
    def kernel_timer_migration(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "kernel_timer_migration")

    @kernel_timer_migration.setter
    def kernel_timer_migration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kernel_timer_migration", value)


if not MYPY:
    class TunerNetParamsArgsDict(TypedDict):
        net_core_rmem_default: NotRequired[pulumi.Input[int]]
        net_core_rmem_max: NotRequired[pulumi.Input[int]]
        net_core_wmem_default: NotRequired[pulumi.Input[int]]
        net_core_wmem_max: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_congestion_control: NotRequired[pulumi.Input[str]]
        net_ipv4_tcp_fastopen: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_low_latency: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_moderate_rcvbuf: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_no_metrics_save: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_rmem: NotRequired[pulumi.Input[str]]
        net_ipv4_tcp_sack: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_timestamps: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_tw_reuse: NotRequired[pulumi.Input[int]]
        net_ipv4_tcp_wmem: NotRequired[pulumi.Input[str]]
elif False:
    TunerNetParamsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TunerNetParamsArgs:
    def __init__(__self__, *,
                 net_core_rmem_default: Optional[pulumi.Input[int]] = None,
                 net_core_rmem_max: Optional[pulumi.Input[int]] = None,
                 net_core_wmem_default: Optional[pulumi.Input[int]] = None,
                 net_core_wmem_max: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_congestion_control: Optional[pulumi.Input[str]] = None,
                 net_ipv4_tcp_fastopen: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_low_latency: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_moderate_rcvbuf: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_no_metrics_save: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_rmem: Optional[pulumi.Input[str]] = None,
                 net_ipv4_tcp_sack: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_timestamps: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_tw_reuse: Optional[pulumi.Input[int]] = None,
                 net_ipv4_tcp_wmem: Optional[pulumi.Input[str]] = None):
        if net_core_rmem_default is not None:
            pulumi.set(__self__, "net_core_rmem_default", net_core_rmem_default)
        if net_core_rmem_max is not None:
            pulumi.set(__self__, "net_core_rmem_max", net_core_rmem_max)
        if net_core_wmem_default is not None:
            pulumi.set(__self__, "net_core_wmem_default", net_core_wmem_default)
        if net_core_wmem_max is not None:
            pulumi.set(__self__, "net_core_wmem_max", net_core_wmem_max)
        if net_ipv4_tcp_congestion_control is not None:
            pulumi.set(__self__, "net_ipv4_tcp_congestion_control", net_ipv4_tcp_congestion_control)
        if net_ipv4_tcp_fastopen is not None:
            pulumi.set(__self__, "net_ipv4_tcp_fastopen", net_ipv4_tcp_fastopen)
        if net_ipv4_tcp_low_latency is not None:
            pulumi.set(__self__, "net_ipv4_tcp_low_latency", net_ipv4_tcp_low_latency)
        if net_ipv4_tcp_moderate_rcvbuf is not None:
            pulumi.set(__self__, "net_ipv4_tcp_moderate_rcvbuf", net_ipv4_tcp_moderate_rcvbuf)
        if net_ipv4_tcp_no_metrics_save is not None:
            pulumi.set(__self__, "net_ipv4_tcp_no_metrics_save", net_ipv4_tcp_no_metrics_save)
        if net_ipv4_tcp_rmem is not None:
            pulumi.set(__self__, "net_ipv4_tcp_rmem", net_ipv4_tcp_rmem)
        if net_ipv4_tcp_sack is not None:
            pulumi.set(__self__, "net_ipv4_tcp_sack", net_ipv4_tcp_sack)
        if net_ipv4_tcp_timestamps is not None:
            pulumi.set(__self__, "net_ipv4_tcp_timestamps", net_ipv4_tcp_timestamps)
        if net_ipv4_tcp_tw_reuse is not None:
            pulumi.set(__self__, "net_ipv4_tcp_tw_reuse", net_ipv4_tcp_tw_reuse)
        if net_ipv4_tcp_wmem is not None:
            pulumi.set(__self__, "net_ipv4_tcp_wmem", net_ipv4_tcp_wmem)

    @property
    @pulumi.getter(name="netCoreRmemDefault")
    def net_core_rmem_default(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_core_rmem_default")

    @net_core_rmem_default.setter
    def net_core_rmem_default(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_core_rmem_default", value)

    @property
    @pulumi.getter(name="netCoreRmemMax")
    def net_core_rmem_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_core_rmem_max")

    @net_core_rmem_max.setter
    def net_core_rmem_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_core_rmem_max", value)

    @property
    @pulumi.getter(name="netCoreWmemDefault")
    def net_core_wmem_default(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_core_wmem_default")

    @net_core_wmem_default.setter
    def net_core_wmem_default(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_core_wmem_default", value)

    @property
    @pulumi.getter(name="netCoreWmemMax")
    def net_core_wmem_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_core_wmem_max")

    @net_core_wmem_max.setter
    def net_core_wmem_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_core_wmem_max", value)

    @property
    @pulumi.getter(name="netIpv4TcpCongestionControl")
    def net_ipv4_tcp_congestion_control(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "net_ipv4_tcp_congestion_control")

    @net_ipv4_tcp_congestion_control.setter
    def net_ipv4_tcp_congestion_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_ipv4_tcp_congestion_control", value)

    @property
    @pulumi.getter(name="netIpv4TcpFastopen")
    def net_ipv4_tcp_fastopen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_fastopen")

    @net_ipv4_tcp_fastopen.setter
    def net_ipv4_tcp_fastopen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_fastopen", value)

    @property
    @pulumi.getter(name="netIpv4TcpLowLatency")
    def net_ipv4_tcp_low_latency(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_low_latency")

    @net_ipv4_tcp_low_latency.setter
    def net_ipv4_tcp_low_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_low_latency", value)

    @property
    @pulumi.getter(name="netIpv4TcpModerateRcvbuf")
    def net_ipv4_tcp_moderate_rcvbuf(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_moderate_rcvbuf")

    @net_ipv4_tcp_moderate_rcvbuf.setter
    def net_ipv4_tcp_moderate_rcvbuf(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_moderate_rcvbuf", value)

    @property
    @pulumi.getter(name="netIpv4TcpNoMetricsSave")
    def net_ipv4_tcp_no_metrics_save(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_no_metrics_save")

    @net_ipv4_tcp_no_metrics_save.setter
    def net_ipv4_tcp_no_metrics_save(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_no_metrics_save", value)

    @property
    @pulumi.getter(name="netIpv4TcpRmem")
    def net_ipv4_tcp_rmem(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "net_ipv4_tcp_rmem")

    @net_ipv4_tcp_rmem.setter
    def net_ipv4_tcp_rmem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_ipv4_tcp_rmem", value)

    @property
    @pulumi.getter(name="netIpv4TcpSack")
    def net_ipv4_tcp_sack(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_sack")

    @net_ipv4_tcp_sack.setter
    def net_ipv4_tcp_sack(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_sack", value)

    @property
    @pulumi.getter(name="netIpv4TcpTimestamps")
    def net_ipv4_tcp_timestamps(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_timestamps")

    @net_ipv4_tcp_timestamps.setter
    def net_ipv4_tcp_timestamps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_timestamps", value)

    @property
    @pulumi.getter(name="netIpv4TcpTwReuse")
    def net_ipv4_tcp_tw_reuse(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "net_ipv4_tcp_tw_reuse")

    @net_ipv4_tcp_tw_reuse.setter
    def net_ipv4_tcp_tw_reuse(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "net_ipv4_tcp_tw_reuse", value)

    @property
    @pulumi.getter(name="netIpv4TcpWmem")
    def net_ipv4_tcp_wmem(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "net_ipv4_tcp_wmem")

    @net_ipv4_tcp_wmem.setter
    def net_ipv4_tcp_wmem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_ipv4_tcp_wmem", value)


if not MYPY:
    class TunerParamsArgsDict(TypedDict):
        cpu_governor: NotRequired[pulumi.Input['CpuGovernor']]
        fs: NotRequired[pulumi.Input['TunerFsParamsArgsDict']]
        kernel: NotRequired[pulumi.Input['TunerKernelParamsArgsDict']]
        net: NotRequired[pulumi.Input['TunerNetParamsArgsDict']]
        vm: NotRequired[pulumi.Input['TunerVmParamsArgsDict']]
elif False:
    TunerParamsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TunerParamsArgs:
    def __init__(__self__, *,
                 cpu_governor: Optional[pulumi.Input['CpuGovernor']] = None,
                 fs: Optional[pulumi.Input['TunerFsParamsArgs']] = None,
                 kernel: Optional[pulumi.Input['TunerKernelParamsArgs']] = None,
                 net: Optional[pulumi.Input['TunerNetParamsArgs']] = None,
                 vm: Optional[pulumi.Input['TunerVmParamsArgs']] = None):
        if cpu_governor is not None:
            pulumi.set(__self__, "cpu_governor", cpu_governor)
        if fs is not None:
            pulumi.set(__self__, "fs", fs)
        if kernel is not None:
            pulumi.set(__self__, "kernel", kernel)
        if net is not None:
            pulumi.set(__self__, "net", net)
        if vm is not None:
            pulumi.set(__self__, "vm", vm)

    @property
    @pulumi.getter(name="cpuGovernor")
    def cpu_governor(self) -> Optional[pulumi.Input['CpuGovernor']]:
        return pulumi.get(self, "cpu_governor")

    @cpu_governor.setter
    def cpu_governor(self, value: Optional[pulumi.Input['CpuGovernor']]):
        pulumi.set(self, "cpu_governor", value)

    @property
    @pulumi.getter
    def fs(self) -> Optional[pulumi.Input['TunerFsParamsArgs']]:
        return pulumi.get(self, "fs")

    @fs.setter
    def fs(self, value: Optional[pulumi.Input['TunerFsParamsArgs']]):
        pulumi.set(self, "fs", value)

    @property
    @pulumi.getter
    def kernel(self) -> Optional[pulumi.Input['TunerKernelParamsArgs']]:
        return pulumi.get(self, "kernel")

    @kernel.setter
    def kernel(self, value: Optional[pulumi.Input['TunerKernelParamsArgs']]):
        pulumi.set(self, "kernel", value)

    @property
    @pulumi.getter
    def net(self) -> Optional[pulumi.Input['TunerNetParamsArgs']]:
        return pulumi.get(self, "net")

    @net.setter
    def net(self, value: Optional[pulumi.Input['TunerNetParamsArgs']]):
        pulumi.set(self, "net", value)

    @property
    @pulumi.getter
    def vm(self) -> Optional[pulumi.Input['TunerVmParamsArgs']]:
        return pulumi.get(self, "vm")

    @vm.setter
    def vm(self, value: Optional[pulumi.Input['TunerVmParamsArgs']]):
        pulumi.set(self, "vm", value)


if not MYPY:
    class TunerVmParamsArgsDict(TypedDict):
        vm_dirty_background_ratio: NotRequired[pulumi.Input[int]]
        vm_dirty_expire_centisecs: NotRequired[pulumi.Input[int]]
        vm_dirty_ratio: NotRequired[pulumi.Input[int]]
        vm_dirty_writeback_centisecs: NotRequired[pulumi.Input[int]]
        vm_dirtytime_expire_seconds: NotRequired[pulumi.Input[int]]
        vm_max_map_count: NotRequired[pulumi.Input[int]]
        vm_min_free_kbytes: NotRequired[pulumi.Input[int]]
        vm_stat_interval: NotRequired[pulumi.Input[int]]
        vm_swappiness: NotRequired[pulumi.Input[int]]
elif False:
    TunerVmParamsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TunerVmParamsArgs:
    def __init__(__self__, *,
                 vm_dirty_background_ratio: Optional[pulumi.Input[int]] = None,
                 vm_dirty_expire_centisecs: Optional[pulumi.Input[int]] = None,
                 vm_dirty_ratio: Optional[pulumi.Input[int]] = None,
                 vm_dirty_writeback_centisecs: Optional[pulumi.Input[int]] = None,
                 vm_dirtytime_expire_seconds: Optional[pulumi.Input[int]] = None,
                 vm_max_map_count: Optional[pulumi.Input[int]] = None,
                 vm_min_free_kbytes: Optional[pulumi.Input[int]] = None,
                 vm_stat_interval: Optional[pulumi.Input[int]] = None,
                 vm_swappiness: Optional[pulumi.Input[int]] = None):
        if vm_dirty_background_ratio is not None:
            pulumi.set(__self__, "vm_dirty_background_ratio", vm_dirty_background_ratio)
        if vm_dirty_expire_centisecs is not None:
            pulumi.set(__self__, "vm_dirty_expire_centisecs", vm_dirty_expire_centisecs)
        if vm_dirty_ratio is not None:
            pulumi.set(__self__, "vm_dirty_ratio", vm_dirty_ratio)
        if vm_dirty_writeback_centisecs is not None:
            pulumi.set(__self__, "vm_dirty_writeback_centisecs", vm_dirty_writeback_centisecs)
        if vm_dirtytime_expire_seconds is not None:
            pulumi.set(__self__, "vm_dirtytime_expire_seconds", vm_dirtytime_expire_seconds)
        if vm_max_map_count is not None:
            pulumi.set(__self__, "vm_max_map_count", vm_max_map_count)
        if vm_min_free_kbytes is not None:
            pulumi.set(__self__, "vm_min_free_kbytes", vm_min_free_kbytes)
        if vm_stat_interval is not None:
            pulumi.set(__self__, "vm_stat_interval", vm_stat_interval)
        if vm_swappiness is not None:
            pulumi.set(__self__, "vm_swappiness", vm_swappiness)

    @property
    @pulumi.getter(name="vmDirtyBackgroundRatio")
    def vm_dirty_background_ratio(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_dirty_background_ratio")

    @vm_dirty_background_ratio.setter
    def vm_dirty_background_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_dirty_background_ratio", value)

    @property
    @pulumi.getter(name="vmDirtyExpireCentisecs")
    def vm_dirty_expire_centisecs(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_dirty_expire_centisecs")

    @vm_dirty_expire_centisecs.setter
    def vm_dirty_expire_centisecs(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_dirty_expire_centisecs", value)

    @property
    @pulumi.getter(name="vmDirtyRatio")
    def vm_dirty_ratio(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_dirty_ratio")

    @vm_dirty_ratio.setter
    def vm_dirty_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_dirty_ratio", value)

    @property
    @pulumi.getter(name="vmDirtyWritebackCentisecs")
    def vm_dirty_writeback_centisecs(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_dirty_writeback_centisecs")

    @vm_dirty_writeback_centisecs.setter
    def vm_dirty_writeback_centisecs(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_dirty_writeback_centisecs", value)

    @property
    @pulumi.getter(name="vmDirtytimeExpireSeconds")
    def vm_dirtytime_expire_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_dirtytime_expire_seconds")

    @vm_dirtytime_expire_seconds.setter
    def vm_dirtytime_expire_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_dirtytime_expire_seconds", value)

    @property
    @pulumi.getter(name="vmMaxMapCount")
    def vm_max_map_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_max_map_count")

    @vm_max_map_count.setter
    def vm_max_map_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_max_map_count", value)

    @property
    @pulumi.getter(name="vmMinFreeKbytes")
    def vm_min_free_kbytes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_min_free_kbytes")

    @vm_min_free_kbytes.setter
    def vm_min_free_kbytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_min_free_kbytes", value)

    @property
    @pulumi.getter(name="vmStatInterval")
    def vm_stat_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_stat_interval")

    @vm_stat_interval.setter
    def vm_stat_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_stat_interval", value)

    @property
    @pulumi.getter(name="vmSwappiness")
    def vm_swappiness(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vm_swappiness")

    @vm_swappiness.setter
    def vm_swappiness(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vm_swappiness", value)


