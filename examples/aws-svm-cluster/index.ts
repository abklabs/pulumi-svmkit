import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@pulumi/svmkit";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";
import * as command from "@pulumi/command";

const config = new pulumi.Config("svmkit");

const ami = config.require("ami");
const numValidators = config.getNumber("validators") || 3;

const bootstrapSshKey = new tls.PrivateKey("ssh-key", {
    algorithm: "ED25519",
});

const keyPair = new aws.ec2.KeyPair("keypair", {
    publicKey: bootstrapSshKey.publicKeyOpenssh,
});

// Faucet for receiving SOL
const faucetKey = new svmkit.KeyPair("faucet-key");

// Treasury used to distribute stake
const treasuryKey = new svmkit.KeyPair("treasury-key");

// Bootstrap node
const bootstrapIdentityKey = new svmkit.KeyPair("identity-key-bootstrap");
const bootstrapVoteAccountKey = new svmkit.KeyPair(
    "vote-account-key-bootstrap",
);
const bootstrapStakeAccountKey = new svmkit.KeyPair(
    "stake-account-key-bootstrap",
);

// Validator nodes
const validatorKeys = Array.from({ length: numValidators }, (_, i) => ({
    identity: new svmkit.KeyPair(`identity-key-${i}`),
    voteAccount: new svmkit.KeyPair(`vote-account-key-${i}`),
    stakeAccount: new svmkit.KeyPair(`stake-account-key-${i}`),
}));

const paths = {
    accounts: "/home/sol/accounts",
    ledger: "/home/sol/ledger",
    log: "/home/sol/log",
};

const validatorSecurityGroup = new aws.ec2.SecurityGroup(
    "security-group-validator",
    {
        description: "Allow SSH and specific inbound traffic",
        ingress: [
            {
                protocol: "tcp",
                fromPort: 22,
                toPort: 22,
                cidrBlocks: ["0.0.0.0/0"],
            },
            {
                protocol: "tcp",
                fromPort: 8000,
                toPort: 8020,
                cidrBlocks: ["0.0.0.0/0"],
            },
            {
                protocol: "udp",
                fromPort: 8000,
                toPort: 8020,
                cidrBlocks: ["0.0.0.0/0"],
            },
        ],
        egress: [
            {
                protocol: "-1",
                fromPort: 0,
                toPort: 0,
                cidrBlocks: ["0.0.0.0/0"],
            },
        ],
    },
);

const bootstrapSecurityGroup = new aws.ec2.SecurityGroup(
    "security-group-bootstrap",
    {
        description:
            "Allow SSH and specific inbound traffic. Open RPC port to the world.",
        ingress: [
            {
                protocol: "tcp",
                fromPort: 22,
                toPort: 22,
                cidrBlocks: ["0.0.0.0/0"],
            },
            {
                protocol: "tcp",
                fromPort: 8000,
                toPort: 8020,
                cidrBlocks: ["0.0.0.0/0"],
            },
            {
                protocol: "tcp",
                fromPort: 8899,
                toPort: 8899,
                cidrBlocks: ["0.0.0.0/0"],
            },
            {
                protocol: "udp",
                fromPort: 8000,
                toPort: 8020,
                cidrBlocks: ["0.0.0.0/0"],
            },
        ],
        egress: [
            {
                protocol: "-1",
                fromPort: 0,
                toPort: 0,
                cidrBlocks: ["0.0.0.0/0"],
            },
        ],
    },
);

const bootstrapInstance = new aws.ec2.Instance("instance-bootstrap", {
    ami,
    instanceType: "t3.large",
    keyName: keyPair.keyName,
    vpcSecurityGroupIds: [bootstrapSecurityGroup.id],
    rootBlockDevice: {
        volumeSize: 200,
    },
});

const privateKey = bootstrapSshKey.privateKeyOpenssh;
const user = "admin";

const connection: command.types.input.remote.ConnectionArgs = {
    host: bootstrapInstance.publicDns,
    user,
    privateKey,
};

const bootstrapReady = new svmkit.Ready("instance-bootstrap-ready", {
    connection,
});

