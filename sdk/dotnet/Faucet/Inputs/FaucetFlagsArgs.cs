// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Faucet.Inputs
{

    public sealed class FaucetFlagsArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowIPs")]
        private InputList<string>? _allowIPs;
        public InputList<string> AllowIPs
        {
            get => _allowIPs ?? (_allowIPs = new InputList<string>());
            set => _allowIPs = value;
        }

        [Input("perRequestCap")]
        public Input<int>? PerRequestCap { get; set; }

        [Input("perTimeCap")]
        public Input<int>? PerTimeCap { get; set; }

        [Input("sliceSeconds")]
        public Input<int>? SliceSeconds { get; set; }

        public FaucetFlagsArgs()
        {
        }
        public static new FaucetFlagsArgs Empty => new FaucetFlagsArgs();
    }
}
