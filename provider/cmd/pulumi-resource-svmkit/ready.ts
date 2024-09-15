import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import * as command from "@pulumi/command";
import { naming } from "./name";

export interface ReadyArgs {
    timeout?: string;
    connection: command.types.input.remote.ConnectionArgs;
}

export class Ready extends pulumi.ComponentResource {
    constructor(
        name: string,
        args: ReadyArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Ready", name, args, opts);
        const parent = this;

        const tag = naming(name);

        const { connection } = args;

        const timeout = args.timeout ?? "10m";

        new command.remote.Command(
            tag("is", "machine", "running"),
            {
                connection: { ...connection, dialErrorLimit: -1 },
                create: "echo connected",
            },
            {
                parent,
                customTimeouts: { create: timeout },
            },
        );
    }
}

export async function constructReady(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const ready = new Ready(name, inputs as ReadyArgs, options);

    return {
        urn: ready.urn,
        state: {},
    };
}
