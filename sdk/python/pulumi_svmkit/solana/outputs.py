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
    'Environment',
    'ExplorerFlags',
    'FaucetFlags',
    'GenesisFlags',
    'PrimorialEntry',
    'StakeAccountKeyPairs',
    'TxnOptions',
    'ValidatorInfo',
    'VoteAccountKeyPairs',
]

@pulumi.output_type
class Environment(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "rpcURL":
            suggest = "rpc_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Environment. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Environment.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Environment.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 rpc_url: str):
        pulumi.set(__self__, "rpc_url", rpc_url)

    @property
    @pulumi.getter(name="rpcURL")
    def rpc_url(self) -> str:
        return pulumi.get(self, "rpc_url")


@pulumi.output_type
class ExplorerFlags(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keepAliveTimeout":
            suggest = "keep_alive_timeout"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExplorerFlags. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExplorerFlags.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExplorerFlags.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 hostname: Optional[str] = None,
                 keep_alive_timeout: Optional[int] = None,
                 port: Optional[int] = None):
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if keep_alive_timeout is not None:
            pulumi.set(__self__, "keep_alive_timeout", keep_alive_timeout)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="keepAliveTimeout")
    def keep_alive_timeout(self) -> Optional[int]:
        return pulumi.get(self, "keep_alive_timeout")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")


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


