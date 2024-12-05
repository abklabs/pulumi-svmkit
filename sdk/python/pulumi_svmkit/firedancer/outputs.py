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
from . import outputs

__all__ = [
    'Config',
    'ConfigConsensus',
    'ConfigGossip',
    'ConfigHugeTLBFS',
    'ConfigLayout',
    'ConfigLedger',
    'ConfigLog',
    'ConfigRPC',
    'ConfigReporting',
    'ConfigSnapshots',
    'KeyPairs',
]

@pulumi.output_type
class Config(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dynamicPortRange":
            suggest = "dynamic_port_range"
        elif key == "extraConfig":
            suggest = "extra_config"
        elif key == "scratchDirectory":
            suggest = "scratch_directory"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Config. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Config.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Config.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 consensus: Optional['outputs.ConfigConsensus'] = None,
                 dynamic_port_range: Optional[str] = None,
                 extra_config: Optional[Sequence[str]] = None,
                 gossip: Optional['outputs.ConfigGossip'] = None,
                 hugetlbfs: Optional['outputs.ConfigHugeTLBFS'] = None,
                 layout: Optional['outputs.ConfigLayout'] = None,
                 ledger: Optional['outputs.ConfigLedger'] = None,
                 log: Optional['outputs.ConfigLog'] = None,
                 name: Optional[str] = None,
                 reporting: Optional['outputs.ConfigReporting'] = None,
                 rpc: Optional['outputs.ConfigRPC'] = None,
                 scratch_directory: Optional[str] = None,
                 snapshots: Optional['outputs.ConfigSnapshots'] = None,
                 user: Optional[str] = None):
        if consensus is not None:
            pulumi.set(__self__, "consensus", consensus)
        if dynamic_port_range is not None:
            pulumi.set(__self__, "dynamic_port_range", dynamic_port_range)
        if extra_config is not None:
            pulumi.set(__self__, "extra_config", extra_config)
        if gossip is not None:
            pulumi.set(__self__, "gossip", gossip)
        if hugetlbfs is not None:
            pulumi.set(__self__, "hugetlbfs", hugetlbfs)
        if layout is not None:
            pulumi.set(__self__, "layout", layout)
        if ledger is not None:
            pulumi.set(__self__, "ledger", ledger)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if reporting is not None:
            pulumi.set(__self__, "reporting", reporting)
        if rpc is not None:
            pulumi.set(__self__, "rpc", rpc)
        if scratch_directory is not None:
            pulumi.set(__self__, "scratch_directory", scratch_directory)
        if snapshots is not None:
            pulumi.set(__self__, "snapshots", snapshots)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def consensus(self) -> Optional['outputs.ConfigConsensus']:
        return pulumi.get(self, "consensus")

    @property
    @pulumi.getter(name="dynamicPortRange")
    def dynamic_port_range(self) -> Optional[str]:
        return pulumi.get(self, "dynamic_port_range")

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "extra_config")

    @property
    @pulumi.getter
    def gossip(self) -> Optional['outputs.ConfigGossip']:
        return pulumi.get(self, "gossip")

    @property
    @pulumi.getter
    def hugetlbfs(self) -> Optional['outputs.ConfigHugeTLBFS']:
        return pulumi.get(self, "hugetlbfs")

    @property
    @pulumi.getter
    def layout(self) -> Optional['outputs.ConfigLayout']:
        return pulumi.get(self, "layout")

    @property
    @pulumi.getter
    def ledger(self) -> Optional['outputs.ConfigLedger']:
        return pulumi.get(self, "ledger")

    @property
    @pulumi.getter
    def log(self) -> Optional['outputs.ConfigLog']:
        return pulumi.get(self, "log")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def reporting(self) -> Optional['outputs.ConfigReporting']:
        return pulumi.get(self, "reporting")

    @property
    @pulumi.getter
    def rpc(self) -> Optional['outputs.ConfigRPC']:
        return pulumi.get(self, "rpc")

    @property
    @pulumi.getter(name="scratchDirectory")
    def scratch_directory(self) -> Optional[str]:
        return pulumi.get(self, "scratch_directory")

    @property
    @pulumi.getter
    def snapshots(self) -> Optional['outputs.ConfigSnapshots']:
        return pulumi.get(self, "snapshots")

    @property
    @pulumi.getter
    def user(self) -> Optional[str]:
        return pulumi.get(self, "user")


