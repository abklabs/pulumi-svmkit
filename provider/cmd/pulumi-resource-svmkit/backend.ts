import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import type { Output } from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";

export interface BackendArgs {
    connection: command.types.input.remote.ConnectionArgs;
    publicDNSName: Output<string>;
    triggers?: command.remote.CopyFileArgs["triggers"];
}

export class Backend extends pulumi.ComponentResource {
    constructor(
        name: string,
        args: BackendArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Backend", name, args, opts);
        const parent = this;

        const genName = (...parts: (string | number)[]) =>
            [name, ...parts].join("-");

        const { connection, triggers } = args;
        const targetDir = "svmkit";
        const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), `${targetDir}-`));
        ["lib.bash", "step-runner", "setup"].forEach((f) =>
            fs.cpSync(
                path.join(__dirname, "assets", f),
                path.join(tempDir, targetDir, f),
            ),
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

        const archive = new pulumi.asset.FileArchive(
            path.join(tempDir, targetDir),
        );
        const copyAssets = new command.remote.CopyToRemote(
            genName("copyAssets"),
            {
                connection,
                source: archive,
                remotePath: ".",
                triggers,
            },
            {
                parent,
                dependsOn: isMachineRunning,
            },
        );

        new command.remote.Command(
            genName("setupInstance"),
            {
                connection,
                create: `bash ./${targetDir}/step-runner setup ./${targetDir}/setup && rm -rf ${targetDir}`,
                environment: {},
                triggers,
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
