# Agave Validator On an Existing Host

This example brings up an Agave validator hooked up to testnet, on an existing host.
In order to use this example you must have:

1. A SSH-accessible host running a recent Debian or Ubuntu install.
2. `sudo` privileges for the user you're connecting as.

# Configuration

The following Pulumi configuration must be in place:

- `ssh:host` - a hostname accessible by SSH on the host Pulumi is running on.
- `ssh:user` - the user who will set up the validator (using `sudo`).
- `ssh:privateKey` - (optional) the private key to use with SSH to connect to the host.

For more information on setting configuration, see the [Pulumi documentation](https://www.pulumi.com/docs/iac/concepts/config/#setting-and-getting-configuration-values).
