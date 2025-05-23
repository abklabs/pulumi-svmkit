// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Geyser.Outputs
{

    [OutputType]
    public sealed class GrpcConfigTokio
    {
        public readonly ImmutableArray<int> Affinity;
        public readonly int? WorkerThreads;

        [OutputConstructor]
        private GrpcConfigTokio(
            ImmutableArray<int> affinity,

            int? workerThreads)
        {
            Affinity = affinity;
            WorkerThreads = workerThreads;
        }
    }
}
