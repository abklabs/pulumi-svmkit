// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Solana.Inputs
{

    public sealed class EnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("rpcURL", required: true)]
        public Input<string> RpcURL { get; set; } = null!;

        public EnvironmentArgs()
        {
        }
        public static new EnvironmentArgs Empty => new EnvironmentArgs();
    }
}
