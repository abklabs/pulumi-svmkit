// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Svmkit.Agave.Inputs
{

    public sealed class FlagsArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockProductionMethod", required: true)]
        public Input<string> BlockProductionMethod { get; set; } = null!;

        [Input("dynamicPortRange", required: true)]
        public Input<string> DynamicPortRange { get; set; } = null!;

        [Input("entryPoint")]
        private InputList<string>? _entryPoint;
        public InputList<string> EntryPoint
        {
            get => _entryPoint ?? (_entryPoint = new InputList<string>());
            set => _entryPoint = value;
        }

        [Input("expectedGenesisHash")]
        public Input<string>? ExpectedGenesisHash { get; set; }

        [Input("fullRpcAPI")]
        public Input<bool>? FullRpcAPI { get; set; }

        [Input("fullSnapshotIntervalSlots", required: true)]
        public Input<int> FullSnapshotIntervalSlots { get; set; } = null!;

        [Input("gossipPort", required: true)]
        public Input<int> GossipPort { get; set; } = null!;

        [Input("knownValidator")]
        private InputList<string>? _knownValidator;
        public InputList<string> KnownValidator
        {
            get => _knownValidator ?? (_knownValidator = new InputList<string>());
            set => _knownValidator = value;
        }

        [Input("limitLedgerSize", required: true)]
        public Input<int> LimitLedgerSize { get; set; } = null!;

        [Input("noVoting")]
        public Input<bool>? NoVoting { get; set; }

        [Input("noWaitForVoteToStartLeader", required: true)]
        public Input<bool> NoWaitForVoteToStartLeader { get; set; } = null!;

        [Input("onlyKnownRPC", required: true)]
        public Input<bool> OnlyKnownRPC { get; set; } = null!;

        [Input("privateRPC", required: true)]
        public Input<bool> PrivateRPC { get; set; } = null!;

        [Input("rpcBindAddress", required: true)]
        public Input<string> RpcBindAddress { get; set; } = null!;

        [Input("rpcPort", required: true)]
        public Input<int> RpcPort { get; set; } = null!;

        [Input("tvuReceiveThreads")]
        public Input<int>? TvuReceiveThreads { get; set; }

        [Input("useSnapshotArchivesAtStartup", required: true)]
        public Input<string> UseSnapshotArchivesAtStartup { get; set; } = null!;

        [Input("walRecoveryMode", required: true)]
        public Input<string> WalRecoveryMode { get; set; } = null!;

        public FlagsArgs()
        {
        }
        public static new FlagsArgs Empty => new FlagsArgs();
    }
}
