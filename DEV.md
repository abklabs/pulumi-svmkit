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
| `Makefile` and `README` | Standard project files for building and documentation.                      |

## Build

```bash
$ make
```

This will build the pulumi provider, generate language sdks, and prepare host to execute the plugin locally.

You can use the locally built pulumi provider by including `.../pulumi-svmkit/bin` in your `$PATH`.  It is
recommended that you do this on a temporary basis, in the environments where you need to test your local
pulumi-svmkit provider build.  DO NOT include this in your permanent `$PATH`.

## Developing against a local `svmkit`

In order to be able to use/test unreleased local changes to `svmkit` it is important that up you update
`provider/go.mod` with a [`replace`](https://go.dev/ref/mod#go-mod-file-replace) directive to point at your local `svmkit` e.g.:

```
--- a/provider/go.mod
+++ b/provider/go.mod
@@ -12,6 +12,8 @@ require (
        github.com/pulumi/pulumi/sdk/v3 v3.138.0
 )

+replace github.com/abklabs/svmkit/pkg => ../../svmkit/pkg
+
 require (
        cloud.google.com/go v0.112.1 // indirect
        cloud.google.com/go/compute v1.25.0 // indirect
```

You can perform this by running:

```
( cd provider && go mod edit -replace github.com/abklabs/svmkit/pkg=../../svmkit/pkg )
```

### pulumi-resource-svmkit

The pulumi-resource-svmkit is a custom Pulumi provider binary that enables Pulumi to manage svmkit resources, such as machines and firewalls, by exposing Go-based logic to Pulumi SDKs in languages like Python and TypeScript. Built as a gRPC server using Pulumi’s Go provider SDK, it handles instructions from pulumi up, maps inputs to resource creation logic, executes CRUD operations defined in the pkg/ directory, and returns outputs for use in your Pulumi stack.

### Using a local python sdk
To use a locally-built python svmkit sdk for testing, you can run the following in the root of your project:
```
pulumi install
./venv/bin/pip install ../pulumi-svmkit/sdk/python/
```
