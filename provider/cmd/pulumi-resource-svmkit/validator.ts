import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import type { Input } from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";
import { flags } from "./agave";
import { naming } from "./name";

export interface ValidatorFlags {
    identityKeyPair: string;
    voteAccountKeyPair: string;
    stakeAccountKeyPair?: string;
    entryPoint?: string[];
    knownValidator?: string[];
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
    noWaitForVoteToStartLeader: boolean;
    fullSnapshotIntervalSlots: number;
    expectedGenesisHash?: string;
    fullRpcAPI?: boolean;
    paths: {
        accounts: string;
        ledger: string;
        log: string;
    };
}

export interface ValidatorArgs {
    connection: command.types.input.remote.ConnectionArgs;
    variant?: string;
    flags: ValidatorFlags;
}

export class Validator extends pulumi.ComponentResource {
    constructor(
        name: string,
        args: ValidatorArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Validator", name, args, {
            ...opts,
        });
        const parent = this;

        const tag = naming(name);

        const { connection } = args;
        const targetDir = "svmkit-validator";

        const VALIDATOR_VARIANT = args.variant ?? "agave";

        this.validate(VALIDATOR_VARIANT);

        // This is gross: the assets we need are inside the
        // built component resource archive.  Unfortunately, it
        // doesn't look like there's an easy way to pass the
        // embedded files in one component, to another component
        // that expects files on the filesystem.
        //
        // If there's a better way to do this, please rip this out.

        const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), `${targetDir}-`));

        const assets = [
            "lib.bash",
            "step-runner",
            "validator",
            "validator-builder",
        ];

        assets.forEach((f) =>
            fs.cpSync(
                path.join(__dirname, "assets", f),
                path.join(tempDir, targetDir, f),
            ),
        );

        const archiveName = `${targetDir}.tar.gz`;

        const assetBuilder = new command.local.Command(
            tag("asset", "builder"),
            {
                archivePaths: [path.join(tempDir, targetDir, "*")],
                dir: path.join(tempDir, targetDir),
                create: `bash ./validator-builder ${archiveName}`,
                environment: {
                    VALIDATOR_VARIANT,
                    // Note: these flags are currently "Agave
                    // specific", but all of our validator
                    // variants are either based off of Agave or
                    // the upstream Solana Labs variant.  This can
                    // get changed in future.
                    VALIDATOR_FLAGS: flags(args.flags).join(" "),
                    IDENTITY_KEYPAIR: args.flags.identityKeyPair,
                    VOTEACCOUNT_KEYPAIR: args.flags.voteAccountKeyPair,
                    STAKEACCOUNT_KEYPAIR: args.flags.stakeAccountKeyPair ?? "",
                },
            },
            {
                parent,
            },
        );

        // CopyFile is deprecated, but there are bugs cropping up
        // in CopyToRemote, so use this for now.
        const copyAssets = new command.remote.CopyFile(
            tag("copy", "assets"),
            {
                connection,
                localPath: pulumi.output(path.join(tempDir, archiveName)),
                remotePath: archiveName,
                triggers: [assetBuilder],
            },
            {
                parent,
                dependsOn: [assetBuilder],
            },
        );

        // Remove the temp directory after the archive has been created.
        copyAssets.localPath.apply((_) =>
            fs.rmSync(tempDir, { recursive: true }),
        );

        new command.remote.Command(
            tag("run", "setup"),
            {
                connection,
                create: [
                    `tar xvzf ${archiveName}`,
                    `bash ./${targetDir}/step-runner setup ./${targetDir}/validator`,
                    `rm -rf ${targetDir} ${archiveName}`,
                ].join(" && "),
                environment: {},
                triggers: [copyAssets],
            },
            {
                parent,
                dependsOn: copyAssets,
            },
        );

        this.registerOutputs({});
    }

    private validate(variant: string) {
        const validVariants = [
            "solana",
            "agave",
            "powerledger",
            "jito",
            "pyth",
        ];

        if (!validVariants.includes(variant)) {
            throw new Error(
                `Validator variant '${variant}' is unknown, valid options are: ${validVariants.join(", ")}`,
            );
        }
    }
}

export async function constructValidator(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const backend = new Validator(name, inputs as ValidatorArgs, options);

    return {
        urn: backend.urn,
        state: {},
    };
}
