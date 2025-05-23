// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Solana.Outputs
{

    [OutputType]
    public sealed class StakeAccountKeyPairs
    {
        public readonly string StakeAccount;
        public readonly string VoteAccount;

        [OutputConstructor]
        private StakeAccountKeyPairs(
            string stakeAccount,

            string voteAccount)
        {
            StakeAccount = stakeAccount;
            VoteAccount = voteAccount;
        }
    }
}
