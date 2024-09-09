import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import type { Input } from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";

export interface GenesisConfig {
    ledgerPath: string;
    identityPubkey: Input<string>;
    votePubkey: Input<string>;
    stakePubkey: Input<string>;
    faucetPubkey: Input<string>;
    faucetLamports?: string;
    targetLamportsPerSignature?: string;
    inflation?: string;
    lamportsPerByteYear?: string;
    slotPerEpoch?: string;
    clusterType?: string;
}

export interface GenesisArgs {
    connection: command.types.input.remote.ConnectionArgs;
    triggers?: command.remote.CopyFileArgs["triggers"];
    genesisConfig: GenesisConfig;
}

export class Genesis extends pulumi.ComponentResource {
    constructor(
        name: string,
        args: GenesisArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Genesis", name, args, opts);
        const parent = this;

        const genName = (...parts: (string | number)[]) =>
            [name, ...parts].join("-");

        const { connection, triggers, genesisConfig } = args;
        const targetDir = "svmkit-genesis";

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

        const assetBuilder = new command.local.Command(
            genName("assetBuilder"),
            {
                archivePaths: [path.join(tempDir, targetDir, "*")],
                dir: path.join(tempDir, targetDir),
                create: `bash ./genesis-builder ${archiveName}`,
                environment: {
                    LEDGER_PATH: genesisConfig.ledgerPath,
                    IDENTITY_PUBKEY: genesisConfig.identityPubkey,
                    VOTE_PUBKEY: genesisConfig.votePubkey,
                    STAKE_PUBKEY: genesisConfig.stakePubkey,
                    FAUCET_PUBKEY: genesisConfig.faucetPubkey,
                    FAUCET_LAMPORTS: genesisConfig.faucetLamports ?? "1000",
                    TARGET_LAMPORTS_PER_SIGNATURE:
                        genesisConfig.targetLamportsPerSignature ?? "0",
                    INFLATION: genesisConfig.inflation ?? "none",
                    LAMPORTS_PER_BYTE_YEAR:
                        genesisConfig.lamportsPerByteYear ?? "1",
                    SLOT_PER_EPOCH: genesisConfig.slotPerEpoch ?? "150",
                    CLUSTER_TYPE: genesisConfig.clusterType ?? "development",
                },
            },
            {
                parent,
            },
        );

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

        copyAssets.localPath.apply((_) =>
            fs.rmSync(tempDir, { recursive: true }),
        );
        new command.remote.Command(
            genName("setupGenesis"),
            {
                connection,
                create: pulumi.interpolate`
                  tar xvzf ${archiveName} && \
                  bash ./${targetDir}/step-runner genesis ./${targetDir}/genesis && \
                  rm -rf ./${targetDir} ${archiveName}
              `,
                triggers: [copyAssets.urn],
            },
            {
                parent,
                dependsOn: copyAssets,
            },
        );

        this.registerOutputs({});
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
        state: {},
    };
}
