// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Networkinfo
{
    public static class GetNetworkInfo
    {
        public static Task<GetNetworkInfoResult> InvokeAsync(GetNetworkInfoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkInfoResult>("svmkit:networkinfo:getNetworkInfo", args ?? new GetNetworkInfoArgs(), options.WithDefaults());

        public static Output<GetNetworkInfoResult> Invoke(GetNetworkInfoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkInfoResult>("svmkit:networkinfo:getNetworkInfo", args ?? new GetNetworkInfoInvokeArgs(), options.WithDefaults());

        public static Output<GetNetworkInfoResult> Invoke(GetNetworkInfoInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkInfoResult>("svmkit:networkinfo:getNetworkInfo", args ?? new GetNetworkInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkInfoArgs : global::Pulumi.InvokeArgs
    {
        [Input("networkName", required: true)]
        public ABKLabs.Svmkit.Solana.NetworkName NetworkName { get; set; }

        public GetNetworkInfoArgs()
        {
        }
        public static new GetNetworkInfoArgs Empty => new GetNetworkInfoArgs();
    }

    public sealed class GetNetworkInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("networkName", required: true)]
        public Input<ABKLabs.Svmkit.Solana.NetworkName> NetworkName { get; set; } = null!;

        public GetNetworkInfoInvokeArgs()
        {
        }
        public static new GetNetworkInfoInvokeArgs Empty => new GetNetworkInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkInfoResult
    {
        public readonly ImmutableArray<string> EntryPoint;
        public readonly string GenesisHash;
        public readonly ImmutableArray<string> KnownValidator;
        public readonly ABKLabs.Svmkit.Solana.NetworkName NetworkName;
        public readonly ImmutableArray<string> RpcURL;

        [OutputConstructor]
        private GetNetworkInfoResult(
            ImmutableArray<string> entryPoint,

            string genesisHash,

            ImmutableArray<string> knownValidator,

            ABKLabs.Svmkit.Solana.NetworkName networkName,

            ImmutableArray<string> rpcURL)
        {
            EntryPoint = entryPoint;
            GenesisHash = genesisHash;
            KnownValidator = knownValidator;
            NetworkName = networkName;
            RpcURL = rpcURL;
        }
    }
}