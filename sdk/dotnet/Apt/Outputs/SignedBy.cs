// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Apt.Outputs
{

    [OutputType]
    public sealed class SignedBy
    {
        public readonly ImmutableArray<string> Paths;
        public readonly string? PublicKey;

        [OutputConstructor]
        private SignedBy(
            ImmutableArray<string> paths,

            string? publicKey)
        {
            Paths = paths;
            PublicKey = publicKey;
        }
    }
}
