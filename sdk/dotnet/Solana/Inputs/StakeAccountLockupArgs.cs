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

    public sealed class StakeAccountLockupArgs : global::Pulumi.ResourceArgs
    {
        [Input("custodianPubkey", required: true)]
        public Input<string> CustodianPubkey { get; set; } = null!;

        [Input("epochAvailable", required: true)]
        public Input<int> EpochAvailable { get; set; } = null!;

        public StakeAccountLockupArgs()
        {
        }
        public static new StakeAccountLockupArgs Empty => new StakeAccountLockupArgs();
    }
}
