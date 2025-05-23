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
    public sealed class GrpcConfigFilterLimitsTransactions
    {
        public readonly int? AccountExcludeMax;
        public readonly int? AccountIncludeMax;
        public readonly ImmutableArray<string> AccountIncludeReject;
        public readonly int? AccountRequiredMax;
        public readonly bool? Any;
        public readonly int? Max;

        [OutputConstructor]
        private GrpcConfigFilterLimitsTransactions(
            int? accountExcludeMax,

            int? accountIncludeMax,

            ImmutableArray<string> accountIncludeReject,

            int? accountRequiredMax,

            bool? any,

            int? max)
        {
            AccountExcludeMax = accountExcludeMax;
            AccountIncludeMax = accountIncludeMax;
            AccountIncludeReject = accountIncludeReject;
            AccountRequiredMax = accountRequiredMax;
            Any = any;
            Max = max;
        }
    }
}
