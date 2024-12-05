// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Firedancer.Inputs
{

    public sealed class ConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("consensus")]
        public Input<Inputs.ConfigConsensusArgs>? Consensus { get; set; }

        [Input("dynamicPortRange")]
        public Input<string>? DynamicPortRange { get; set; }

        [Input("extraConfig")]
        private InputList<string>? _extraConfig;
        public InputList<string> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputList<string>());
            set => _extraConfig = value;
        }

        [Input("gossip")]
        public Input<Inputs.ConfigGossipArgs>? Gossip { get; set; }

        [Input("hugetlbfs")]
        public Input<Inputs.ConfigHugeTLBFSArgs>? Hugetlbfs { get; set; }

        [Input("layout")]
        public Input<Inputs.ConfigLayoutArgs>? Layout { get; set; }

        [Input("ledger")]
        public Input<Inputs.ConfigLedgerArgs>? Ledger { get; set; }

        [Input("log")]
        public Input<Inputs.ConfigLogArgs>? Log { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("reporting")]
        public Input<Inputs.ConfigReportingArgs>? Reporting { get; set; }

        [Input("rpc")]
        public Input<Inputs.ConfigRPCArgs>? Rpc { get; set; }

        [Input("scratchDirectory")]
        public Input<string>? ScratchDirectory { get; set; }

        [Input("snapshots")]
        public Input<Inputs.ConfigSnapshotsArgs>? Snapshots { get; set; }

        [Input("user")]
        public Input<string>? User { get; set; }

        public ConfigArgs()
        {
        }
        public static new ConfigArgs Empty => new ConfigArgs();
    }
}
