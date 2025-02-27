# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'NetworkName',
]


class NetworkName(str, Enum):
    DEVNET = "devnet"
    """
    Solana devnet
    """
    TESTNET = "testnet"
    """
    Solana testnet
    """
    MAINNET_BETA = "mainnet-beta"
    """
    Solana mainnet-beta
    """