@pulumi.output_type
class GenesisFlags(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "identityPubkey":
            suggest = "identity_pubkey"
        elif key == "ledgerPath":
            suggest = "ledger_path"
        elif key == "stakePubkey":
            suggest = "stake_pubkey"
        elif key == "votePubkey":
            suggest = "vote_pubkey"
        elif key == "bootstrapStakeAuthorizedPubkey":
            suggest = "bootstrap_stake_authorized_pubkey"
        elif key == "bootstrapValidatorLamports":
            suggest = "bootstrap_validator_lamports"
        elif key == "bootstrapValidatorStakeLamports":
            suggest = "bootstrap_validator_stake_lamports"
        elif key == "clusterType":
            suggest = "cluster_type"
        elif key == "creationTime":
            suggest = "creation_time"
        elif key == "deactivateFeatures":
            suggest = "deactivate_features"
        elif key == "enableWarmupEpochs":
            suggest = "enable_warmup_epochs"
        elif key == "extraFlags":
            suggest = "extra_flags"
        elif key == "faucetLamports":
            suggest = "faucet_lamports"
        elif key == "faucetPubkey":
            suggest = "faucet_pubkey"
        elif key == "feeBurnPercentage":
            suggest = "fee_burn_percentage"
        elif key == "hashesPerTick":
            suggest = "hashes_per_tick"
        elif key == "lamportsPerByteYear":
            suggest = "lamports_per_byte_year"
        elif key == "maxGenesisArchiveUnpackedSize":
            suggest = "max_genesis_archive_unpacked_size"
        elif key == "rentBurnPercentage":
            suggest = "rent_burn_percentage"
        elif key == "rentExemptionThreshold":
            suggest = "rent_exemption_threshold"
        elif key == "slotPerEpoch":
            suggest = "slot_per_epoch"
        elif key == "targetLamportsPerSignature":
            suggest = "target_lamports_per_signature"
        elif key == "targetSignaturesPerSlot":
            suggest = "target_signatures_per_slot"
        elif key == "targetTickDuration":
            suggest = "target_tick_duration"
        elif key == "ticksPerSlot":
            suggest = "ticks_per_slot"
        elif key == "voteCommissionPercentage":
            suggest = "vote_commission_percentage"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GenesisFlags. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GenesisFlags.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GenesisFlags.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 identity_pubkey: str,
                 ledger_path: str,
                 stake_pubkey: str,
                 vote_pubkey: str,
                 bootstrap_stake_authorized_pubkey: Optional[str] = None,
                 bootstrap_validator_lamports: Optional[int] = None,
                 bootstrap_validator_stake_lamports: Optional[int] = None,
                 cluster_type: Optional[str] = None,
                 creation_time: Optional[str] = None,
                 deactivate_features: Optional[Sequence[str]] = None,
                 enable_warmup_epochs: Optional[bool] = None,
                 extra_flags: Optional[Sequence[str]] = None,
                 faucet_lamports: Optional[int] = None,
                 faucet_pubkey: Optional[str] = None,
                 fee_burn_percentage: Optional[int] = None,
                 hashes_per_tick: Optional[str] = None,
                 inflation: Optional[str] = None,
                 lamports_per_byte_year: Optional[int] = None,
                 max_genesis_archive_unpacked_size: Optional[int] = None,
                 rent_burn_percentage: Optional[int] = None,
                 rent_exemption_threshold: Optional[int] = None,
                 slot_per_epoch: Optional[int] = None,
                 target_lamports_per_signature: Optional[int] = None,
                 target_signatures_per_slot: Optional[int] = None,
                 target_tick_duration: Optional[int] = None,
                 ticks_per_slot: Optional[int] = None,
                 url: Optional[str] = None,
                 vote_commission_percentage: Optional[int] = None):
        pulumi.set(__self__, "identity_pubkey", identity_pubkey)
        pulumi.set(__self__, "ledger_path", ledger_path)
        pulumi.set(__self__, "stake_pubkey", stake_pubkey)
        pulumi.set(__self__, "vote_pubkey", vote_pubkey)
        if bootstrap_stake_authorized_pubkey is not None:
            pulumi.set(__self__, "bootstrap_stake_authorized_pubkey", bootstrap_stake_authorized_pubkey)
        if bootstrap_validator_lamports is not None:
            pulumi.set(__self__, "bootstrap_validator_lamports", bootstrap_validator_lamports)
        if bootstrap_validator_stake_lamports is not None:
            pulumi.set(__self__, "bootstrap_validator_stake_lamports", bootstrap_validator_stake_lamports)
        if cluster_type is not None:
            pulumi.set(__self__, "cluster_type", cluster_type)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if deactivate_features is not None:
            pulumi.set(__self__, "deactivate_features", deactivate_features)
        if enable_warmup_epochs is not None:
            pulumi.set(__self__, "enable_warmup_epochs", enable_warmup_epochs)
        if extra_flags is not None:
            pulumi.set(__self__, "extra_flags", extra_flags)
        if faucet_lamports is not None:
            pulumi.set(__self__, "faucet_lamports", faucet_lamports)
        if faucet_pubkey is not None:
            pulumi.set(__self__, "faucet_pubkey", faucet_pubkey)
        if fee_burn_percentage is not None:
            pulumi.set(__self__, "fee_burn_percentage", fee_burn_percentage)
        if hashes_per_tick is not None:
            pulumi.set(__self__, "hashes_per_tick", hashes_per_tick)
        if inflation is not None:
            pulumi.set(__self__, "inflation", inflation)
        if lamports_per_byte_year is not None:
            pulumi.set(__self__, "lamports_per_byte_year", lamports_per_byte_year)
        if max_genesis_archive_unpacked_size is not None:
            pulumi.set(__self__, "max_genesis_archive_unpacked_size", max_genesis_archive_unpacked_size)
        if rent_burn_percentage is not None:
            pulumi.set(__self__, "rent_burn_percentage", rent_burn_percentage)
        if rent_exemption_threshold is not None:
            pulumi.set(__self__, "rent_exemption_threshold", rent_exemption_threshold)
        if slot_per_epoch is not None:
            pulumi.set(__self__, "slot_per_epoch", slot_per_epoch)
        if target_lamports_per_signature is not None:
            pulumi.set(__self__, "target_lamports_per_signature", target_lamports_per_signature)
        if target_signatures_per_slot is not None:
            pulumi.set(__self__, "target_signatures_per_slot", target_signatures_per_slot)
        if target_tick_duration is not None:
            pulumi.set(__self__, "target_tick_duration", target_tick_duration)
        if ticks_per_slot is not None:
            pulumi.set(__self__, "ticks_per_slot", ticks_per_slot)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vote_commission_percentage is not None:
            pulumi.set(__self__, "vote_commission_percentage", vote_commission_percentage)

    @property
    @pulumi.getter(name="identityPubkey")
    def identity_pubkey(self) -> str:
        return pulumi.get(self, "identity_pubkey")

    @property
    @pulumi.getter(name="ledgerPath")
    def ledger_path(self) -> str:
        return pulumi.get(self, "ledger_path")

    @property
    @pulumi.getter(name="stakePubkey")
    def stake_pubkey(self) -> str:
        return pulumi.get(self, "stake_pubkey")

    @property
    @pulumi.getter(name="votePubkey")
    def vote_pubkey(self) -> str:
        return pulumi.get(self, "vote_pubkey")

    @property
    @pulumi.getter(name="bootstrapStakeAuthorizedPubkey")
    def bootstrap_stake_authorized_pubkey(self) -> Optional[str]:
        return pulumi.get(self, "bootstrap_stake_authorized_pubkey")

    @property
    @pulumi.getter(name="bootstrapValidatorLamports")
    def bootstrap_validator_lamports(self) -> Optional[int]:
        return pulumi.get(self, "bootstrap_validator_lamports")

    @property
    @pulumi.getter(name="bootstrapValidatorStakeLamports")
    def bootstrap_validator_stake_lamports(self) -> Optional[int]:
        return pulumi.get(self, "bootstrap_validator_stake_lamports")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[str]:
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deactivateFeatures")
    def deactivate_features(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "deactivate_features")

    @property
    @pulumi.getter(name="enableWarmupEpochs")
    def enable_warmup_epochs(self) -> Optional[bool]:
        return pulumi.get(self, "enable_warmup_epochs")

    @property
    @pulumi.getter(name="extraFlags")
    def extra_flags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "extra_flags")

    @property
    @pulumi.getter(name="faucetLamports")
    def faucet_lamports(self) -> Optional[int]:
        return pulumi.get(self, "faucet_lamports")

    @property
    @pulumi.getter(name="faucetPubkey")
    def faucet_pubkey(self) -> Optional[str]:
        return pulumi.get(self, "faucet_pubkey")

    @property
    @pulumi.getter(name="feeBurnPercentage")
    def fee_burn_percentage(self) -> Optional[int]:
        return pulumi.get(self, "fee_burn_percentage")

    @property
    @pulumi.getter(name="hashesPerTick")
    def hashes_per_tick(self) -> Optional[str]:
        return pulumi.get(self, "hashes_per_tick")

    @property
    @pulumi.getter
    def inflation(self) -> Optional[str]:
        return pulumi.get(self, "inflation")

    @property
    @pulumi.getter(name="lamportsPerByteYear")
    def lamports_per_byte_year(self) -> Optional[int]:
        return pulumi.get(self, "lamports_per_byte_year")

    @property
    @pulumi.getter(name="maxGenesisArchiveUnpackedSize")
    def max_genesis_archive_unpacked_size(self) -> Optional[int]:
        return pulumi.get(self, "max_genesis_archive_unpacked_size")

    @property
    @pulumi.getter(name="rentBurnPercentage")
    def rent_burn_percentage(self) -> Optional[int]:
        return pulumi.get(self, "rent_burn_percentage")

    @property
    @pulumi.getter(name="rentExemptionThreshold")
    def rent_exemption_threshold(self) -> Optional[int]:
        return pulumi.get(self, "rent_exemption_threshold")

    @property
    @pulumi.getter(name="slotPerEpoch")
    def slot_per_epoch(self) -> Optional[int]:
        return pulumi.get(self, "slot_per_epoch")

    @property
    @pulumi.getter(name="targetLamportsPerSignature")
    def target_lamports_per_signature(self) -> Optional[int]:
        return pulumi.get(self, "target_lamports_per_signature")

    @property
    @pulumi.getter(name="targetSignaturesPerSlot")
    def target_signatures_per_slot(self) -> Optional[int]:
        return pulumi.get(self, "target_signatures_per_slot")

    @property
    @pulumi.getter(name="targetTickDuration")
    def target_tick_duration(self) -> Optional[int]:
        return pulumi.get(self, "target_tick_duration")

    @property
    @pulumi.getter(name="ticksPerSlot")
    def ticks_per_slot(self) -> Optional[int]:
        return pulumi.get(self, "ticks_per_slot")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="voteCommissionPercentage")
    def vote_commission_percentage(self) -> Optional[int]:
        return pulumi.get(self, "vote_commission_percentage")


