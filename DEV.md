# Development

This section provides all the necessary information to start contributing to or building SVMKit. It is divided into several subsections for clarity.

## Requirement

Ensure the following tools are installed and available in your `$PATH`:

- [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
- [Go 1.22](https://golang.org/dl/) or 1.latest
- [NodeJS](https://nodejs.org/en/) 14.x (We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations)
- [Yarn](https://yarnpkg.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Python](https://www.python.org/downloads/) (referred to as `python3`; the system-installed version is sufficient for recent MacOS versions)
- [.NET](https://dotnet.microsoft.com/download)

## Repository Structure

The repository is organized as follows:

| Directory/File          | Description                                                                 |
|-------------------------|-----------------------------------------------------------------------------|
| `provider/`             | Contains the build and implementation logic for the SVMkit Pulumi provider. |
| `sdk/`                  | Houses the generated code libraries.                                        |
| `examples/`             | Includes Pulumi programs for local testing and CI usage.                    |
| `Makefile` and `README` | Standard project files for building and documentation.                      |

## Build

```bash
$ make
```

This will build the pulumi provider, generate language sdks, and prepare host to execute the plugin locally.

## Implementation Examples

You can find a catalog of example Pulumi projects to help you get started with SVMkit [here](./examples).

```bash
$ cd examples/aws-agave-validator
$ yarn link @svmkit/pulumi-svmkit
$ yarn install
$ pulumi stack init demo
$ pulumi up
```

In this example, an Agave validator is installed on a machine via SSH, joining the Solana testnet.
