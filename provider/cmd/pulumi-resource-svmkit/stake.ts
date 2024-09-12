import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import type { Input } from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";

import { naming } from "./name";

export interface StakeConfig {
    amount: number;
}

export interface StakeArgs {
    connection: command.types.input.remote.ConnectionArgs;
    stakeConfig: StakeConfig;
}

export class Stake extends pulumi.ComponentResource {
    constructor(
        name: string,
        args: StakeArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Stake", name, args, opts);
        const parent = this;

        const tag = naming(name);

        const { connection, stakeConfig } = args;
        const targetDir = "svmkit-stake";

        const isMachineRunning = new command.remote.Command(
            tag("isMachineRunning"),
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

        const assets = ["stake-builder", "stake", "lib.bash", "step-runner"];

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
                create: `bash ./stake-builder ${archiveName}`,
                environment: {
                    STAKE_AMOUNT: stakeConfig.amount.toString(),
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
            tag("distribute", "stake"),
            {
                connection,
                create: `tar xvzf ${archiveName} && \
                  bash ./${targetDir}/step-runner stake ./${targetDir}/stake && \
                  rm -rf ./${targetDir} ${archiveName}`,
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

export async function constructStake(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const genesis = new Stake(name, inputs as StakeArgs, options);

    return {
        urn: genesis.urn,
        state: {},
    };
}
