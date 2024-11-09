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

## Developing against a local `svmkit`

In order to be able to use/test unreleased local changes to `svmkit` it is important that up you update
`provider/go.mod` with a [`replace`](https://go.dev/ref/mod#go-mod-file-replace) directive to point at your local `svmkit` e.g.:

```
--- a/provider/go.mod
+++ b/provider/go.mod
@@ -12,6 +12,8 @@ require (
        github.com/pulumi/pulumi/sdk/v3 v3.138.0
 )
 
+replaces github.com/abklabs/svmkit/pkg => ../../svmkit/pkg
+
 require (
        cloud.google.com/go v0.112.1 // indirect
        cloud.google.com/go/compute v1.25.0 // indirect
```

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

Keep in mind, these examples work using your local build of `pulumi-svmkit` if the `Pulumi.yaml` includes:

```
plugins:
  providers:
    - name: svmkit
      path: ../../bin
```

If you are using an example which doesn't, then you must make sure that `pulumi-svmkit/bin` is in your `PATH` before you run the example.  Otherwise it will attempt to download an unreleased `pulumi-svmkit` which will most likely fail.

