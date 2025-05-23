// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Genesis.Inputs
{

    public sealed class PrimordialAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("executable")]
        public Input<bool>? Executable { get; set; }

        [Input("lamports", required: true)]
        public Input<int> Lamports { get; set; } = null!;

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("pubkey", required: true)]
        public Input<string> Pubkey { get; set; } = null!;

        public PrimordialAccountArgs()
        {
        }
        public static new PrimordialAccountArgs Empty => new PrimordialAccountArgs();
    }
}
