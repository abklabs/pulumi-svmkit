// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Watchtower.Outputs
{

    [OutputType]
    public sealed class TwilioConfig
    {
        public readonly string AccountSid;
        public readonly string AuthToken;
        public readonly string FromNumber;
        public readonly string ToNumber;

        [OutputConstructor]
        private TwilioConfig(
            string accountSid,

            string authToken,

            string fromNumber,

            string toNumber)
        {
            AccountSid = accountSid;
            AuthToken = authToken;
            FromNumber = fromNumber;
            ToNumber = toNumber;
        }
    }
}
