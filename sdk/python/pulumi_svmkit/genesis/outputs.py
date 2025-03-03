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
    'BootstrapAccount',
    'BootstrapValidator',
    'GenesisFlags',
    'PrimordialAccount',
]

@pulumi.output_type
class BootstrapAccount(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "identityPubkey":
            suggest = "identity_pubkey"
        elif key == "stakePubkey":
            suggest = "stake_pubkey"
        elif key == "votePubkey":
            suggest = "vote_pubkey"
        elif key == "balanceLamports":
            suggest = "balance_lamports"
        elif key == "stakeLamports":
            suggest = "stake_lamports"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BootstrapAccount. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BootstrapAccount.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BootstrapAccount.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 identity_pubkey: str,
                 stake_pubkey: str,
                 vote_pubkey: str,
                 balance_lamports: Optional[int] = None,
                 stake_lamports: Optional[int] = None):
        pulumi.set(__self__, "identity_pubkey", identity_pubkey)
        pulumi.set(__self__, "stake_pubkey", stake_pubkey)
        pulumi.set(__self__, "vote_pubkey", vote_pubkey)
        if balance_lamports is not None:
            pulumi.set(__self__, "balance_lamports", balance_lamports)
        if stake_lamports is not None:
            pulumi.set(__self__, "stake_lamports", stake_lamports)

    @property
    @pulumi.getter(name="identityPubkey")
    def identity_pubkey(self) -> str:
        return pulumi.get(self, "identity_pubkey")

    @property
    @pulumi.getter(name="stakePubkey")
    def stake_pubkey(self) -> str:
        return pulumi.get(self, "stake_pubkey")

    @property
    @pulumi.getter(name="votePubkey")
    def vote_pubkey(self) -> str:
        return pulumi.get(self, "vote_pubkey")

    @property
    @pulumi.getter(name="balanceLamports")
    def balance_lamports(self) -> Optional[int]:
        return pulumi.get(self, "balance_lamports")

    @property
    @pulumi.getter(name="stakeLamports")
    def stake_lamports(self) -> Optional[int]:
        return pulumi.get(self, "stake_lamports")


@pulumi.output_type
class BootstrapValidator(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "identityPubkey":
            suggest = "identity_pubkey"
        elif key == "stakePubkey":
            suggest = "stake_pubkey"
        elif key == "votePubkey":
            suggest = "vote_pubkey"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BootstrapValidator. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BootstrapValidator.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BootstrapValidator.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 identity_pubkey: str,
                 stake_pubkey: str,
                 vote_pubkey: str):
        pulumi.set(__self__, "identity_pubkey", identity_pubkey)
        pulumi.set(__self__, "stake_pubkey", stake_pubkey)
        pulumi.set(__self__, "vote_pubkey", vote_pubkey)

    @property
    @pulumi.getter(name="identityPubkey")
    def identity_pubkey(self) -> str:
        return pulumi.get(self, "identity_pubkey")

    @property
    @pulumi.getter(name="stakePubkey")
    def stake_pubkey(self) -> str:
        return pulumi.get(self, "stake_pubkey")

    @property
    @pulumi.getter(name="votePubkey")
    def vote_pubkey(self) -> str:
        return pulumi.get(self, "vote_pubkey")


@pulumi.output_type
class GenesisFlags(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bootstrapValidators":
            suggest = "bootstrap_validators"
        elif key == "ledgerPath":
            suggest = "ledger_path"
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
        elif key == "slotsPerEpoch":
            suggest = "slots_per_epoch"
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
                 bootstrap_validators: Sequence['outputs.BootstrapValidator'],
                 ledger_path: str,
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
                 slots_per_epoch: Optional[int] = None,
                 target_lamports_per_signature: Optional[int] = None,
                 target_signatures_per_slot: Optional[int] = None,
                 target_tick_duration: Optional[int] = None,
                 ticks_per_slot: Optional[int] = None,
                 url: Optional[str] = None,
                 vote_commission_percentage: Optional[int] = None):
        pulumi.set(__self__, "bootstrap_validators", bootstrap_validators)
        pulumi.set(__self__, "ledger_path", ledger_path)
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
        if slots_per_epoch is not None:
            pulumi.set(__self__, "slots_per_epoch", slots_per_epoch)
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
    @pulumi.getter(name="bootstrapValidators")
    def bootstrap_validators(self) -> Sequence['outputs.BootstrapValidator']:
        return pulumi.get(self, "bootstrap_validators")

    @property
    @pulumi.getter(name="ledgerPath")
    def ledger_path(self) -> str:
        return pulumi.get(self, "ledger_path")

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
    @pulumi.getter(name="slotsPerEpoch")
    def slots_per_epoch(self) -> Optional[int]:
        return pulumi.get(self, "slots_per_epoch")

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
class PrimordialAccount(dict):
    def __init__(__self__, *,
                 lamports: int,
                 pubkey: str,
                 data: Optional[str] = None,
                 executable: Optional[bool] = None,
                 owner: Optional[str] = None):
        pulumi.set(__self__, "lamports", lamports)
        pulumi.set(__self__, "pubkey", pubkey)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if executable is not None:
            pulumi.set(__self__, "executable", executable)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)

    @property
    @pulumi.getter
    def lamports(self) -> int:
        return pulumi.get(self, "lamports")

    @property
    @pulumi.getter
    def pubkey(self) -> str:
        return pulumi.get(self, "pubkey")

    @property
    @pulumi.getter
    def data(self) -> Optional[str]:
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def executable(self) -> Optional[bool]:
        return pulumi.get(self, "executable")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        return pulumi.get(self, "owner")


