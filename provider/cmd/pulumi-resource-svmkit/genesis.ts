import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import type { Input } from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";
import { naming } from "./name";

export interface GenesisFlags {
    ledgerPath: string;
    identityPubkey: string;
    votePubkey: string;
    stakePubkey: string;
    faucetPubkey: string;
    faucetLamports?: string;
    targetLamportsPerSignature?: string;
    inflation?: string;
    lamportsPerByteYear?: string;
    slotPerEpoch?: string;
    clusterType?: string;
}

export interface PrimordialConfig {
    treasuryPubkey: Input<string>;
    treasuryLamports?: string;
    initialValidatorPubkeys: Input<string>[];
    initialValidatorLamports?: string;
}

export type PrimordialEntry = [publicKey: string, lamports: string];

export interface GenesisArgs {
    connection: command.types.input.remote.ConnectionArgs;
    flags: GenesisFlags;
    primordial?: PrimordialEntry[];
}

export class Genesis extends pulumi.ComponentResource {
    public readonly genesisHash: pulumi.Output<string>;

    constructor(
        name: string,
        args: GenesisArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Genesis", name, args, opts);
        const parent = this;

        const tag = naming(name);

        const { connection, flags, primordial } = args;
        const targetDir = "svmkit-genesis";

        const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), `${targetDir}-`));

        const assets = [
            "genesis-builder",
            "genesis",
            "lib.bash",
            "step-runner",
        ];

        assets.forEach((f) =>
            fs.cpSync(
                path.join(__dirname, "assets", f),
                path.join(tempDir, targetDir, f),
            ),
        );

        const archiveName = `${targetDir}.tar.gz`;

        const PRIMORDIAL = primordial ? primordial.flat().join(",") : "";

        const assetBuilder = new command.local.Command(
            tag("asset", "builder"),
            {
                archivePaths: [path.join(tempDir, targetDir, "*")],
                dir: path.join(tempDir, targetDir),
                create: `bash ./genesis-builder ${archiveName}`,
                environment: {
                    LEDGER_PATH: flags.ledgerPath,
                    IDENTITY_PUBKEY: flags.identityPubkey,
                    VOTE_PUBKEY: flags.votePubkey,
                    STAKE_PUBKEY: flags.stakePubkey,
                    FAUCET_PUBKEY: flags.faucetPubkey,
                    FAUCET_LAMPORTS: flags.faucetLamports ?? "1000",
                    TARGET_LAMPORTS_PER_SIGNATURE:
                        flags.targetLamportsPerSignature ?? "0",
                    INFLATION: flags.inflation ?? "none",
                    LAMPORTS_PER_BYTE_YEAR: flags.lamportsPerByteYear ?? "1",
                    SLOT_PER_EPOCH: flags.slotPerEpoch ?? "150",
                    CLUSTER_TYPE: flags.clusterType ?? "development",
                    PRIMORDIAL,
                },
            },
            {
                parent,
            },
        );

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

        copyAssets.localPath.apply((_) =>
            fs.rmSync(tempDir, { recursive: true }),
        );

        const setupGenesis = new command.remote.Command(
            tag("setup", "genesis"),
            {
                connection,
                create: [
                    `tar xvzf ${archiveName}`,
                    `bash ./${targetDir}/step-runner genesis ./${targetDir}/genesis`,
                    `rm -rf ./${targetDir} ${archiveName}`,
                ].join(" && "),
                triggers: [copyAssets.urn],
            },
            {
                parent,
                dependsOn: copyAssets,
            },
        );

        const genesisHashCommand = new command.remote.Command(
            tag("genesis", "hash"),
            {
                connection,
                create: "solana genesis-hash",
                triggers: [copyAssets.urn],
            },
            {
                parent,
                dependsOn: setupGenesis,
            },
        );

        this.genesisHash = genesisHashCommand.stdout;

        this.registerOutputs({
            genesisHash: genesisHashCommand.stdout,
        });
    }
}

export async function constructGenesis(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const genesis = new Genesis(name, inputs as GenesisArgs, options);

    return {
        urn: genesis.urn,
        state: {
            genesisHash: genesis.genesisHash,
        },
    };
}