@pulumi.output_type
class ConfigConsensus(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "authorizedVoterPaths":
            suggest = "authorized_voter_paths"
        elif key == "expectedBankHash":
            suggest = "expected_bank_hash"
        elif key == "expectedGenesisHash":
            suggest = "expected_genesis_hash"
        elif key == "expectedShredVersion":
            suggest = "expected_shred_version"
        elif key == "genesisFetch":
            suggest = "genesis_fetch"
        elif key == "hardForkAtSlots":
            suggest = "hard_fork_at_slots"
        elif key == "identityPath":
            suggest = "identity_path"
        elif key == "knownValidators":
            suggest = "known_validators"
        elif key == "osNetworkLimitsTest":
            suggest = "os_network_limits_test"
        elif key == "pohSpeedTest":
            suggest = "poh_speed_test"
        elif key == "snapshotFetch":
            suggest = "snapshot_fetch"
        elif key == "voteAccountPath":
            suggest = "vote_account_path"
        elif key == "waitForSupermajorityAtSlot":
            suggest = "wait_for_supermajority_at_slot"
        elif key == "waitForVoteToStartLeader":
            suggest = "wait_for_vote_to_start_leader"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigConsensus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigConsensus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigConsensus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 authorized_voter_paths: Optional[Sequence[str]] = None,
                 expected_bank_hash: Optional[str] = None,
                 expected_genesis_hash: Optional[str] = None,
                 expected_shred_version: Optional[int] = None,
                 genesis_fetch: Optional[bool] = None,
                 hard_fork_at_slots: Optional[Sequence[str]] = None,
                 identity_path: Optional[str] = None,
                 known_validators: Optional[Sequence[str]] = None,
                 os_network_limits_test: Optional[bool] = None,
                 poh_speed_test: Optional[bool] = None,
                 snapshot_fetch: Optional[bool] = None,
                 vote_account_path: Optional[str] = None,
                 wait_for_supermajority_at_slot: Optional[int] = None,
                 wait_for_vote_to_start_leader: Optional[bool] = None):
        if authorized_voter_paths is not None:
            pulumi.set(__self__, "authorized_voter_paths", authorized_voter_paths)
        if expected_bank_hash is not None:
            pulumi.set(__self__, "expected_bank_hash", expected_bank_hash)
        if expected_genesis_hash is not None:
            pulumi.set(__self__, "expected_genesis_hash", expected_genesis_hash)
        if expected_shred_version is not None:
            pulumi.set(__self__, "expected_shred_version", expected_shred_version)
        if genesis_fetch is not None:
            pulumi.set(__self__, "genesis_fetch", genesis_fetch)
        if hard_fork_at_slots is not None:
            pulumi.set(__self__, "hard_fork_at_slots", hard_fork_at_slots)
        if identity_path is not None:
            pulumi.set(__self__, "identity_path", identity_path)
        if known_validators is not None:
            pulumi.set(__self__, "known_validators", known_validators)
        if os_network_limits_test is not None:
            pulumi.set(__self__, "os_network_limits_test", os_network_limits_test)
        if poh_speed_test is not None:
            pulumi.set(__self__, "poh_speed_test", poh_speed_test)
        if snapshot_fetch is not None:
            pulumi.set(__self__, "snapshot_fetch", snapshot_fetch)
        if vote_account_path is not None:
            pulumi.set(__self__, "vote_account_path", vote_account_path)
        if wait_for_supermajority_at_slot is not None:
            pulumi.set(__self__, "wait_for_supermajority_at_slot", wait_for_supermajority_at_slot)
        if wait_for_vote_to_start_leader is not None:
            pulumi.set(__self__, "wait_for_vote_to_start_leader", wait_for_vote_to_start_leader)

    @property
    @pulumi.getter(name="authorizedVoterPaths")
    def authorized_voter_paths(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "authorized_voter_paths")

    @property
    @pulumi.getter(name="expectedBankHash")
    def expected_bank_hash(self) -> Optional[str]:
        return pulumi.get(self, "expected_bank_hash")

    @property
    @pulumi.getter(name="expectedGenesisHash")
    def expected_genesis_hash(self) -> Optional[str]:
        return pulumi.get(self, "expected_genesis_hash")

    @property
    @pulumi.getter(name="expectedShredVersion")
    def expected_shred_version(self) -> Optional[int]:
        return pulumi.get(self, "expected_shred_version")

    @property
    @pulumi.getter(name="genesisFetch")
    def genesis_fetch(self) -> Optional[bool]:
        return pulumi.get(self, "genesis_fetch")

    @property
    @pulumi.getter(name="hardForkAtSlots")
    def hard_fork_at_slots(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "hard_fork_at_slots")

    @property
    @pulumi.getter(name="identityPath")
    def identity_path(self) -> Optional[str]:
        return pulumi.get(self, "identity_path")

    @property
    @pulumi.getter(name="knownValidators")
    def known_validators(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "known_validators")

    @property
    @pulumi.getter(name="osNetworkLimitsTest")
    def os_network_limits_test(self) -> Optional[bool]:
        return pulumi.get(self, "os_network_limits_test")

    @property
    @pulumi.getter(name="pohSpeedTest")
    def poh_speed_test(self) -> Optional[bool]:
        return pulumi.get(self, "poh_speed_test")

    @property
    @pulumi.getter(name="snapshotFetch")
    def snapshot_fetch(self) -> Optional[bool]:
        return pulumi.get(self, "snapshot_fetch")

    @property
    @pulumi.getter(name="voteAccountPath")
    def vote_account_path(self) -> Optional[str]:
        return pulumi.get(self, "vote_account_path")

    @property
    @pulumi.getter(name="waitForSupermajorityAtSlot")
    def wait_for_supermajority_at_slot(self) -> Optional[int]:
        return pulumi.get(self, "wait_for_supermajority_at_slot")

    @property
    @pulumi.getter(name="waitForVoteToStartLeader")
    def wait_for_vote_to_start_leader(self) -> Optional[bool]:
        return pulumi.get(self, "wait_for_vote_to_start_leader")


