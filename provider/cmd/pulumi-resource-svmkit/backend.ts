import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import type { Input } from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";
import { generateAgaveValidatorFlags } from "./agave";

export interface BackendValidatorConfig {
    validatorVariant?: string;
    identityKeyPair: Input<string>;
    voteAccountKeyPair: Input<string>;
    entryPoint: string[];
    knownValidator: string[];
    useSnapshotArchivesAtStartup: string;
    rpcPort: number;
    privateRPC: boolean;
    onlyKnownRPC: boolean;
    dynamicPortRange: string;
    gossipPort: number;
    rpcBindAddress: string;
    walRecoveryMode: string;
    limitLedgerSize: number;
    blockProductionMethod: string;
    tvuReceiveThreads?: number;
    paths: {
        accounts: string;
        ledger: string;
        log: string;
    };
}

export interface BackendArgs {
    connection: command.types.input.remote.ConnectionArgs;
    triggers?: command.remote.CopyFileArgs["triggers"];
    validatorConfig: BackendValidatorConfig;
    dependsOn?: pulumi.Input<pulumi.Resource>[];
}

export class Backend extends pulumi.ComponentResource {
    constructor(
        name: string,
        args: BackendArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Backend", name, args, {
            ...opts,
            dependsOn: args.dependsOn,
        });
        const parent = this;

        const genName = (...parts: (string | number)[]) =>
            [name, ...parts].join("-");

        const { connection, triggers } = args;
        const targetDir = "svmkit";

        const VALIDATOR_VARIANT =
            args.validatorConfig.validatorVariant ?? "agave";

        // Note: this list must match the names defined inside solana-build.
        const validVariants = [
            "solana",
            "agave",
            "powerledger",
            "jito",
            "pyth",
        ];

        if (!validVariants.includes(VALIDATOR_VARIANT)) {
            throw new Error(
                `Validator variant '${VALIDATOR_VARIANT}' is unknown, valid options are: ${validVariants.join(", ")}`,
            );
        }

        // This is gross: the assets we need are inside the
        // built component resource archive.  Unfortunately, it
        // doesn't look like there's an easy way to pass the
        // embedded files in one component, to another component
        // that expects files on the filesystem.
        //
        // If there's a better way to do this, please rip this out.

        const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), `${targetDir}-`));

        const assets = ["lib.bash", "step-runner", "setup", "asset-builder"];

        assets.forEach((f) =>
            fs.cpSync(
                path.join(__dirname, "assets", f),
                path.join(tempDir, targetDir, f),
            ),
        );

        const archiveName = `${targetDir}.tar.gz`;

        const assetBuilder = new command.local.Command(
            genName("assetBuilder"),
            {
                archivePaths: [path.join(tempDir, targetDir, "*")],
                dir: path.join(tempDir, targetDir),
                create: `bash ./asset-builder ${archiveName}`,
                environment: {
                    VALIDATOR_VARIANT,
                    // Note: these flags are currently "Agave
                    // specific", but all of our validator
                    // variants are either based off of Agave or
                    // the upstream Solana Labs variant.  This can
                    // get changed in future.
                    VALIDATOR_FLAGS: generateAgaveValidatorFlags(
                        args.validatorConfig,
                    ).join(" "),
                    IDENTITY_KEYPAIR: args.validatorConfig.identityKeyPair,
                    VOTEACCOUNT_KEYPAIR:
                        args.validatorConfig.voteAccountKeyPair,
                },
            },
            {
                parent,
            },
        );

        const isMachineRunning = new command.remote.Command(
            genName("isMachineRunning"),
            {
                connection: { ...connection, dialErrorLimit: -1 },
                create: "echo connected",
            },
            {
                parent,
                customTimeouts: { create: "10m" },
            },
        );

        // CopyFile is deprecated, but there are bugs cropping up
        // in CopyToRemote, so use this for now.
        const copyAssets = new command.remote.CopyFile(
            genName("copyAssets"),
            {
                connection,
                localPath: pulumi.output(path.join(tempDir, archiveName)),
                remotePath: archiveName,
                triggers: [triggers, assetBuilder],
            },
            {
                parent,
                dependsOn: [assetBuilder, isMachineRunning],
            },
        );

        // Remove the temp directory after the archive has been created.
        copyAssets.localPath.apply((_) =>
            fs.rmSync(tempDir, { recursive: true }),
        );

        new command.remote.Command(
            genName("setupInstance"),
            {
                connection,
                create: `tar xvzf ${archiveName} && bash ./${targetDir}/step-runner setup ./${targetDir}/setup && rm -rf ${targetDir} ${archiveName}`,
                environment: {},
                triggers: [triggers, copyAssets],
            },
            {
                parent,
                dependsOn: copyAssets,
            },
        );

        this.registerOutputs({});
    }
}

export async function constructBackend(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const backend = new Backend(name, inputs as BackendArgs, options);

    return {
        urn: backend.urn,
        state: {},
    };
}
