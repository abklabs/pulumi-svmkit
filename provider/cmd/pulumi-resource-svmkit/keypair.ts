import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";

import * as solanaWeb3 from "@solana/web3.js";

export interface KeyPairArgs {}

export class KeyPair extends pulumi.ComponentResource {
    public readonly publicKey: pulumi.Output<string>;
    public readonly privateKey: pulumi.Output<number[]>;
    public readonly json: pulumi.Output<string>;

    constructor(
        name: string,
        args: KeyPairArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:KeyPair", name, args, opts);

        const keyPair = solanaWeb3.Keypair.generate();
        this.publicKey = pulumi.output(keyPair.publicKey.toBase58());
        this.privateKey = pulumi.secret(Array.from(keyPair.secretKey));

        // Note: the "secretKey" includes private and public parts.
        this.json = pulumi.secret(
            JSON.stringify(Array.from(keyPair.secretKey)),
        );

        this.registerOutputs({
            publicKey: this.publicKey,
            privateKey: this.privateKey,
            json: this.json,
        });
    }
}

export async function constructKeyPair(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const keypair = new KeyPair(name, inputs as KeyPairArgs, options);

    return {
        urn: keypair.urn,
        state: {
            publicKey: keypair.publicKey,
            private: keypair.privateKey,
            json: keypair.json,
        },
    };
}
