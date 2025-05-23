// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Firedancer.Outputs
{

    [OutputType]
    public sealed class ConfigRPC
    {
        public readonly bool? BigtableLedgerStorage;
        public readonly bool? ExtendedTxMetadataStorage;
        public readonly bool? FullApi;
        public readonly bool? OnlyKnown;
        public readonly int? Port;
        public readonly bool? Private;
        public readonly bool? PubsubEnableBlockSubscription;
        public readonly bool? PubsubEnableVoteSubscription;
        public readonly bool? TransactionHistory;

        [OutputConstructor]
        private ConfigRPC(
            bool? bigtableLedgerStorage,

            bool? extendedTxMetadataStorage,

            bool? fullApi,

            bool? onlyKnown,

            int? port,

            bool? @private,

            bool? pubsubEnableBlockSubscription,

            bool? pubsubEnableVoteSubscription,

            bool? transactionHistory)
        {
            BigtableLedgerStorage = bigtableLedgerStorage;
            ExtendedTxMetadataStorage = extendedTxMetadataStorage;
            FullApi = fullApi;
            OnlyKnown = onlyKnown;
            Port = port;
            Private = @private;
            PubsubEnableBlockSubscription = pubsubEnableBlockSubscription;
            PubsubEnableVoteSubscription = pubsubEnableVoteSubscription;
            TransactionHistory = transactionHistory;
        }
    }
}
