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

faucet_key = svmkit.KeyPair("faucet-key")

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
            "from_port": 9900,
            "to_port": 9900,
            "cidr_blocks": ["0.0.0.0/0"],
        },
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
    instance_type="t3.xlarge",
    key_name=key_pair.key_name,
    vpc_security_group_ids=[security_group.id],
    tags={
        "Name": "svmkit-faucet",
    },
)

svmkit.faucet.Faucet(
    "faucet",
    connection={
        "host": instance.public_dns,
        "user": "admin",
        "private_key": ssh_key.private_key_openssh
    },
    version="1.18.24-1",
    keypair=faucet_key.json,
    flags={
        "per_request_cap": 10,
    },
    opts=pulumi.ResourceOptions(depends_on=[instance])
)

pulumi.export("nodes_name", ["instance"])
pulumi.export("nodes_public_ip", [instance.public_ip])
pulumi.export("nodes_private_key", [ssh_key.private_key_openssh])
