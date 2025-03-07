// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Account
{
    [SvmkitResourceType("svmkit:account:VoteAccount")]
    public partial class VoteAccount : global::Pulumi.CustomResource
    {
        [Output("authVoterPubkey")]
        public Output<string?> AuthVoterPubkey { get; private set; } = null!;

        [Output("closeRecipientPubkey")]
        public Output<string?> CloseRecipientPubkey { get; private set; } = null!;

        [Output("connection")]
        public Output<ABKLabs.Svmkit.Ssh.Outputs.Connection> Connection { get; private set; } = null!;

        [Output("keyPairs")]
        public Output<ABKLabs.Svmkit.Solana.Outputs.VoteAccountKeyPairs> KeyPairs { get; private set; } = null!;

        [Output("runnerConfig")]
        public Output<ABKLabs.Svmkit.Runner.Outputs.Config?> RunnerConfig { get; private set; } = null!;


        /// <summary>
        /// Create a VoteAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VoteAccount(string name, VoteAccountArgs args, CustomResourceOptions? options = null)
            : base("svmkit:account:VoteAccount", name, args ?? new VoteAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VoteAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("svmkit:account:VoteAccount", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VoteAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VoteAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VoteAccount(name, id, options);
        }
    }

    public sealed class VoteAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("authVoterPubkey")]
        public Input<string>? AuthVoterPubkey { get; set; }

        [Input("closeRecipientPubkey")]
        public Input<string>? CloseRecipientPubkey { get; set; }

        [Input("connection", required: true)]
        public Input<ABKLabs.Svmkit.Ssh.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("keyPairs", required: true)]
        public Input<ABKLabs.Svmkit.Solana.Inputs.VoteAccountKeyPairsArgs> KeyPairs { get; set; } = null!;

        [Input("runnerConfig")]
        public Input<ABKLabs.Svmkit.Runner.Inputs.ConfigArgs>? RunnerConfig { get; set; }

        public VoteAccountArgs()
        {
        }
        public static new VoteAccountArgs Empty => new VoteAccountArgs();
    }
}
