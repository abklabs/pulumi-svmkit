// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Tuner.Outputs
{

    [OutputType]
    public sealed class TunerParams
    {
        public readonly ABKLabs.Svmkit.Tuner.CpuGovernor? CpuGovernor;
        public readonly Outputs.TunerFsParams? Fs;
        public readonly Outputs.TunerKernelParams? Kernel;
        public readonly Outputs.TunerNetParams? Net;
        public readonly Outputs.TunerVmParams? Vm;

        [OutputConstructor]
        private TunerParams(
            ABKLabs.Svmkit.Tuner.CpuGovernor? cpuGovernor,

            Outputs.TunerFsParams? fs,

            Outputs.TunerKernelParams? kernel,

            Outputs.TunerNetParams? net,

            Outputs.TunerVmParams? vm)
        {
            CpuGovernor = cpuGovernor;
            Fs = fs;
            Kernel = kernel;
            Net = net;
            Vm = vm;
        }
    }
}
