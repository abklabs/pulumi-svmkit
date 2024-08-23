import * as path from "node:path";
import * as fs from "node:fs";
import * as os from "node:os";

import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";

export interface BackendArgs {
    connection: command.types.input.remote.ConnectionArgs;
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

        // This is gross: the assets we need are inside the
        // built component resource archive.  Unfortunately, it
        // doesn't look like there's an easy way to pass the
        // embedded files in one component, to another component
        // that expects files on the filesystem.
        //
        // If there's a better way to do this, please rip this out.

        const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), `${targetDir}-`));

        ["lib.bash", "step-runner", "setup"].forEach((f) =>
            fs.cpSync(
                path.join(__dirname, "assets", f),
                path.join(tempDir, targetDir, f),
            ),
        );

        const archive = new pulumi.asset.FileArchive(
            path.join(tempDir, targetDir),
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

        // Remove the temp directory after the archive has been created.
        copyAssets.source.apply((_) => fs.rmSync(tempDir, { recursive: true }));

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
