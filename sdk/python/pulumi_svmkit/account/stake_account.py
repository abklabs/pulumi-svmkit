# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from .. import solana as _solana
from .. import ssh as _ssh

__all__ = ['StakeAccountArgs', 'StakeAccount']

@pulumi.input_type
class StakeAccountArgs:
    def __init__(__self__, *,
                 amount: pulumi.Input[float],
                 connection: pulumi.Input['_ssh.ConnectionArgs'],
                 key_pairs: pulumi.Input['_solana.StakeAccountKeyPairsArgs']):
        """
        The set of arguments for constructing a StakeAccount resource.
        """
        pulumi.set(__self__, "amount", amount)
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "key_pairs", key_pairs)

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Input[float]:
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: pulumi.Input[float]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['_ssh.ConnectionArgs']:
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['_ssh.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter(name="keyPairs")
    def key_pairs(self) -> pulumi.Input['_solana.StakeAccountKeyPairsArgs']:
        return pulumi.get(self, "key_pairs")

    @key_pairs.setter
    def key_pairs(self, value: pulumi.Input['_solana.StakeAccountKeyPairsArgs']):
        pulumi.set(self, "key_pairs", value)


class StakeAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount: Optional[pulumi.Input[float]] = None,
                 connection: Optional[pulumi.Input[Union['_ssh.ConnectionArgs', '_ssh.ConnectionArgsDict']]] = None,
                 key_pairs: Optional[pulumi.Input[Union['_solana.StakeAccountKeyPairsArgs', '_solana.StakeAccountKeyPairsArgsDict']]] = None,
                 __props__=None):
        """
        Create a StakeAccount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StakeAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a StakeAccount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param StakeAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StakeAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount: Optional[pulumi.Input[float]] = None,
                 connection: Optional[pulumi.Input[Union['_ssh.ConnectionArgs', '_ssh.ConnectionArgsDict']]] = None,
                 key_pairs: Optional[pulumi.Input[Union['_solana.StakeAccountKeyPairsArgs', '_solana.StakeAccountKeyPairsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StakeAccountArgs.__new__(StakeAccountArgs)

            if amount is None and not opts.urn:
                raise TypeError("Missing required property 'amount'")
            __props__.__dict__["amount"] = amount
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            if key_pairs is None and not opts.urn:
                raise TypeError("Missing required property 'key_pairs'")
            __props__.__dict__["key_pairs"] = key_pairs
        super(StakeAccount, __self__).__init__(
            'svmkit:account:StakeAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StakeAccount':
        """
        Get an existing StakeAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StakeAccountArgs.__new__(StakeAccountArgs)

        __props__.__dict__["amount"] = None
        __props__.__dict__["connection"] = None
        __props__.__dict__["key_pairs"] = None
        return StakeAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Output[float]:
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['_ssh.outputs.Connection']:
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="keyPairs")
    def key_pairs(self) -> pulumi.Output['_solana.outputs.StakeAccountKeyPairs']:
        return pulumi.get(self, "key_pairs")

