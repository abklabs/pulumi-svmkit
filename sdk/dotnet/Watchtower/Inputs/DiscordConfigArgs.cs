// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Watchtower.Inputs
{

    public sealed class DiscordConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("webhookUrl", required: true)]
        public Input<string> WebhookUrl { get; set; } = null!;

        public DiscordConfigArgs()
        {
        }
        public static new DiscordConfigArgs Empty => new DiscordConfigArgs();
    }
}
