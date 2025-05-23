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

    public sealed class TunerParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cpuGovernor")]
        public Input<ABKLabs.Svmkit.Tuner.CpuGovernor>? CpuGovernor { get; set; }

        [Input("fs")]
        public Input<Inputs.TunerFsParamsArgs>? Fs { get; set; }

        [Input("kernel")]
        public Input<Inputs.TunerKernelParamsArgs>? Kernel { get; set; }

        [Input("net")]
        public Input<Inputs.TunerNetParamsArgs>? Net { get; set; }

        [Input("vm")]
        public Input<Inputs.TunerVmParamsArgs>? Vm { get; set; }

        public TunerParamsArgs()
        {
        }
        public static new TunerParamsArgs Empty => new TunerParamsArgs();
    }
}
