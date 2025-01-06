"""An AWS Python Pulumi program"""

import pulumi
import pulumi_aws as aws
import pulumi_tls as tls
import pulumi_svmkit as svmkit

# replace this with the actual RPC host and port
RPC_HOST = "localhost"
RPC_PORT = 8899

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
        "Name": "svmkit-explorer",
    },
)

sol_env = svmkit.solana.EnvironmentArgs(
    rpc_url=f"http://{RPC_HOST}:{RPC_PORT}"
)


svmkit.explorer.Explorer(
    "explorer",
    connection={
        "host": instance.public_dns,
        "user": "admin",
        "private_key": ssh_key.private_key_openssh
    },
    environment=sol_env,
    flags={
        "hostname": "0.0.0.0",
        "port": 3000,
    },
    opts=pulumi.ResourceOptions(depends_on=[instance])
)

pulumi.export("nodes_name", ["instance"])
pulumi.export("nodes_public_ip", [instance.public_ip])
pulumi.export("nodes_private_key", [ssh_key.private_key_openssh])
