# Pulumi SVMKit
`Pulumi SVMKit` is our Infrastructure-as-Code (IaC) project for managing validators and networks, built on top of Pulumi. This tool allows developers and network operators to easily deploy and manage Solana Virtual Machine (SVM) validators and bootstrap nodes across various environments.

## Disclaimer

**This project is intended for use in testing environments only and is not yet considered production-ready.** 

We are actively developing and improving the functionality. Users should exercise caution when deploying this software in mission-critical or production systems. Please report any issues or feedback to help us enhance its stability and reliability.

## What Does Pulumi SVMKit Do?

`Pulumi SVMKit` simplifies the process of deploying and managing blockchain infrastructure for permissioned and permissionless networks. The current components include:

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

4. Import the repository's GPG key:

```bash
curl -s https://apt.abklabs.com/keys/abklabs-archive-dev.asc | sudo apt-key add -
```

5. Update your package lists again:

```bash
sudo apt-get update
```

6. Install SVMKit's build of the Agave validator:

```bash
sudo apt-get install zuma-agave-validator
```

=======

# Pulumi Native Provider Boilerplate

This repository is a boilerplate showing how to create and locally test a native Pulumi provider.

## Authoring a Pulumi Native Provider

This boilerplate creates a working Pulumi-owned provider named `svm`.
It implements a random number generator that you can [build and test out for yourself](#test-against-the-example) and then replace the Random code with code specific to your provider.

### Prerequisites

Prerequisites for this repository are already satisfied by the [Pulumi Devcontainer](https://github.com/pulumi/devcontainer) if you are using Github Codespaces, or VSCode.

If you are not using VSCode, you will need to ensure the following tools are installed and present in your `$PATH`:

- [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
- [Go 1.21](https://golang.org/dl/) or 1.latest
- [NodeJS](https://nodejs.org/en/) 14.x. We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations.
- [Yarn](https://yarnpkg.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Python](https://www.python.org/downloads/) (called as `python3`). For recent versions of MacOS, the system-installed version is fine.
- [.NET](https://dotnet.microsoft.com/download)

### Build & test the boilerplate XYZ provider

1. Create a new Github CodeSpaces environment using this repository.
1. Open a terminal in the CodeSpaces environment.
1. Run `make build install` to build and install the provider.
1. Run `make gen_examples` to generate the example programs in `examples/` off of the source `examples/yaml` example program.
1. Run `make up` to run the example program in `examples/yaml`.
1. Run `make down` to tear down the example program.

### Creating a new provider repository

Pulumi offers this repository as a [GitHub template repository](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template) for convenience. From this repository:

1. Click "Use this template".
1. Set the following options:
   - Owner: pulumi
   - Repository name: pulumi-xyz-native (replace "xyz" with the name of your provider)
   - Description: Pulumi provider for xyz
   - # Repository type: Public
   * Owner: pulumi
   * Repository name: pulumi-svm-native (replace "svm" with the name of your provider)
   * Description: Pulumi provider for svm
   * Repository type: Public
     > > > > > > > c183d24 (feat: Setup provider svm and port initial KeyPair component)
1. Clone the generated repository.

From the templated repository:

1. Run the following command to update files to use the name of your provider (third-party: use your GitHub organization/username):

   ```bash
   make prepare NAME=foo REPOSITORY=github.com/pulumi/pulumi-foo ORG=myorg
   ```

   This will do the following:

   - rename folders in `provider/cmd` to `pulumi-resource-{NAME}`
   - replace dependencies in `provider/go.mod` to reflect your repository name
   - find and replace all instances of the boilerplate `svm` with the `NAME` of your provider.
   - find and replace all instances of the boilerplate `abklabs` with the `ORG` of your provider.
   - replace all instances of the `github.com/abklabs/pulumi-svm` repository with the `REPOSITORY` location

#### Build the provider and install the plugin

```bash
$ make build install
```

This will:

1. Create the SDK codegen binary and place it in a `./bin` folder (gitignored)
2. Create the provider binary and place it in the `./bin` folder (gitignored)
3. Generate the dotnet, Go, Node, and Python SDKs and place them in the `./sdk` folder
4. Install the provider on your machine.

#### Test against the example

```bash
$ cd examples/simple
$ yarn link @pulumi/svm
$ yarn install
$ pulumi stack init test
$ pulumi up
```

Now that you have completed all of the above steps, you have a working provider that generates a random string for you.

#### A brief repository overview

You now have:

1. A `provider/` folder containing the building and implementation logic
   1. `cmd/pulumi-resource-svm/main.go` - holds the provider's sample implementation logic.
2. `deployment-templates` - a set of files to help you around deployment and publication
3. `sdk` - holds the generated code libraries created by `pulumi-gen-svm/main.go`
4. `examples` a folder of Pulumi programs to try locally and/or use in CI.
5. A `Makefile` and this `README`.

#### Additional Details

This repository depends on the pulumi-go-provider library. For more details on building providers, please check
the [Pulumi Go Provider docs](https://github.com/pulumi/pulumi-go-provider).

### Build Examples

Create an example program using the resources defined in your provider, and place it in the `examples/` folder.

You can now repeat the steps for [build, install, and test](#test-against-the-example).

## Configuring CI and releases

1. Follow the instructions laid out in the [deployment templates](./deployment-templates/README-DEPLOYMENT.md).

## References

Other resources/examples for implementing providers:

- [Pulumi Command provider](https://github.com/pulumi/pulumi-command/blob/master/provider/pkg/provider/provider.go)
- [Pulumi Go Provider repository](https://github.com/pulumi/pulumi-go-provider)
  > > > > > > > c9c5061 (Initial commit)
