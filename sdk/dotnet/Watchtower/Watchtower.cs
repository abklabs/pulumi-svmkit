// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Watchtower
{
    [SvmkitResourceType("svmkit:watchtower:Watchtower")]
    public partial class Watchtower : global::Pulumi.CustomResource
    {
        [Output("connection")]
        public Output<ABKLabs.Svmkit.Ssh.Outputs.Connection> Connection { get; private set; } = null!;

        [Output("environment")]
        public Output<ABKLabs.Svmkit.Solana.Outputs.Environment> Environment { get; private set; } = null!;

        [Output("flags")]
        public Output<Outputs.WatchtowerFlags> Flags { get; private set; } = null!;

        [Output("notifications")]
        public Output<Outputs.NotificationConfig> Notifications { get; private set; } = null!;

        [Output("runnerConfig")]
        public Output<ABKLabs.Svmkit.Runner.Outputs.Config?> RunnerConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Watchtower resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Watchtower(string name, WatchtowerArgs args, CustomResourceOptions? options = null)
            : base("svmkit:watchtower:Watchtower", name, args ?? new WatchtowerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Watchtower(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("svmkit:watchtower:Watchtower", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/abklabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Watchtower resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Watchtower Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Watchtower(name, id, options);
        }
    }

    public sealed class WatchtowerArgs : global::Pulumi.ResourceArgs
    {
        [Input("connection", required: true)]
        public Input<ABKLabs.Svmkit.Ssh.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("environment", required: true)]
        public Input<ABKLabs.Svmkit.Solana.Inputs.EnvironmentArgs> Environment { get; set; } = null!;

        [Input("flags", required: true)]
        public Input<Inputs.WatchtowerFlagsArgs> Flags { get; set; } = null!;

        [Input("notifications", required: true)]
        public Input<Inputs.NotificationConfigArgs> Notifications { get; set; } = null!;

        [Input("runnerConfig")]
        public Input<ABKLabs.Svmkit.Runner.Inputs.ConfigArgs>? RunnerConfig { get; set; }

        public WatchtowerArgs()
        {
        }
        public static new WatchtowerArgs Empty => new WatchtowerArgs();
    }
}
