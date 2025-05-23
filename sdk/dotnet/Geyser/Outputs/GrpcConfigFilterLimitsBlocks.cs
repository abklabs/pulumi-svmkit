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
    public sealed class GrpcConfigFilterLimitsBlocks
    {
        public readonly bool? AccountIncludeAny;
        public readonly int? AccountIncludeMax;
        public readonly ImmutableArray<string> AccountIncludeReject;
        public readonly bool? IncludeAccounts;
        public readonly bool? IncludeEntries;
        public readonly bool? IncludeTransactions;
        public readonly int? Max;

        [OutputConstructor]
        private GrpcConfigFilterLimitsBlocks(
            bool? accountIncludeAny,

            int? accountIncludeMax,

            ImmutableArray<string> accountIncludeReject,

            bool? includeAccounts,

            bool? includeEntries,

            bool? includeTransactions,

            int? max)
        {
            AccountIncludeAny = accountIncludeAny;
            AccountIncludeMax = accountIncludeMax;
            AccountIncludeReject = accountIncludeReject;
            IncludeAccounts = includeAccounts;
            IncludeEntries = includeEntries;
            IncludeTransactions = includeTransactions;
            Max = max;
        }
    }
}
