// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Watchtower.Inputs
{

    public sealed class SlackConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("webhookUrl", required: true)]
        public Input<string> WebhookUrl { get; set; } = null!;

        public SlackConfigArgs()
        {
        }
        public static new SlackConfigArgs Empty => new SlackConfigArgs();
    }
}
