// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Deb.Inputs
{

    public sealed class PackageConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("additional")]
        private InputList<string>? _additional;
        public InputList<string> Additional
        {
            get => _additional ?? (_additional = new InputList<string>());
            set => _additional = value;
        }

        [Input("override")]
        private InputList<Inputs.PackageArgs>? _override;
        public InputList<Inputs.PackageArgs> Override
        {
            get => _override ?? (_override = new InputList<Inputs.PackageArgs>());
            set => _override = value;
        }

        public PackageConfigArgs()
        {
        }
        public static new PackageConfigArgs Empty => new PackageConfigArgs();
    }
}
