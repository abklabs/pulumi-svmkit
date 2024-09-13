import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";
import {
    Connection,
    Keypair,
    VoteProgram,
    Transaction,
    sendAndConfirmTransaction,
    TransactionSignature,
} from "@solana/web3.js";

// The parameters for the vote transaction.
export type VoteTransactionParams = {
    // The reward commissions for the validator. Defaults to 0.
    commission?: number;
    // The authorized withdrawer is the entity that can withdraw rewards from the vote account.
    authorizedWithdrawer: string;
    // The vote account keypair to register with the network.
    voteAccount: string;
    // The keypair for the validator identity.
    identity: string;
};

export interface VoteArgs {
    // The RPC endpoint to submit the transaction.
    endpoint: string;
    // The parameters for the vote transaction.
    params: VoteTransactionParams;
}

// Vote is a Pulumi component resource that creates a vote account on a SVM network.
export class Vote extends pulumi.ComponentResource {
    public readonly signature: Promise<pulumi.Output<TransactionSignature>>;

    constructor(
        name: string,
        args: VoteArgs,
        opts?: pulumi.ComponentResourceOptions,
    ) {
        super("svmkit:index:Vote", name, args, opts);

        this.signature = Vote.create(args.endpoint, args.params);

        this.registerOutputs({
            signature: this.signature,
        });
    }

    public static async create(
        endpoint: string,
        params: VoteTransactionParams,
    ): Promise<pulumi.Output<TransactionSignature>> {
        const conn = new Connection(endpoint, "confirmed");

        const authorizedWithdrawerKeypair = Keypair.fromSecretKey(
            Buffer.from(JSON.parse(params.authorizedWithdrawer)),
        );
        const voteAccountKeypair = Keypair.fromSecretKey(
            Buffer.from(JSON.parse(params.voteAccount)),
        );
        const validatorKeypair = Keypair.fromSecretKey(
            Buffer.from(JSON.parse(params.identity)),
        );

        const instruction = VoteProgram.initializeAccount({
            voteInit: {
                nodePubkey: validatorKeypair.publicKey,
                authorizedVoter: validatorKeypair.publicKey,
                authorizedWithdrawer: authorizedWithdrawerKeypair.publicKey,
                commission: params.commission ?? 0,
            },
            votePubkey: voteAccountKeypair.publicKey,
            nodePubkey: validatorKeypair.publicKey,
        });

        const transaction = new Transaction().add(instruction);

        const signatue = await sendAndConfirmTransaction(
            conn,
            transaction,
            [validatorKeypair],
            { skipPreflight: true },
        );

        return pulumi.output(signatue);
    }
}

export async function constructVote(
    name: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions,
): Promise<provider.ConstructResult> {
    const args = inputs as VoteArgs;
    const vote = new Vote(name, args, options);

    return {
        urn: vote.urn,
        state: {
            signature: vote.signature,
        },
    };
}
