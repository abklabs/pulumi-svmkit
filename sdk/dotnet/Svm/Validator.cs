// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Svm.Svm
{
    [SvmResourceType("svm:svm:Validator")]
    public partial class Validator : global::Pulumi.CustomResource
    {
        [Output("flags")]
        public Output<Pulumi.Svm.Agave.Outputs.ValidatorFlags> Flags { get; private set; } = null!;

        [Output("keyPairs")]
        public Output<Pulumi.Svm.Agave.Outputs.ValidatorKeyPairs> KeyPairs { get; private set; } = null!;

        [Output("variant")]
        public Output<string?> Variant { get; private set; } = null!;


        /// <summary>
        /// Create a Validator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Validator(string name, ValidatorArgs args, CustomResourceOptions? options = null)
            : base("svm:svm:Validator", name, args ?? new ValidatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Validator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("svm:svm:Validator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "keyPairs",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Validator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Validator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Validator(name, id, options);
        }
    }

    public sealed class ValidatorArgs : global::Pulumi.ResourceArgs
    {
        [Input("flags", required: true)]
        public Input<Pulumi.Svm.Agave.Inputs.ValidatorFlagsArgs> Flags { get; set; } = null!;

        [Input("keyPairs", required: true)]
        private Input<Pulumi.Svm.Agave.Inputs.ValidatorKeyPairsArgs>? _keyPairs;
        public Input<Pulumi.Svm.Agave.Inputs.ValidatorKeyPairsArgs>? KeyPairs
        {
            get => _keyPairs;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyPairs = Output.Tuple<Input<Pulumi.Svm.Agave.Inputs.ValidatorKeyPairsArgs>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("variant")]
        public Input<string>? Variant { get; set; }

        public ValidatorArgs()
        {
        }
        public static new ValidatorArgs Empty => new ValidatorArgs();
    }
}
