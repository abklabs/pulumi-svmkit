import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@pulumi/svmkit";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";
import * as command from "@pulumi/command";

const config = new pulumi.Config("svmkit");

const ami = config.require("ami");

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

const paths = {
    accounts: "/home/sol/accounts",
    ledger: "/home/sol/ledger",
    log: "/home/sol/log",
};

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

const instance = new aws.ec2.Instance("instance", {
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
    host: instance.publicDns,
    user,
    privateKey,
};

const ready = new svmkit.Ready(
    "instance-ready",
    {
        connection,
    },
    { dependsOn: [instance] },
);

const genesis = pulumi
    .all([
        bootstrapIdentityKey.publicKey,
        bootstrapVoteAccountKey.publicKey,
        bootstrapStakeAccountKey.publicKey,
        faucetKey.publicKey,
        treasuryKey.publicKey,
    ])
    .apply(
        ([
            identityPubkey,
            votePubkey,
            stakePubkey,
            faucetPubkey,
            treasuryPubkey,
        ]) => {
            const primordial = [
                [treasuryPubkey, "1000000000000"], // 1000 SOL
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
                { dependsOn: [instance, ready] },
            );
        },
    );

new svmkit.Validator(
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

export const PUBLIC_DNS_NAME = instance.publicDns;
export const SSH_PRIVATE_KEY = bootstrapSshKey.privateKeyOpenssh;
