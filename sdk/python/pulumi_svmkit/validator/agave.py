# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from .. import agave
from .. import agave as _agave
from .. import deb as _deb
from .. import deletion
from .. import geyser as _geyser
from .. import runner as _runner
from .. import solana as _solana
from .. import ssh as _ssh

__all__ = ['AgaveArgs', 'Agave']

@pulumi.input_type
class AgaveArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['_ssh.ConnectionArgs'],
                 flags: pulumi.Input['_agave.FlagsArgs'],
                 key_pairs: pulumi.Input['_agave.KeyPairsArgs'],
                 deletion_policy: Optional[pulumi.Input['deletion.Policy']] = None,
                 environment: Optional[pulumi.Input['_solana.EnvironmentArgs']] = None,
                 geyser_plugin: Optional[pulumi.Input['_geyser.GeyserPluginArgs']] = None,
                 info: Optional[pulumi.Input['_solana.ValidatorInfoArgs']] = None,
                 metrics: Optional[pulumi.Input['_agave.MetricsArgs']] = None,
                 runner_config: Optional[pulumi.Input['_runner.ConfigArgs']] = None,
                 shutdown_policy: Optional[pulumi.Input['_agave.ShutdownPolicyArgs']] = None,
                 startup_policy: Optional[pulumi.Input['_agave.StartupPolicyArgs']] = None,
                 timeout_config: Optional[pulumi.Input['_agave.TimeoutConfigArgs']] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 variant: Optional[pulumi.Input['agave.Variant']] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Agave resource.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "flags", flags)
        pulumi.set(__self__, "key_pairs", key_pairs)
        if deletion_policy is not None:
            pulumi.set(__self__, "deletion_policy", deletion_policy)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if geyser_plugin is not None:
            pulumi.set(__self__, "geyser_plugin", geyser_plugin)
        if info is not None:
            pulumi.set(__self__, "info", info)
        if metrics is not None:
            pulumi.set(__self__, "metrics", metrics)
        if runner_config is not None:
            pulumi.set(__self__, "runner_config", runner_config)
        if shutdown_policy is not None:
            pulumi.set(__self__, "shutdown_policy", shutdown_policy)
        if startup_policy is not None:
            pulumi.set(__self__, "startup_policy", startup_policy)
        if timeout_config is not None:
            pulumi.set(__self__, "timeout_config", timeout_config)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)
        if variant is not None:
            pulumi.set(__self__, "variant", variant)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['_ssh.ConnectionArgs']:
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['_ssh.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter
    def flags(self) -> pulumi.Input['_agave.FlagsArgs']:
        return pulumi.get(self, "flags")

    @flags.setter
    def flags(self, value: pulumi.Input['_agave.FlagsArgs']):
        pulumi.set(self, "flags", value)

    @property
    @pulumi.getter(name="keyPairs")
    def key_pairs(self) -> pulumi.Input['_agave.KeyPairsArgs']:
        return pulumi.get(self, "key_pairs")

    @key_pairs.setter
    def key_pairs(self, value: pulumi.Input['_agave.KeyPairsArgs']):
        pulumi.set(self, "key_pairs", value)

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> Optional[pulumi.Input['deletion.Policy']]:
        return pulumi.get(self, "deletion_policy")

    @deletion_policy.setter
    def deletion_policy(self, value: Optional[pulumi.Input['deletion.Policy']]):
        pulumi.set(self, "deletion_policy", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input['_solana.EnvironmentArgs']]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input['_solana.EnvironmentArgs']]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="geyserPlugin")
    def geyser_plugin(self) -> Optional[pulumi.Input['_geyser.GeyserPluginArgs']]:
        return pulumi.get(self, "geyser_plugin")

    @geyser_plugin.setter
    def geyser_plugin(self, value: Optional[pulumi.Input['_geyser.GeyserPluginArgs']]):
        pulumi.set(self, "geyser_plugin", value)

    @property
    @pulumi.getter
    def info(self) -> Optional[pulumi.Input['_solana.ValidatorInfoArgs']]:
        return pulumi.get(self, "info")

    @info.setter
    def info(self, value: Optional[pulumi.Input['_solana.ValidatorInfoArgs']]):
        pulumi.set(self, "info", value)

    @property
    @pulumi.getter
    def metrics(self) -> Optional[pulumi.Input['_agave.MetricsArgs']]:
        return pulumi.get(self, "metrics")

    @metrics.setter
    def metrics(self, value: Optional[pulumi.Input['_agave.MetricsArgs']]):
        pulumi.set(self, "metrics", value)

    @property
    @pulumi.getter(name="runnerConfig")
    def runner_config(self) -> Optional[pulumi.Input['_runner.ConfigArgs']]:
        return pulumi.get(self, "runner_config")

    @runner_config.setter
    def runner_config(self, value: Optional[pulumi.Input['_runner.ConfigArgs']]):
        pulumi.set(self, "runner_config", value)

    @property
    @pulumi.getter(name="shutdownPolicy")
    def shutdown_policy(self) -> Optional[pulumi.Input['_agave.ShutdownPolicyArgs']]:
        return pulumi.get(self, "shutdown_policy")

    @shutdown_policy.setter
    def shutdown_policy(self, value: Optional[pulumi.Input['_agave.ShutdownPolicyArgs']]):
        pulumi.set(self, "shutdown_policy", value)

    @property
    @pulumi.getter(name="startupPolicy")
    def startup_policy(self) -> Optional[pulumi.Input['_agave.StartupPolicyArgs']]:
        return pulumi.get(self, "startup_policy")

    @startup_policy.setter
    def startup_policy(self, value: Optional[pulumi.Input['_agave.StartupPolicyArgs']]):
        pulumi.set(self, "startup_policy", value)

    @property
    @pulumi.getter(name="timeoutConfig")
    def timeout_config(self) -> Optional[pulumi.Input['_agave.TimeoutConfigArgs']]:
        return pulumi.get(self, "timeout_config")

    @timeout_config.setter
    def timeout_config(self, value: Optional[pulumi.Input['_agave.TimeoutConfigArgs']]):
        pulumi.set(self, "timeout_config", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Sequence[Any]]]:
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Sequence[Any]]]):
        pulumi.set(self, "triggers", value)

    @property
    @pulumi.getter
    def variant(self) -> Optional[pulumi.Input['agave.Variant']]:
        return pulumi.get(self, "variant")

    @variant.setter
    def variant(self, value: Optional[pulumi.Input['agave.Variant']]):
        pulumi.set(self, "variant", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


@pulumi.type_token("svmkit:validator:Agave")
class Agave(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[Union['_ssh.ConnectionArgs', '_ssh.ConnectionArgsDict']]] = None,
                 deletion_policy: Optional[pulumi.Input['deletion.Policy']] = None,
                 environment: Optional[pulumi.Input[Union['_solana.EnvironmentArgs', '_solana.EnvironmentArgsDict']]] = None,
                 flags: Optional[pulumi.Input[Union['_agave.FlagsArgs', '_agave.FlagsArgsDict']]] = None,
                 geyser_plugin: Optional[pulumi.Input[Union['_geyser.GeyserPluginArgs', '_geyser.GeyserPluginArgsDict']]] = None,
                 info: Optional[pulumi.Input[Union['_solana.ValidatorInfoArgs', '_solana.ValidatorInfoArgsDict']]] = None,
                 key_pairs: Optional[pulumi.Input[Union['_agave.KeyPairsArgs', '_agave.KeyPairsArgsDict']]] = None,
                 metrics: Optional[pulumi.Input[Union['_agave.MetricsArgs', '_agave.MetricsArgsDict']]] = None,
                 runner_config: Optional[pulumi.Input[Union['_runner.ConfigArgs', '_runner.ConfigArgsDict']]] = None,
                 shutdown_policy: Optional[pulumi.Input[Union['_agave.ShutdownPolicyArgs', '_agave.ShutdownPolicyArgsDict']]] = None,
                 startup_policy: Optional[pulumi.Input[Union['_agave.StartupPolicyArgs', '_agave.StartupPolicyArgsDict']]] = None,
                 timeout_config: Optional[pulumi.Input[Union['_agave.TimeoutConfigArgs', '_agave.TimeoutConfigArgsDict']]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 variant: Optional[pulumi.Input['agave.Variant']] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a Agave resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgaveArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Agave resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AgaveArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgaveArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[Union['_ssh.ConnectionArgs', '_ssh.ConnectionArgsDict']]] = None,
                 deletion_policy: Optional[pulumi.Input['deletion.Policy']] = None,
                 environment: Optional[pulumi.Input[Union['_solana.EnvironmentArgs', '_solana.EnvironmentArgsDict']]] = None,
                 flags: Optional[pulumi.Input[Union['_agave.FlagsArgs', '_agave.FlagsArgsDict']]] = None,
                 geyser_plugin: Optional[pulumi.Input[Union['_geyser.GeyserPluginArgs', '_geyser.GeyserPluginArgsDict']]] = None,
                 info: Optional[pulumi.Input[Union['_solana.ValidatorInfoArgs', '_solana.ValidatorInfoArgsDict']]] = None,
                 key_pairs: Optional[pulumi.Input[Union['_agave.KeyPairsArgs', '_agave.KeyPairsArgsDict']]] = None,
                 metrics: Optional[pulumi.Input[Union['_agave.MetricsArgs', '_agave.MetricsArgsDict']]] = None,
                 runner_config: Optional[pulumi.Input[Union['_runner.ConfigArgs', '_runner.ConfigArgsDict']]] = None,
                 shutdown_policy: Optional[pulumi.Input[Union['_agave.ShutdownPolicyArgs', '_agave.ShutdownPolicyArgsDict']]] = None,
                 startup_policy: Optional[pulumi.Input[Union['_agave.StartupPolicyArgs', '_agave.StartupPolicyArgsDict']]] = None,
                 timeout_config: Optional[pulumi.Input[Union['_agave.TimeoutConfigArgs', '_agave.TimeoutConfigArgsDict']]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 variant: Optional[pulumi.Input['agave.Variant']] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AgaveArgs.__new__(AgaveArgs)

            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["deletion_policy"] = deletion_policy
            __props__.__dict__["environment"] = environment
            if flags is None and not opts.urn:
                raise TypeError("Missing required property 'flags'")
            __props__.__dict__["flags"] = flags
            __props__.__dict__["geyser_plugin"] = geyser_plugin
            __props__.__dict__["info"] = info
            if key_pairs is None and not opts.urn:
                raise TypeError("Missing required property 'key_pairs'")
            __props__.__dict__["key_pairs"] = key_pairs
            __props__.__dict__["metrics"] = metrics
            __props__.__dict__["runner_config"] = runner_config
            __props__.__dict__["shutdown_policy"] = shutdown_policy
            __props__.__dict__["startup_policy"] = startup_policy
            __props__.__dict__["timeout_config"] = timeout_config
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["variant"] = variant
            __props__.__dict__["version"] = version
            __props__.__dict__["systemd_service_name"] = None
        super(Agave, __self__).__init__(
            'svmkit:validator:Agave',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Agave':
        """
        Get an existing Agave resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgaveArgs.__new__(AgaveArgs)

        __props__.__dict__["connection"] = None
        __props__.__dict__["deletion_policy"] = None
        __props__.__dict__["environment"] = None
        __props__.__dict__["flags"] = None
        __props__.__dict__["geyser_plugin"] = None
        __props__.__dict__["info"] = None
        __props__.__dict__["key_pairs"] = None
        __props__.__dict__["metrics"] = None
        __props__.__dict__["runner_config"] = None
        __props__.__dict__["shutdown_policy"] = None
        __props__.__dict__["startup_policy"] = None
        __props__.__dict__["systemd_service_name"] = None
        __props__.__dict__["timeout_config"] = None
        __props__.__dict__["triggers"] = None
        __props__.__dict__["variant"] = None
        __props__.__dict__["version"] = None
        return Agave(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['_ssh.outputs.Connection']:
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> pulumi.Output[Optional['deletion.Policy']]:
        return pulumi.get(self, "deletion_policy")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional['_solana.outputs.Environment']]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def flags(self) -> pulumi.Output['_agave.outputs.Flags']:
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter(name="geyserPlugin")
    def geyser_plugin(self) -> pulumi.Output[Optional['_geyser.outputs.GeyserPlugin']]:
        return pulumi.get(self, "geyser_plugin")

    @property
    @pulumi.getter
    def info(self) -> pulumi.Output[Optional['_solana.outputs.ValidatorInfo']]:
        return pulumi.get(self, "info")

    @property
    @pulumi.getter(name="keyPairs")
    def key_pairs(self) -> pulumi.Output['_agave.outputs.KeyPairs']:
        return pulumi.get(self, "key_pairs")

    @property
    @pulumi.getter
    def metrics(self) -> pulumi.Output[Optional['_agave.outputs.Metrics']]:
        return pulumi.get(self, "metrics")

    @property
    @pulumi.getter(name="runnerConfig")
    def runner_config(self) -> pulumi.Output[Optional['_runner.outputs.Config']]:
        return pulumi.get(self, "runner_config")

    @property
    @pulumi.getter(name="shutdownPolicy")
    def shutdown_policy(self) -> pulumi.Output[Optional['_agave.outputs.ShutdownPolicy']]:
        return pulumi.get(self, "shutdown_policy")

    @property
    @pulumi.getter(name="startupPolicy")
    def startup_policy(self) -> pulumi.Output[Optional['_agave.outputs.StartupPolicy']]:
        return pulumi.get(self, "startup_policy")

    @property
    @pulumi.getter(name="systemdServiceName")
    def systemd_service_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "systemd_service_name")

    @property
    @pulumi.getter(name="timeoutConfig")
    def timeout_config(self) -> pulumi.Output[Optional['_agave.outputs.TimeoutConfig']]:
        return pulumi.get(self, "timeout_config")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Optional[Sequence[Any]]]:
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter
    def variant(self) -> pulumi.Output[Optional['agave.Variant']]:
        return pulumi.get(self, "variant")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "version")

