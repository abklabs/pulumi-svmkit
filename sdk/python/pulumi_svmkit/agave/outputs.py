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
    'Flags',
    'KeyPairs',
    'Metrics',
]

@pulumi.output_type
class Flags(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "blockProductionMethod":
            suggest = "block_production_method"
        elif key == "dynamicPortRange":
            suggest = "dynamic_port_range"
        elif key == "fullSnapshotIntervalSlots":
            suggest = "full_snapshot_interval_slots"
        elif key == "gossipPort":
            suggest = "gossip_port"
        elif key == "limitLedgerSize":
            suggest = "limit_ledger_size"
        elif key == "noWaitForVoteToStartLeader":
            suggest = "no_wait_for_vote_to_start_leader"
        elif key == "onlyKnownRPC":
            suggest = "only_known_rpc"
        elif key == "privateRPC":
            suggest = "private_rpc"
        elif key == "rpcBindAddress":
            suggest = "rpc_bind_address"
        elif key == "rpcPort":
            suggest = "rpc_port"
        elif key == "useSnapshotArchivesAtStartup":
            suggest = "use_snapshot_archives_at_startup"
        elif key == "walRecoveryMode":
            suggest = "wal_recovery_mode"
        elif key == "allowPrivateAddr":
            suggest = "allow_private_addr"
        elif key == "entryPoint":
            suggest = "entry_point"
        elif key == "expectedGenesisHash":
            suggest = "expected_genesis_hash"
        elif key == "extraFlags":
            suggest = "extra_flags"
        elif key == "fullRpcAPI":
            suggest = "full_rpc_api"
        elif key == "gossipHost":
            suggest = "gossip_host"
        elif key == "knownValidator":
            suggest = "known_validator"
        elif key == "noVoting":
            suggest = "no_voting"
        elif key == "tvuReceiveThreads":
            suggest = "tvu_receive_threads"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Flags. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Flags.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Flags.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 block_production_method: str,
                 dynamic_port_range: str,
                 full_snapshot_interval_slots: int,
                 gossip_port: int,
                 limit_ledger_size: int,
                 no_wait_for_vote_to_start_leader: bool,
                 only_known_rpc: bool,
                 private_rpc: bool,
                 rpc_bind_address: str,
                 rpc_port: int,
                 use_snapshot_archives_at_startup: str,
                 wal_recovery_mode: str,
                 allow_private_addr: Optional[bool] = None,
                 entry_point: Optional[Sequence[str]] = None,
                 expected_genesis_hash: Optional[str] = None,
                 extra_flags: Optional[Sequence[str]] = None,
                 full_rpc_api: Optional[bool] = None,
                 gossip_host: Optional[str] = None,
                 known_validator: Optional[Sequence[str]] = None,
                 no_voting: Optional[bool] = None,
                 tvu_receive_threads: Optional[int] = None):
        pulumi.set(__self__, "block_production_method", block_production_method)
        pulumi.set(__self__, "dynamic_port_range", dynamic_port_range)
        pulumi.set(__self__, "full_snapshot_interval_slots", full_snapshot_interval_slots)
        pulumi.set(__self__, "gossip_port", gossip_port)
        pulumi.set(__self__, "limit_ledger_size", limit_ledger_size)
        pulumi.set(__self__, "no_wait_for_vote_to_start_leader", no_wait_for_vote_to_start_leader)
        pulumi.set(__self__, "only_known_rpc", only_known_rpc)
        pulumi.set(__self__, "private_rpc", private_rpc)
        pulumi.set(__self__, "rpc_bind_address", rpc_bind_address)
        pulumi.set(__self__, "rpc_port", rpc_port)
        pulumi.set(__self__, "use_snapshot_archives_at_startup", use_snapshot_archives_at_startup)
        pulumi.set(__self__, "wal_recovery_mode", wal_recovery_mode)
        if allow_private_addr is not None:
            pulumi.set(__self__, "allow_private_addr", allow_private_addr)
        if entry_point is not None:
            pulumi.set(__self__, "entry_point", entry_point)
        if expected_genesis_hash is not None:
            pulumi.set(__self__, "expected_genesis_hash", expected_genesis_hash)
        if extra_flags is not None:
            pulumi.set(__self__, "extra_flags", extra_flags)
        if full_rpc_api is not None:
            pulumi.set(__self__, "full_rpc_api", full_rpc_api)
        if gossip_host is not None:
            pulumi.set(__self__, "gossip_host", gossip_host)
        if known_validator is not None:
            pulumi.set(__self__, "known_validator", known_validator)
        if no_voting is not None:
            pulumi.set(__self__, "no_voting", no_voting)
        if tvu_receive_threads is not None:
            pulumi.set(__self__, "tvu_receive_threads", tvu_receive_threads)

    @property
    @pulumi.getter(name="blockProductionMethod")
    def block_production_method(self) -> str:
        return pulumi.get(self, "block_production_method")

    @property
    @pulumi.getter(name="dynamicPortRange")
    def dynamic_port_range(self) -> str:
        return pulumi.get(self, "dynamic_port_range")

    @property
    @pulumi.getter(name="fullSnapshotIntervalSlots")
    def full_snapshot_interval_slots(self) -> int:
        return pulumi.get(self, "full_snapshot_interval_slots")

    @property
    @pulumi.getter(name="gossipPort")
    def gossip_port(self) -> int:
        return pulumi.get(self, "gossip_port")

    @property
    @pulumi.getter(name="limitLedgerSize")
    def limit_ledger_size(self) -> int:
        return pulumi.get(self, "limit_ledger_size")

    @property
    @pulumi.getter(name="noWaitForVoteToStartLeader")
    def no_wait_for_vote_to_start_leader(self) -> bool:
        return pulumi.get(self, "no_wait_for_vote_to_start_leader")

    @property
    @pulumi.getter(name="onlyKnownRPC")
    def only_known_rpc(self) -> bool:
        return pulumi.get(self, "only_known_rpc")

    @property
    @pulumi.getter(name="privateRPC")
    def private_rpc(self) -> bool:
        return pulumi.get(self, "private_rpc")

    @property
    @pulumi.getter(name="rpcBindAddress")
    def rpc_bind_address(self) -> str:
        return pulumi.get(self, "rpc_bind_address")

    @property
    @pulumi.getter(name="rpcPort")
    def rpc_port(self) -> int:
        return pulumi.get(self, "rpc_port")

    @property
    @pulumi.getter(name="useSnapshotArchivesAtStartup")
    def use_snapshot_archives_at_startup(self) -> str:
        return pulumi.get(self, "use_snapshot_archives_at_startup")

    @property
    @pulumi.getter(name="walRecoveryMode")
    def wal_recovery_mode(self) -> str:
        return pulumi.get(self, "wal_recovery_mode")

    @property
    @pulumi.getter(name="allowPrivateAddr")
    def allow_private_addr(self) -> Optional[bool]:
        return pulumi.get(self, "allow_private_addr")

    @property
    @pulumi.getter(name="entryPoint")
    def entry_point(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "entry_point")

    @property
    @pulumi.getter(name="expectedGenesisHash")
    def expected_genesis_hash(self) -> Optional[str]:
        return pulumi.get(self, "expected_genesis_hash")

    @property
    @pulumi.getter(name="extraFlags")
    def extra_flags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "extra_flags")

    @property
    @pulumi.getter(name="fullRpcAPI")
    def full_rpc_api(self) -> Optional[bool]:
        return pulumi.get(self, "full_rpc_api")

    @property
    @pulumi.getter(name="gossipHost")
    def gossip_host(self) -> Optional[str]:
        return pulumi.get(self, "gossip_host")

    @property
    @pulumi.getter(name="knownValidator")
    def known_validator(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "known_validator")

    @property
    @pulumi.getter(name="noVoting")
    def no_voting(self) -> Optional[bool]:
        return pulumi.get(self, "no_voting")

    @property
    @pulumi.getter(name="tvuReceiveThreads")
    def tvu_receive_threads(self) -> Optional[int]:
        return pulumi.get(self, "tvu_receive_threads")


@pulumi.output_type
class KeyPairs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "voteAccount":
            suggest = "vote_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KeyPairs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KeyPairs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KeyPairs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 identity: str,
                 vote_account: str):
        pulumi.set(__self__, "identity", identity)
        pulumi.set(__self__, "vote_account", vote_account)

    @property
    @pulumi.getter
    def identity(self) -> str:
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="voteAccount")
    def vote_account(self) -> str:
        return pulumi.get(self, "vote_account")


@pulumi.output_type
class Metrics(dict):
    def __init__(__self__, *,
                 database: str,
                 password: str,
                 url: str,
                 user: str):
        pulumi.set(__self__, "database", database)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "url", url)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def database(self) -> str:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def user(self) -> str:
        return pulumi.get(self, "user")


