import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@svmkit/pulumi-svmkit";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";

const nodeConfig = new pulumi.Config("node");

const ami = pulumi.output(
    aws.ec2.getAmi({
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
    }),
).id;

const sshKey = new tls.PrivateKey("ssh-key", {
    algorithm: "ED25519",
});

const keyPair = new aws.ec2.KeyPair("keypair", {
    publicKey: sshKey.publicKeyOpenssh,
});

const validatorKey = new svmkit.KeyPair("validator-key");
const voteAccountKey = new svmkit.KeyPair("vote-account-key");

const securityGroup = new aws.ec2.SecurityGroup("security-group", {
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
        {
            protocol: "udp",
            fromPort: 8900,
            toPort: 8920,
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
    ami,
    instanceType: nodeConfig.get("instanceType") ?? "r7a.8xlarge",
    keyName: keyPair.keyName,
    vpcSecurityGroupIds: [securityGroup.id],
    ebsBlockDevices: [
        {
            deviceName: "/dev/sdf",
            volumeSize: 500,
            volumeType: "io2",
            iops: 16000,
        },
        {
            deviceName: "/dev/sdg",
            volumeSize: 1024,
            volumeType: "io2",
            iops: 16000,
        },
    ],
    userData: `#!/bin/bash
mkfs -t ext4 /dev/sdf
mkfs -t ext4 /dev/sdg
mkdir -p /home/sol/accounts
mkdir -p /home/sol/ledger
cat <<EOF >> /etc/fstab
/dev/sdf	/home/sol/accounts	ext4	defaults	0	0
/dev/sdg	/home/sol/ledger	ext4	defaults	0	0
EOF
systemctl daemon-reload
mount -a
`,
});

const connection = {
    host: instance.publicDns,
    user: "admin",
    privateKey: sshKey.privateKeyOpenssh,
};

new svmkit.validator.Firedancer(
    "fd",
    {
        connection,
        keyPairs: {
            identity: validatorKey.json,
            voteAccount: voteAccountKey.json,
        },
        config: {
            user: "sol",
            gossip: {
                host: instance.publicIp,
                entrypoints: [
                    "entrypoint.devnet.solana.com:8001",
                    "entrypoint2.devnet.solana.com:8001",
                    "entrypoint3.devnet.solana.com:8001",
                    "entrypoint4.devnet.solana.com:8001",
                    "entrypoint5.devnet.solana.com:8001",
                ],
            },
            consensus: {
                identityPath: "/home/sol/validator-keypair.json",
                voteAccountPath: "/home/sol/vote-account-keypair.json",
                knownValidators: [
                    "dv1ZAGvdsz5hHLwWXsVnM94hWf1pjbKVau1QVkaMJ92",
                    "dv2eQHeP4RFrJZ6UeiZWoc3XTtmtZCUKxxCApCDcRNV",
                    "dv4ACNkpYPcE3aKmYDqZm9G5EB3J4MRoeE7WNDRBVJB",
                    "dv3qDFk1DTF36Z62bNvrCXe9sKATA6xvVy6A798xxAS",
                ],
            },
            ledger: {
                path: "/home/sol/ledger",
                accountsPath: "/home/sol/accounts",
            },
            rpc: {
                port: 8899,
                fullApi: true,
                private: true,
            },
        },
    },
    {
        dependsOn: [instance],
    },
);

export const username = connection.user;
export const ipAddress = connection.host;
export const sshPrivateKey = connection.privateKey;

export const validatorPublicKey = validatorKey.publicKey;
export const voteAccountPublicKey = voteAccountKey.publicKey;

export const PUBLIC_DNS_NAME = instance.publicDns;
export const SSH_PRIVATE_KEY = sshKey.privateKeyOpenssh;
