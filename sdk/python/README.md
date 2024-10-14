# Pulumi SVMKit

The SVMKit Pulumi provider, located in the [`provider`](/provider/) directory, facilitates infrastructure as code (IaC) for managing Solana nodes and related resources. With this provider, users can easily define, deploy, and manage their Solana infrastructure using Pulumi.

## Development

This section provides all the necessary information to start contributing to or building SVMKit. It is divided into several subsections for clarity.

#### Requirement

Ensure the following tools are installed and available in your `$PATH`:

- [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
- [Go 1.22](https://golang.org/dl/) or 1.latest
- [NodeJS](https://nodejs.org/en/) 14.x (We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations)
- [Yarn](https://yarnpkg.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Python](https://www.python.org/downloads/) (referred to as `python3`; the system-installed version is sufficient for recent MacOS versions)
- [.NET](https://dotnet.microsoft.com/download)

#### Build

```bash
$ make build install
```

This will build the pulumi provider, generate language sdks, and prepare host to execute the plugin locally.

#### Demo

You can find a catalog of example Pulumi projects to help you get started with SVMkit [here](./examples).

```bash
$ cd examples/aws-agave-validator
$ yarn link @svmkit/pulumi-svmkit
$ yarn install
$ pulumi stack init demo
$ pulumi up
```

In this example, an Agave validator is installed on a machine via SSH, joining the Solana testnet.

Teams can add more validator clients to SVMkit, which will be accessible through the `validator` namespace in `@svmkit/pulumi-svmkit`.

```typescript
new svmkit.validator.Agave(
  "validator",
  {
    connection,
    keyPairs: {
      identity: validatorKey.json,
      voteAccount: voteAccountKey.json,
    },
    flags: {
      entryPoint: [
        "entrypoint.testnet.solana.com:8001",
        "entrypoint2.testnet.solana.com:8001",
        "entrypoint3.testnet.solana.com:8001",
      ],
      knownValidator: [
        "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
        "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
        "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
        "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
      ],
      expectedGenesisHash: "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
      useSnapshotArchivesAtStartup: "when-newest",
      rpcPort: 8899,
      privateRPC: true,
      onlyKnownRPC: true,
      dynamicPortRange: "8002-8020",
      gossipPort: 8001,
      rpcBindAddress: "0.0.0.0",
      walRecoveryMode: "skip_any_corrupted_record",
      limitLedgerSize: 50000000,
      blockProductionMethod: "central-scheduler",
      fullSnapshotIntervalSlots: 1000,
      noWaitForVoteToStartLeader: true,
    },
  },
  {
    dependsOn: [instance],
  }
);
```

### Repository Structure

The repository is organized as follows:

| Directory/File          | Description                                                                 |
|-------------------------|-----------------------------------------------------------------------------|
| `provider/`             | Contains the build and implementation logic for the SVMkit Pulumi provider. |
| `sdk/`                  | Houses the generated code libraries.                                        |
| `examples/`             | Includes Pulumi programs for local testing and CI usage.                    |
| `Makefile` and `README` | Standard project files for building and documentation.                      |
