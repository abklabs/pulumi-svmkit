# pulumi-svmkit

# Developing pulumi-svmkit

1. For OS X, install the following via homebrew:

```bash
brew install node yarn python python-setuptools go dotnet pulumi
```

2. Build and install a local version of the component resource:

```bash
make install
```

3. Attempt to use the `examples/aws-basic-host` example.

```bash
cd examples/aws-basic-host
yarn install
yarn link @pulumi/svmkit
PATH=$PATH:$PWD/../../bin
pulumi install
pulumi up
```

Note: You will need AWS credentials that pulumi can find to test this example. `SSH_PRIVATE_KEY` and `PUBLIC_DNS_NAME` can be used to access the server; the default username for Debian is `admin`.
