// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Firedancer.Outputs
{

    [OutputType]
    public sealed class ConfigLog
    {
        public readonly string? Colorize;
        public readonly string? LevelFlush;
        public readonly string? LevelLogfile;
        public readonly string? LevelStderr;
        public readonly string? Path;

        [OutputConstructor]
        private ConfigLog(
            string? colorize,

            string? levelFlush,

            string? levelLogfile,

            string? levelStderr,

            string? path)
        {
            Colorize = colorize;
            LevelFlush = levelFlush;
            LevelLogfile = levelLogfile;
            LevelStderr = levelStderr;
            Path = path;
        }
    }
}