@pulumi.output_type
class PrimorialEntry(dict):
    def __init__(__self__, *,
                 lamports: str,
                 pubkey: str):
        pulumi.set(__self__, "lamports", lamports)
        pulumi.set(__self__, "pubkey", pubkey)

    @property
    @pulumi.getter
    def lamports(self) -> str:
        return pulumi.get(self, "lamports")

    @property
    @pulumi.getter
    def pubkey(self) -> str:
        return pulumi.get(self, "pubkey")


@pulumi.output_type
class StakeAccountKeyPairs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "stakeAccount":
            suggest = "stake_account"
        elif key == "voteAccount":
            suggest = "vote_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StakeAccountKeyPairs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StakeAccountKeyPairs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StakeAccountKeyPairs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 stake_account: str,
                 vote_account: str):
        pulumi.set(__self__, "stake_account", stake_account)
        pulumi.set(__self__, "vote_account", vote_account)

    @property
    @pulumi.getter(name="stakeAccount")
    def stake_account(self) -> str:
        return pulumi.get(self, "stake_account")

    @property
    @pulumi.getter(name="voteAccount")
    def vote_account(self) -> str:
        return pulumi.get(self, "vote_account")


@pulumi.output_type
class TxnOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "blockHash":
            suggest = "block_hash"
        elif key == "feePayer":
            suggest = "fee_payer"
        elif key == "from":
            suggest = "from_"
        elif key == "keyPair":
            suggest = "key_pair"
        elif key == "nonceAuthority":
            suggest = "nonce_authority"
        elif key == "withComputeUnitPrice":
            suggest = "with_compute_unit_price"
        elif key == "withMemo":
            suggest = "with_memo"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TxnOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TxnOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TxnOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 block_hash: Optional[str] = None,
                 commitment: Optional[str] = None,
                 fee_payer: Optional[str] = None,
                 from_: Optional[str] = None,
                 key_pair: Optional[str] = None,
                 nonce: Optional[str] = None,
                 nonce_authority: Optional[str] = None,
                 signer: Optional[Sequence[str]] = None,
                 url: Optional[str] = None,
                 with_compute_unit_price: Optional[float] = None,
                 with_memo: Optional[str] = None,
                 ws: Optional[str] = None):
        if block_hash is not None:
            pulumi.set(__self__, "block_hash", block_hash)
        if commitment is not None:
            pulumi.set(__self__, "commitment", commitment)
        if fee_payer is not None:
            pulumi.set(__self__, "fee_payer", fee_payer)
        if from_ is not None:
            pulumi.set(__self__, "from_", from_)
        if key_pair is not None:
            pulumi.set(__self__, "key_pair", key_pair)
        if nonce is not None:
            pulumi.set(__self__, "nonce", nonce)
        if nonce_authority is not None:
            pulumi.set(__self__, "nonce_authority", nonce_authority)
        if signer is not None:
            pulumi.set(__self__, "signer", signer)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if with_compute_unit_price is not None:
            pulumi.set(__self__, "with_compute_unit_price", with_compute_unit_price)
        if with_memo is not None:
            pulumi.set(__self__, "with_memo", with_memo)
        if ws is not None:
            pulumi.set(__self__, "ws", ws)

    @property
    @pulumi.getter(name="blockHash")
    def block_hash(self) -> Optional[str]:
        return pulumi.get(self, "block_hash")

    @property
    @pulumi.getter
    def commitment(self) -> Optional[str]:
        return pulumi.get(self, "commitment")

    @property
    @pulumi.getter(name="feePayer")
    def fee_payer(self) -> Optional[str]:
        return pulumi.get(self, "fee_payer")

    @property
    @pulumi.getter(name="from")
    def from_(self) -> Optional[str]:
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> Optional[str]:
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def nonce(self) -> Optional[str]:
        return pulumi.get(self, "nonce")

    @property
    @pulumi.getter(name="nonceAuthority")
    def nonce_authority(self) -> Optional[str]:
        return pulumi.get(self, "nonce_authority")

    @property
    @pulumi.getter
    def signer(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "signer")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="withComputeUnitPrice")
    def with_compute_unit_price(self) -> Optional[float]:
        return pulumi.get(self, "with_compute_unit_price")

    @property
    @pulumi.getter(name="withMemo")
    def with_memo(self) -> Optional[str]:
        return pulumi.get(self, "with_memo")

    @property
    @pulumi.getter
    def ws(self) -> Optional[str]:
        return pulumi.get(self, "ws")


