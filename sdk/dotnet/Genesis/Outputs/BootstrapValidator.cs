// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Genesis.Outputs
{

    [OutputType]
    public sealed class BootstrapValidator
    {
        public readonly string IdentityPubkey;
        public readonly string StakePubkey;
        public readonly string VotePubkey;

        [OutputConstructor]
        private BootstrapValidator(
            string identityPubkey,

            string stakePubkey,

            string votePubkey)
        {
            IdentityPubkey = identityPubkey;
            StakePubkey = stakePubkey;
            VotePubkey = votePubkey;
        }
    }
}
