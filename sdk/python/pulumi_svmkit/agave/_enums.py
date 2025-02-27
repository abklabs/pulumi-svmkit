# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'Variant',
]


class Variant(str, Enum):
    SOLANA = "solana"
    """
    The Solana validator
    """
    AGAVE = "agave"
    """
    The Agave validator
    """
    POWERLEDGER = "powerledger"
    """
    The Powerledger validator
    """
    JITO = "jito"
    """
    The Jito validator
    """
    PYTH = "pyth"
    """
    The Pyth validator
    """
    MANTIS = "mantis"
    """
    The Mantis validator
    """
    XEN = "xen"
    """
    The Xen validator
    """
    TACHYON = "tachyon"
    """
    The Tachyon validator
    """