const genesis = pulumi
    .all([
        bootstrapIdentityKey.publicKey,
        bootstrapVoteAccountKey.publicKey,
        bootstrapStakeAccountKey.publicKey,
        faucetKey.publicKey,
        treasuryKey.publicKey,
        ...validatorKeys.map((keys) => keys.identity.publicKey),
    ])
    .apply(
        ([
            identityPubkey,
            votePubkey,
            stakePubkey,
            faucetPubkey,
            treasuryPubkey,
            ...validatorPubkeys
        ]) => {
            const primordial = [
                [treasuryPubkey, "1000000000000"], // 1000 SOL
                ...validatorPubkeys.map((pubkey) => [
                    pubkey,
                    "100000000000", // 100 SOL
                ]),
            ];

            return new svmkit.Genesis(
                "genesis",
                {
                    connection,
                    flags: {
                        ledgerPath: paths.ledger,
                        identityPubkey,
                        votePubkey,
                        stakePubkey,
                        faucetPubkey,
                    },
                    primordial,
                },
                { dependsOn: [bootstrapInstance, bootstrapReady] },
            );
        },
    );

const bootstrapValidator = new svmkit.Validator(
    "validator-bootstrap",
    {
        connection,
        flags: {
            identityKeyPair: bootstrapIdentityKey.json,
            voteAccountKeyPair: bootstrapVoteAccountKey.json,
            stakeAccountKeyPair: bootstrapStakeAccountKey.json,
            useSnapshotArchivesAtStartup: "when-newest",
            rpcPort: 8899,
            privateRPC: false,
            onlyKnownRPC: false,
            dynamicPortRange: "8002-8020",
            gossipPort: 8001,
            rpcBindAddress: "0.0.0.0",
            walRecoveryMode: "skip_any_corrupted_record",
            limitLedgerSize: 50000000,
            blockProductionMethod: "central-scheduler",
            tvuReceiveThreads: 2,
            noWaitForVoteToStartLeader: true,
            fullSnapshotIntervalSlots: 1000,
            paths,
        },
    },
    { dependsOn: genesis },
);

export const VALIDATOR_PUBLIC_DNS_NAMES = pulumi
    .all([
        bootstrapInstance.publicIp,
        genesis.genesisHash,
        bootstrapIdentityKey.publicKey,
    ])
    .apply(
        ([publicIp, expectedGenesisHash, bootstrapPubkey]: [
            string,
            string,
            string,
        ]) => {
            const entrypoint = `${publicIp}:8001`;
            const flags = {
                entryPoint: [entrypoint],
                knownValidator: [bootstrapPubkey],
                expectedGenesisHash,
                useSnapshotArchivesAtStartup: "when-newest",
                rpcPort: 8899,
                privateRPC: true,
                onlyKnownRPC: true,
                dynamicPortRange: "8002-8020",
                gossipPort: 8001,
                rpcBindAddress: "0.0.0.0",
                walRecoveryMode: "skip_any_corrupted_record",
                limitLedgerSize: 50000000,
                blockProductionMethod: "central-scheduler",
                tvuReceiveThreads: 2,
                paths,
            };

            const validatorInstances = validatorKeys.map((keys, i) => {
                const instance = new aws.ec2.Instance(`instance-${i}`, {
                    ami,
                    instanceType: "t3.large",
                    keyName: keyPair.keyName,
                    vpcSecurityGroupIds: [validatorSecurityGroup.id],
                    rootBlockDevice: {
                        volumeSize: 200,
                    },
                });

                const connection: command.types.input.remote.ConnectionArgs = {
                    host: instance.publicDns,
                    user,
                    privateKey,
                };

                const ready = new svmkit.Ready(`ready-${i}`, {
                    connection,
                });

                new svmkit.Validator(
                    `validator-${i}`,
                    {
                        connection,
                        flags: {
                            ...flags,
                            identityKeyPair: keys.identity.json,
                            voteAccountKeyPair: keys.voteAccount.json,
                        },
                    },
                    { dependsOn: [instance, bootstrapValidator, ready] },
                );

                return instance;
            });

            return pulumi.output(
                validatorInstances.map((instance) => instance.publicDns),
            );
        },
    );

export const SSH_PRIVATE_KEY = bootstrapSshKey.privateKeyOpenssh;
export const BOOTSTRAP_PUBLIC_DNS_NAME = bootstrapInstance.publicDns;
