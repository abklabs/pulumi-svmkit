// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Tuner.Inputs
{

    public sealed class TunerFsParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("fsNrOpen")]
        public Input<int>? FsNrOpen { get; set; }

        public TunerFsParamsArgs()
        {
        }
        public static new TunerFsParamsArgs Empty => new TunerFsParamsArgs();
    }
}