@pulumi.output_type
class ConfigGossip(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "portCheck":
            suggest = "port_check"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigGossip. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigGossip.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigGossip.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 entrypoints: Optional[Sequence[str]] = None,
                 host: Optional[str] = None,
                 port: Optional[int] = None,
                 port_check: Optional[bool] = None):
        if entrypoints is not None:
            pulumi.set(__self__, "entrypoints", entrypoints)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if port_check is not None:
            pulumi.set(__self__, "port_check", port_check)

    @property
    @pulumi.getter
    def entrypoints(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "entrypoints")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="portCheck")
    def port_check(self) -> Optional[bool]:
        return pulumi.get(self, "port_check")


@pulumi.output_type
class ConfigHugeTLBFS(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "mountPath":
            suggest = "mount_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigHugeTLBFS. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigHugeTLBFS.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigHugeTLBFS.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mount_path: Optional[str] = None):
        if mount_path is not None:
            pulumi.set(__self__, "mount_path", mount_path)

    @property
    @pulumi.getter(name="mountPath")
    def mount_path(self) -> Optional[str]:
        return pulumi.get(self, "mount_path")


@pulumi.output_type
class ConfigLayout(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "agaveAffinity":
            suggest = "agave_affinity"
        elif key == "bankTileCount":
            suggest = "bank_tile_count"
        elif key == "netTileCount":
            suggest = "net_tile_count"
        elif key == "quicTileCount":
            suggest = "quic_tile_count"
        elif key == "resolvTileCount":
            suggest = "resolv_tile_count"
        elif key == "shredTileCount":
            suggest = "shred_tile_count"
        elif key == "verifyTileCount":
            suggest = "verify_tile_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigLayout. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigLayout.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigLayout.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 affinity: Optional[str] = None,
                 agave_affinity: Optional[str] = None,
                 bank_tile_count: Optional[int] = None,
                 net_tile_count: Optional[int] = None,
                 quic_tile_count: Optional[int] = None,
                 resolv_tile_count: Optional[int] = None,
                 shred_tile_count: Optional[int] = None,
                 verify_tile_count: Optional[int] = None):
        if affinity is not None:
            pulumi.set(__self__, "affinity", affinity)
        if agave_affinity is not None:
            pulumi.set(__self__, "agave_affinity", agave_affinity)
        if bank_tile_count is not None:
            pulumi.set(__self__, "bank_tile_count", bank_tile_count)
        if net_tile_count is not None:
            pulumi.set(__self__, "net_tile_count", net_tile_count)
        if quic_tile_count is not None:
            pulumi.set(__self__, "quic_tile_count", quic_tile_count)
        if resolv_tile_count is not None:
            pulumi.set(__self__, "resolv_tile_count", resolv_tile_count)
        if shred_tile_count is not None:
            pulumi.set(__self__, "shred_tile_count", shred_tile_count)
        if verify_tile_count is not None:
            pulumi.set(__self__, "verify_tile_count", verify_tile_count)

    @property
    @pulumi.getter
    def affinity(self) -> Optional[str]:
        return pulumi.get(self, "affinity")

    @property
    @pulumi.getter(name="agaveAffinity")
    def agave_affinity(self) -> Optional[str]:
        return pulumi.get(self, "agave_affinity")

    @property
    @pulumi.getter(name="bankTileCount")
    def bank_tile_count(self) -> Optional[int]:
        return pulumi.get(self, "bank_tile_count")

    @property
    @pulumi.getter(name="netTileCount")
    def net_tile_count(self) -> Optional[int]:
        return pulumi.get(self, "net_tile_count")

    @property
    @pulumi.getter(name="quicTileCount")
    def quic_tile_count(self) -> Optional[int]:
        return pulumi.get(self, "quic_tile_count")

    @property
    @pulumi.getter(name="resolvTileCount")
    def resolv_tile_count(self) -> Optional[int]:
        return pulumi.get(self, "resolv_tile_count")

    @property
    @pulumi.getter(name="shredTileCount")
    def shred_tile_count(self) -> Optional[int]:
        return pulumi.get(self, "shred_tile_count")

    @property
    @pulumi.getter(name="verifyTileCount")
    def verify_tile_count(self) -> Optional[int]:
        return pulumi.get(self, "verify_tile_count")


@pulumi.output_type
class ConfigLedger(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountIndexExcludeKeys":
            suggest = "account_index_exclude_keys"
        elif key == "accountIndexIncludeKeys":
            suggest = "account_index_include_keys"
        elif key == "accountIndexes":
            suggest = "account_indexes"
        elif key == "accountsPath":
            suggest = "accounts_path"
        elif key == "limitSize":
            suggest = "limit_size"
        elif key == "requireTower":
            suggest = "require_tower"
        elif key == "snapshotArchiveFormat":
            suggest = "snapshot_archive_format"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigLedger. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigLedger.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigLedger.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_index_exclude_keys: Optional[Sequence[str]] = None,
                 account_index_include_keys: Optional[Sequence[str]] = None,
                 account_indexes: Optional[Sequence[str]] = None,
                 accounts_path: Optional[str] = None,
                 limit_size: Optional[int] = None,
                 path: Optional[str] = None,
                 require_tower: Optional[bool] = None,
                 snapshot_archive_format: Optional[str] = None):
        if account_index_exclude_keys is not None:
            pulumi.set(__self__, "account_index_exclude_keys", account_index_exclude_keys)
        if account_index_include_keys is not None:
            pulumi.set(__self__, "account_index_include_keys", account_index_include_keys)
        if account_indexes is not None:
            pulumi.set(__self__, "account_indexes", account_indexes)
        if accounts_path is not None:
            pulumi.set(__self__, "accounts_path", accounts_path)
        if limit_size is not None:
            pulumi.set(__self__, "limit_size", limit_size)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if require_tower is not None:
            pulumi.set(__self__, "require_tower", require_tower)
        if snapshot_archive_format is not None:
            pulumi.set(__self__, "snapshot_archive_format", snapshot_archive_format)

    @property
    @pulumi.getter(name="accountIndexExcludeKeys")
    def account_index_exclude_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "account_index_exclude_keys")

    @property
    @pulumi.getter(name="accountIndexIncludeKeys")
    def account_index_include_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "account_index_include_keys")

    @property
    @pulumi.getter(name="accountIndexes")
    def account_indexes(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "account_indexes")

    @property
    @pulumi.getter(name="accountsPath")
    def accounts_path(self) -> Optional[str]:
        return pulumi.get(self, "accounts_path")

    @property
    @pulumi.getter(name="limitSize")
    def limit_size(self) -> Optional[int]:
        return pulumi.get(self, "limit_size")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="requireTower")
    def require_tower(self) -> Optional[bool]:
        return pulumi.get(self, "require_tower")

    @property
    @pulumi.getter(name="snapshotArchiveFormat")
    def snapshot_archive_format(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_archive_format")


@pulumi.output_type
class ConfigLog(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "levelFlush":
            suggest = "level_flush"
        elif key == "levelLogfile":
            suggest = "level_logfile"
        elif key == "levelStderr":
            suggest = "level_stderr"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigLog. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigLog.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigLog.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 colorize: Optional[str] = None,
                 level_flush: Optional[str] = None,
                 level_logfile: Optional[str] = None,
                 level_stderr: Optional[str] = None,
                 path: Optional[str] = None):
        if colorize is not None:
            pulumi.set(__self__, "colorize", colorize)
        if level_flush is not None:
            pulumi.set(__self__, "level_flush", level_flush)
        if level_logfile is not None:
            pulumi.set(__self__, "level_logfile", level_logfile)
        if level_stderr is not None:
            pulumi.set(__self__, "level_stderr", level_stderr)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def colorize(self) -> Optional[str]:
        return pulumi.get(self, "colorize")

    @property
    @pulumi.getter(name="levelFlush")
    def level_flush(self) -> Optional[str]:
        return pulumi.get(self, "level_flush")

    @property
    @pulumi.getter(name="levelLogfile")
    def level_logfile(self) -> Optional[str]:
        return pulumi.get(self, "level_logfile")

    @property
    @pulumi.getter(name="levelStderr")
    def level_stderr(self) -> Optional[str]:
        return pulumi.get(self, "level_stderr")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")


@pulumi.output_type
class ConfigRPC(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bigtableLedgerStorage":
            suggest = "bigtable_ledger_storage"
        elif key == "extendedTxMetadataStorage":
            suggest = "extended_tx_metadata_storage"
        elif key == "fullApi":
            suggest = "full_api"
        elif key == "onlyKnown":
            suggest = "only_known"
        elif key == "pubsubEnableBlockSubscription":
            suggest = "pubsub_enable_block_subscription"
        elif key == "pubsubEnableVoteSubscription":
            suggest = "pubsub_enable_vote_subscription"
        elif key == "transactionHistory":
            suggest = "transaction_history"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigRPC. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigRPC.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigRPC.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bigtable_ledger_storage: Optional[bool] = None,
                 extended_tx_metadata_storage: Optional[bool] = None,
                 full_api: Optional[bool] = None,
                 only_known: Optional[bool] = None,
                 port: Optional[int] = None,
                 private: Optional[bool] = None,
                 pubsub_enable_block_subscription: Optional[bool] = None,
                 pubsub_enable_vote_subscription: Optional[bool] = None,
                 transaction_history: Optional[bool] = None):
        if bigtable_ledger_storage is not None:
            pulumi.set(__self__, "bigtable_ledger_storage", bigtable_ledger_storage)
        if extended_tx_metadata_storage is not None:
            pulumi.set(__self__, "extended_tx_metadata_storage", extended_tx_metadata_storage)
        if full_api is not None:
            pulumi.set(__self__, "full_api", full_api)
        if only_known is not None:
            pulumi.set(__self__, "only_known", only_known)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if private is not None:
            pulumi.set(__self__, "private", private)
        if pubsub_enable_block_subscription is not None:
            pulumi.set(__self__, "pubsub_enable_block_subscription", pubsub_enable_block_subscription)
        if pubsub_enable_vote_subscription is not None:
            pulumi.set(__self__, "pubsub_enable_vote_subscription", pubsub_enable_vote_subscription)
        if transaction_history is not None:
            pulumi.set(__self__, "transaction_history", transaction_history)

    @property
    @pulumi.getter(name="bigtableLedgerStorage")
    def bigtable_ledger_storage(self) -> Optional[bool]:
        return pulumi.get(self, "bigtable_ledger_storage")

    @property
    @pulumi.getter(name="extendedTxMetadataStorage")
    def extended_tx_metadata_storage(self) -> Optional[bool]:
        return pulumi.get(self, "extended_tx_metadata_storage")

    @property
    @pulumi.getter(name="fullApi")
    def full_api(self) -> Optional[bool]:
        return pulumi.get(self, "full_api")

    @property
    @pulumi.getter(name="onlyKnown")
    def only_known(self) -> Optional[bool]:
        return pulumi.get(self, "only_known")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def private(self) -> Optional[bool]:
        return pulumi.get(self, "private")

    @property
    @pulumi.getter(name="pubsubEnableBlockSubscription")
    def pubsub_enable_block_subscription(self) -> Optional[bool]:
        return pulumi.get(self, "pubsub_enable_block_subscription")

    @property
    @pulumi.getter(name="pubsubEnableVoteSubscription")
    def pubsub_enable_vote_subscription(self) -> Optional[bool]:
        return pulumi.get(self, "pubsub_enable_vote_subscription")

    @property
    @pulumi.getter(name="transactionHistory")
    def transaction_history(self) -> Optional[bool]:
        return pulumi.get(self, "transaction_history")


@pulumi.output_type
class ConfigReporting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "solanaMetricsConfig":
            suggest = "solana_metrics_config"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigReporting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigReporting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigReporting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 solana_metrics_config: Optional[str] = None):
        if solana_metrics_config is not None:
            pulumi.set(__self__, "solana_metrics_config", solana_metrics_config)

    @property
    @pulumi.getter(name="solanaMetricsConfig")
    def solana_metrics_config(self) -> Optional[str]:
        return pulumi.get(self, "solana_metrics_config")


