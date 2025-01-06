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
            "from_port": 22,
            "to_port": 22,
            "cidr_blocks": ["0.0.0.0/0"],
        },
        {
            "protocol": "tcp",
            "from_port": 8000,
            "to_port": 8020,
            "cidr_blocks": ["0.0.0.0/0"],
        },
        {
            "protocol": "udp",
            "from_port": 8000,
            "to_port": 8020,
            "cidr_blocks": ["0.0.0.0/0"],
        }
    ],
    egress=[
        {
            "protocol": "-1",
            "from_port": 0,
            "to_port": 0,
            "cidr_blocks": ["0.0.0.0/0"],
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
            "device_name": "/dev/sdf",
            "volume_size": 500,
            "volume_type": "io2",
            "iops": 16000,
        },
        {
            "device_name": "/dev/sdg",
            "volume_size": 1024,
            "volume_type": "io2",
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

svmkit.validator.Agave(
    "validator",
    connection={
        "host": instance.public_dns,
        "user": "admin",
        "private_key": ssh_key.private_key_openssh
    },
    version="1.18.24-1",
    key_pairs={
        "identity": validator_key.json,
        "vote_account": vote_account_key.json,
    },
    flags={
        "entry_point": [
            "entrypoint.testnet.solana.com:8001",
            "entrypoint2.testnet.solana.com:8001",
            "entrypoint3.testnet.solana.com:8001",
        ],
        "known_validator": [
            "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
            "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
            "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
            "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
        ],
        "expected_genesis_hash": "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
        "use_snapshot_archives_at_startup": "when-newest",
        "rpc_port": 8899,
        "private_rpc": True,
        "only_known_rpc": True,
        "dynamic_port_range": "8002-8020",
        "gossip_port": 8001,
        "rpc_bind_address": "0.0.0.0",
        "wal_recovery_mode": "skip_any_corrupted_record",
        "limit_ledger_size": 50000000,
        "block_production_method": "central-scheduler",
        "full_snapshot_interval_slots": 1000,
        "no_wait_for_vote_to_start_leader": True,
    },
    opts=pulumi.ResourceOptions(depends_on=[instance])
)
pulumi.export("nodes_name", ["instance"])
pulumi.export("nodes_public_ip", [instance.public_ip])
pulumi.export("nodes_private_key", [ssh_key.private_key_openssh])
