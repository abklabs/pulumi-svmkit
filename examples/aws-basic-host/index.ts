import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@pulumi/svmkit";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";
import * as command from "@pulumi/command";

const config = new pulumi.Config();

const sshKey = new tls.PrivateKey("sshKey", {
    algorithm: "ED25519",
});

const keyPair = new aws.ec2.KeyPair("keypair", {
    publicKey: sshKey.publicKeyOpenssh,
});

const validatorKey = new svmkit.KeyPair("validatorKey");
const voteAccountKey = new svmkit.KeyPair("voteAccount");

const validatorConfig = {
    identityKeyPair: validatorKey.json,
    voteAccountKeyPair: voteAccountKey.json,
    entryPoint: [
        "entrypoint.testnet.solana.com:8001",
        "entrypoint2.testnet.solana.com:8001",
        "entrypoint3.testnet.solana.com:8001",
    ],
    knownValidator: [
        "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
        "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
        "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
        "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
    ],
    expectedGenesisHash: "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
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

const instance = new aws.ec2.Instance("instance", {
    ami: amiIdToUse,
    instanceType: "t3.large",
    keyName: keyPair.keyName,
    vpcSecurityGroupIds: [securityGroup.id],
    rootBlockDevice: {
        volumeSize: 200,
    },
});

const sshConnInfo: command.types.input.remote.ConnectionArgs = {
    host: instance.publicDns,
    user: "admin",
    privateKey: sshKey.privateKeyOpenssh,
};

const backendSetup = new svmkit.Backend("backendSetup", {
    connection: sshConnInfo,
    triggers: [instance.arn],
    validatorConfig,
});

export const PUBLIC_DNS_NAME = instance.publicDns;
export const SSH_PRIVATE_KEY = sshKey.privateKeyOpenssh;
