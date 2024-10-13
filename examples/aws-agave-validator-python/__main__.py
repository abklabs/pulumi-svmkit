"""An AWS Python Pulumi program"""

import pulumi
import pulumi_aws as aws
import pulumi_tls as tls
import pulumi_svmkit as svmkit

ami = aws.ec2.get_ami(
    filters=[
        {
            "name": "name",
            "values": ["debian-12-*"],
        },
        {
            "name": "architecture",
            "values": ["x86_64"],
        },
    ],
    owners=["136693071363"],  # Debian
    most_recent=True,
).id

ssh_key = tls.PrivateKey("ssh-key", algorithm="ED25519")
key_pair = aws.ec2.KeyPair("keypair", public_key=ssh_key.public_key_openssh)

validator_key = svmkit.KeyPair("validator-key")
vote_account_key = svmkit.KeyPair("vote-account-key")

security_group = aws.ec2.SecurityGroup(
    "security-group",
    description="Allow SSH and specific inbound traffic",
    ingress=[
        {
            "protocol": "tcp",
            "fromPort": 22,
            "toPort": 22,
            "cidrBlocks": ["0.0.0.0/0"],
        },
        {
            "protocol": "tcp",
            "fromPort": 8000,
            "toPort": 8020,
            "cidrBlocks": ["0.0.0.0/0"],
        },
        {
            "protocol": "udp",
            "fromPort": 8000,
            "toPort": 8020,
            "cidrBlocks": ["0.0.0.0/0"],
        }
    ],
    egress=[
        {
            "protocol": "-1",
            "fromPort": 0,
            "toPort": 0,
            "cidrBlocks": ["0.0.0.0/0"],
        }
    ]
)

instance = aws.ec2.Instance(
    "instance",
    ami=ami,
    instance_type="m5.2xlarge",
    key_name=key_pair.key_name,
    vpc_security_group_ids=[security_group.id],
    ebs_block_devices=[
        {
            "deviceName": "/dev/sdf",
            "volumeSize": 500,
            "volumeType": "io2",
            "iops": 16000,
        },
        {
            "deviceName": "/dev/sdg",
            "volumeSize": 1024,
            "volumeType": "io2",
            "iops": 16000,
        },
    ],
    user_data="""#!/bin/bash
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
"""
)

connection = {
    "host": instance.public_dns,
    "user": "admin",
    "privateKey": ssh_key.private_key_openssh,
}

svmkit.validator.Agave(
    "validator",
    connection=connection,
    version="1.18.24-1",
    key_pairs={
        "identity": validator_key.json,
        "voteAccount": vote_account_key.json,
    },
    flags={
        "entryPoint": [
            "entrypoint.testnet.solana.com:8001",
            "entrypoint2.testnet.solana.com:8001",
            "entrypoint3.testnet.solana.com:8001",
        ],
        "knownValidator": [
            "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
            "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
            "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
            "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
        ],
        "expectedGenesisHash": "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
        "useSnapshotArchivesAtStartup": "when-newest",
        "rpcPort": 8899,
        "privateRPC": True,
        "onlyKnownRPC": True,
        "dynamicPortRange": "8002-8020",
        "gossipPort": 8001,
        "rpcBindAddress": "0.0.0.0",
        "walRecoveryMode": "skip_any_corrupted_record",
        "limitLedgerSize": 50000000,
        "blockProductionMethod": "central-scheduler",
        "fullSnapshotIntervalSlots": 1000,
        "noWaitForVoteToStartLeader": True,
    },
    opts=pulumi.ResourceOptions(depends_on=[instance])
)

pulumi.export("PUBLIC_DNS_NAME", instance.public_dns)
pulumi.export("SSH_PRIVATE_KEY", ssh_key.private_key_openssh)