@pulumi.output_type
class ValidatorInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "iconURL":
            suggest = "icon_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ValidatorInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ValidatorInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ValidatorInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 details: Optional[str] = None,
                 icon_url: Optional[str] = None,
                 website: Optional[str] = None):
        pulumi.set(__self__, "name", name)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if icon_url is not None:
            pulumi.set(__self__, "icon_url", icon_url)
        if website is not None:
            pulumi.set(__self__, "website", website)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def details(self) -> Optional[str]:
        return pulumi.get(self, "details")

    @property
    @pulumi.getter(name="iconURL")
    def icon_url(self) -> Optional[str]:
        return pulumi.get(self, "icon_url")

    @property
    @pulumi.getter
    def website(self) -> Optional[str]:
        return pulumi.get(self, "website")


@pulumi.output_type
class VoteAccountKeyPairs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "authWithdrawer":
            suggest = "auth_withdrawer"
        elif key == "voteAccount":
            suggest = "vote_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VoteAccountKeyPairs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VoteAccountKeyPairs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VoteAccountKeyPairs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auth_withdrawer: str,
                 identity: str,
                 vote_account: str):
        pulumi.set(__self__, "auth_withdrawer", auth_withdrawer)
        pulumi.set(__self__, "identity", identity)
        pulumi.set(__self__, "vote_account", vote_account)

    @property
    @pulumi.getter(name="authWithdrawer")
    def auth_withdrawer(self) -> str:
        return pulumi.get(self, "auth_withdrawer")

    @property
    @pulumi.getter
    def identity(self) -> str:
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="voteAccount")
    def vote_account(self) -> str:
        return pulumi.get(self, "vote_account")


