// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Firedancer.Inputs
{

    public sealed class KeyPairsArgs : global::Pulumi.ResourceArgs
    {
        [Input("identity", required: true)]
        private Input<string>? _identity;
        public Input<string>? Identity
        {
            get => _identity;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _identity = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("voteAccount", required: true)]
        private Input<string>? _voteAccount;
        public Input<string>? VoteAccount
        {
            get => _voteAccount;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _voteAccount = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public KeyPairsArgs()
        {
        }
        public static new KeyPairsArgs Empty => new KeyPairsArgs();
    }
}
