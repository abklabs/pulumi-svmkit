# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .key_pair import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_svmkit.agave as __agave
    agave = __agave
    import pulumi_svmkit.ssh as __ssh
    ssh = __ssh
    import pulumi_svmkit.validator as __validator
    validator = __validator
else:
    agave = _utilities.lazy_import('pulumi_svmkit.agave')
    ssh = _utilities.lazy_import('pulumi_svmkit.ssh')
    validator = _utilities.lazy_import('pulumi_svmkit.validator')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "svmkit",
  "mod": "index",
  "fqn": "pulumi_svmkit",
  "classes": {
   "svmkit:index:KeyPair": "KeyPair"
  }
 },
 {
  "pkg": "svmkit",
  "mod": "validator",
  "fqn": "pulumi_svmkit.validator",
  "classes": {
   "svmkit:validator:Agave": "Agave"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "svmkit",
  "token": "pulumi:providers:svmkit",
  "fqn": "pulumi_svmkit",
  "class": "Provider"
 }
]
"""
)
