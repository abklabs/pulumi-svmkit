# Pulumi SVMKit
`Pulumi SVMKit` is our Infrastructure-as-Code (IaC) project for managing validators and networks, built on top of Pulumi. This tool allows developers and network operators to easily deploy and manage Solana Virtual Machine (SVM) validators and bootstrap nodes across various environments.
## What Does Pulumi SVMKit Do?

`Pulumi SVMKit` simplifies the process of deploying and managing blockchain infrastructure for permissioned networks. The current components include:

**Validator Installation**: Quickly spin up a validator node on any Debian-based operating system.

**Bootstrap Node Setup**: Launch a bootstrap node to initialize a new SVM-powered blockchain network, allowing you to manage network peers and maintain secure connections.

# Key Features

**Pulumi Integration**: Manage your blockchain infrastructure using Pulumi’s declarative language, enabling reusable and modular deployments.

**Apt Repository Support**: We host an apt repository where validator clients like Agave, Jito, and others are available for quick installation on Debian-based systems. This streamlines the setup process, ensuring you always have access to the latest versions of supported clients. [Explore information about the apt repository here.](https://github.com/abklabs/svmkit)

**Multi-Client Support**: The infrastructure is preconfigured to work with multiple validator clients, allowing flexibility in choosing the right client for your network setup.

# Examples

Explore example deployments and get started quickly:

**Single Validator on Testnet**: Demonstrates deploying a single validator node to a testnet environment.
[Example Code](https://github.com/abklabs/pulumi-svmkit/tree/main/examples/aws-basic-host)

**Bootstrap Node for New Network**: Shows how to set up a bootstrap node to launch a new permissioned or test network.
[Example Code](https://github.com/abklabs/pulumi-svmkit/tree/main/examples/aws-svm-bootstrap)

# Upcoming Features
We are actively working on expanding `Pulumi SVMKit` with new examples and features:

**Genesis to Multi-Node Cluster**: We’re adding comprehensive examples showing how to go from a genesis block to a fully operational multi-node cluster. This will help users deploy production-grade permissioned environments.

**Permissioned Environments Support**: Next week, we will introduce a new feature to enable permissioned environments, making it easier for operators to manage node permissions, restrict access, and control governance in their blockchain networks.

Stay tuned for these updates!

# Getting Started / Developing `Pulumi SVMKit`

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
