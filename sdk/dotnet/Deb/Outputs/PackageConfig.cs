// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Deb.Outputs
{

    [OutputType]
    public sealed class PackageConfig
    {
        public readonly ImmutableArray<string> Additional;
        public readonly ImmutableArray<Outputs.Package> Override;

        [OutputConstructor]
        private PackageConfig(
            ImmutableArray<string> additional,

            ImmutableArray<Outputs.Package> @override)
        {
            Additional = additional;
            Override = @override;
        }
    }
}
