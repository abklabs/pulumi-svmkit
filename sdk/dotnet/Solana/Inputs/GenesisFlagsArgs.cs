// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Solana.Inputs
{

    public sealed class GenesisFlagsArgs : global::Pulumi.ResourceArgs
    {
        [Input("bootstrapStakeAuthorizedPubkey")]
        public Input<string>? BootstrapStakeAuthorizedPubkey { get; set; }

        [Input("bootstrapValidatorLamports")]
        public Input<int>? BootstrapValidatorLamports { get; set; }

        [Input("bootstrapValidatorStakeLamports")]
        public Input<int>? BootstrapValidatorStakeLamports { get; set; }

        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        [Input("deactivateFeatures")]
        private InputList<string>? _deactivateFeatures;
        public InputList<string> DeactivateFeatures
        {
            get => _deactivateFeatures ?? (_deactivateFeatures = new InputList<string>());
            set => _deactivateFeatures = value;
        }

        [Input("enableWarmupEpochs")]
        public Input<bool>? EnableWarmupEpochs { get; set; }

        [Input("extraFlags")]
        private InputList<string>? _extraFlags;
        public InputList<string> ExtraFlags
        {
            get => _extraFlags ?? (_extraFlags = new InputList<string>());
            set => _extraFlags = value;
        }

        [Input("faucetLamports")]
        public Input<int>? FaucetLamports { get; set; }

        [Input("faucetPubkey")]
        public Input<string>? FaucetPubkey { get; set; }

        [Input("feeBurnPercentage")]
        public Input<int>? FeeBurnPercentage { get; set; }

        [Input("hashesPerTick")]
        public Input<string>? HashesPerTick { get; set; }

        [Input("identityPubkey", required: true)]
        public Input<string> IdentityPubkey { get; set; } = null!;

        [Input("inflation")]
        public Input<string>? Inflation { get; set; }

        [Input("lamportsPerByteYear")]
        public Input<int>? LamportsPerByteYear { get; set; }

        [Input("ledgerPath", required: true)]
        public Input<string> LedgerPath { get; set; } = null!;

        [Input("maxGenesisArchiveUnpackedSize")]
        public Input<int>? MaxGenesisArchiveUnpackedSize { get; set; }

        [Input("rentBurnPercentage")]
        public Input<int>? RentBurnPercentage { get; set; }

        [Input("rentExemptionThreshold")]
        public Input<int>? RentExemptionThreshold { get; set; }

        [Input("slotPerEpoch")]
        public Input<int>? SlotPerEpoch { get; set; }

        [Input("stakePubkey", required: true)]
        public Input<string> StakePubkey { get; set; } = null!;

        [Input("targetLamportsPerSignature")]
        public Input<int>? TargetLamportsPerSignature { get; set; }

        [Input("targetSignaturesPerSlot")]
        public Input<int>? TargetSignaturesPerSlot { get; set; }

        [Input("targetTickDuration")]
        public Input<int>? TargetTickDuration { get; set; }

        [Input("ticksPerSlot")]
        public Input<int>? TicksPerSlot { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("voteCommissionPercentage")]
        public Input<int>? VoteCommissionPercentage { get; set; }

        [Input("votePubkey", required: true)]
        public Input<string> VotePubkey { get; set; } = null!;

        public GenesisFlagsArgs()
        {
        }
        public static new GenesisFlagsArgs Empty => new GenesisFlagsArgs();
    }
}