@pulumi.output_type
class ConfigSnapshots(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fullSnapshotIntervalSlots":
            suggest = "full_snapshot_interval_slots"
        elif key == "incrementalPath":
            suggest = "incremental_path"
        elif key == "incrementalSnapshotIntervalSlots":
            suggest = "incremental_snapshot_interval_slots"
        elif key == "incrementalSnapshots":
            suggest = "incremental_snapshots"
        elif key == "maximumFullSnapshotsToRetain":
            suggest = "maximum_full_snapshots_to_retain"
        elif key == "maximumIncrementalSnapshotsToRetain":
            suggest = "maximum_incremental_snapshots_to_retain"
        elif key == "minimumSnapshotDownloadSpeed":
            suggest = "minimum_snapshot_download_speed"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigSnapshots. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigSnapshots.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigSnapshots.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 full_snapshot_interval_slots: Optional[int] = None,
                 incremental_path: Optional[str] = None,
                 incremental_snapshot_interval_slots: Optional[int] = None,
                 incremental_snapshots: Optional[bool] = None,
                 maximum_full_snapshots_to_retain: Optional[int] = None,
                 maximum_incremental_snapshots_to_retain: Optional[int] = None,
                 minimum_snapshot_download_speed: Optional[int] = None,
                 path: Optional[str] = None):
        if full_snapshot_interval_slots is not None:
            pulumi.set(__self__, "full_snapshot_interval_slots", full_snapshot_interval_slots)
        if incremental_path is not None:
            pulumi.set(__self__, "incremental_path", incremental_path)
        if incremental_snapshot_interval_slots is not None:
            pulumi.set(__self__, "incremental_snapshot_interval_slots", incremental_snapshot_interval_slots)
        if incremental_snapshots is not None:
            pulumi.set(__self__, "incremental_snapshots", incremental_snapshots)
        if maximum_full_snapshots_to_retain is not None:
            pulumi.set(__self__, "maximum_full_snapshots_to_retain", maximum_full_snapshots_to_retain)
        if maximum_incremental_snapshots_to_retain is not None:
            pulumi.set(__self__, "maximum_incremental_snapshots_to_retain", maximum_incremental_snapshots_to_retain)
        if minimum_snapshot_download_speed is not None:
            pulumi.set(__self__, "minimum_snapshot_download_speed", minimum_snapshot_download_speed)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="fullSnapshotIntervalSlots")
    def full_snapshot_interval_slots(self) -> Optional[int]:
        return pulumi.get(self, "full_snapshot_interval_slots")

    @property
    @pulumi.getter(name="incrementalPath")
    def incremental_path(self) -> Optional[str]:
        return pulumi.get(self, "incremental_path")

    @property
    @pulumi.getter(name="incrementalSnapshotIntervalSlots")
    def incremental_snapshot_interval_slots(self) -> Optional[int]:
        return pulumi.get(self, "incremental_snapshot_interval_slots")

    @property
    @pulumi.getter(name="incrementalSnapshots")
    def incremental_snapshots(self) -> Optional[bool]:
        return pulumi.get(self, "incremental_snapshots")

    @property
    @pulumi.getter(name="maximumFullSnapshotsToRetain")
    def maximum_full_snapshots_to_retain(self) -> Optional[int]:
        return pulumi.get(self, "maximum_full_snapshots_to_retain")

    @property
    @pulumi.getter(name="maximumIncrementalSnapshotsToRetain")
    def maximum_incremental_snapshots_to_retain(self) -> Optional[int]:
        return pulumi.get(self, "maximum_incremental_snapshots_to_retain")

    @property
    @pulumi.getter(name="minimumSnapshotDownloadSpeed")
    def minimum_snapshot_download_speed(self) -> Optional[int]:
        return pulumi.get(self, "minimum_snapshot_download_speed")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")


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


