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
    public sealed class Package
    {
        public readonly string Name;
        public readonly string? Path;
        public readonly string? TargetRelease;
        public readonly string? Version;

        [OutputConstructor]
        private Package(
            string name,

            string? path,

            string? targetRelease,

            string? version)
        {
            Name = name;
            Path = path;
            TargetRelease = targetRelease;
            Version = version;
        }
    }
}
