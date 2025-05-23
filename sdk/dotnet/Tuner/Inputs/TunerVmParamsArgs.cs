// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Tuner.Inputs
{

    public sealed class TunerVmParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("vmDirtyBackgroundRatio")]
        public Input<int>? VmDirtyBackgroundRatio { get; set; }

        [Input("vmDirtyExpireCentisecs")]
        public Input<int>? VmDirtyExpireCentisecs { get; set; }

        [Input("vmDirtyRatio")]
        public Input<int>? VmDirtyRatio { get; set; }

        [Input("vmDirtyWritebackCentisecs")]
        public Input<int>? VmDirtyWritebackCentisecs { get; set; }

        [Input("vmDirtytimeExpireSeconds")]
        public Input<int>? VmDirtytimeExpireSeconds { get; set; }

        [Input("vmMaxMapCount")]
        public Input<int>? VmMaxMapCount { get; set; }

        [Input("vmMinFreeKbytes")]
        public Input<int>? VmMinFreeKbytes { get; set; }

        [Input("vmStatInterval")]
        public Input<int>? VmStatInterval { get; set; }

        [Input("vmSwappiness")]
        public Input<int>? VmSwappiness { get; set; }

        public TunerVmParamsArgs()
        {
        }
        public static new TunerVmParamsArgs Empty => new TunerVmParamsArgs();
    }
}
