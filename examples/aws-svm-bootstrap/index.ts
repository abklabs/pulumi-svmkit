import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@pulumi/svmkit";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";
import * as command from "@pulumi/command";

const config = new pulumi.Config();

const bootstrapSshKey = new tls.PrivateKey("sshKey", {
    algorithm: "ED25519",
});

const keyPair = new aws.ec2.KeyPair("keypair", {
    publicKey: bootstrapSshKey.publicKeyOpenssh,
});

const validatorKey = new svmkit.KeyPair("validatorKey");
const voteAccountKey = new svmkit.KeyPair("voteAccount");
const stakeAccountKey = new svmkit.KeyPair("stakeAccount");
const faucetKey = new svmkit.KeyPair("faucet");

const bootstrapConfig = {
    identityKeyPair: validatorKey.json,
    voteAccountKeyPair: voteAccountKey.json,
    stakeAccountKeyPair: stakeAccountKey.json,
    useSnapshotArchivesAtStartup: "when-newest",
    rpcPort: 8899,
    privateRPC: true,
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
    paths: {
        accounts: "/home/sol/accounts",
        ledger: "/home/sol/ledger",
        log: "/home/sol/log",
    },
};

const amiIdToUse = (async () => {
    let v = config.get("instance.amiId");

    if (v) {
        return v;
    }

    const amiInfo = await aws.ec2.getAmi({
        filters: [
            {
                name: "name",
                values: ["debian-12-*"],
            },
            {
                name: "architecture",
                values: ["x86_64"],
            },
        ],
        owners: ["136693071363"], // Debian
        mostRecent: true,
    });

    return amiInfo.id;
})();

const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {
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
});

const bootstrapInstance = new aws.ec2.Instance("bootstrap", {
    ami: amiIdToUse,
    instanceType: "t3.large",
    keyName: keyPair.keyName,
    vpcSecurityGroupIds: [securityGroup.id],
    rootBlockDevice: {
        volumeSize: 200,
    },
});

const bootstrapSshConnInfo: command.types.input.remote.ConnectionArgs = {
    host: bootstrapInstance.publicDns,
    user: "admin",
    privateKey: bootstrapSshKey.privateKeyOpenssh,
};

const genesis = new svmkit.Genesis("genesis", {
    connection: bootstrapSshConnInfo,
    triggers: [bootstrapInstance.arn],
    genesisConfig: {
        ledgerPath: bootstrapConfig.paths.ledger,
        identityPubkey: validatorKey.publicKey,
        votePubkey: voteAccountKey.publicKey,
        stakePubkey: stakeAccountKey.publicKey,
        faucetPubkey: faucetKey.publicKey,
    },
});

const bootstrapBackendSetup = new svmkit.Backend("bootstrapBackendSetup", {
    connection: bootstrapSshConnInfo,
    triggers: [genesis.urn],
    validatorConfig: bootstrapConfig,
});

export const BOOTSTRAP_PUBLIC_DNS_NAME = bootstrapInstance.publicDns;
export const BOOTSTRAP_SSH_PRIVATE_KEY = bootstrapSshKey.privateKeyOpenssh;
