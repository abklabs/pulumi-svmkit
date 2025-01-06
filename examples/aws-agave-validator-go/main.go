package main

import (
	"github.com/abklabs/pulumi-svmkit/sdk/go"
	"github.com/abklabs/pulumi-svmkit/sdk/go/agave"
	"github.com/abklabs/pulumi-svmkit/sdk/go/ssh"
	"github.com/abklabs/pulumi-svmkit/sdk/go/validator"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		sshKey, err := tls.NewPrivateKey(ctx, "ssh-key", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("ED25519"),
		})

		if err != nil {
			return err
		}

		keyPair, err := ec2.NewKeyPair(ctx, "keypair", &ec2.KeyPairArgs{
			PublicKey: sshKey.PublicKeyOpenssh,
		})

		if err != nil {
			return err
		}

		ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Owners: []string{
				"136693071363",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"debian-12-*",
					},
				},
				{
					Name: "architecture",
					Values: []string{
						"x86_64",
					},
				},
			},
		})

		if err != nil {
			return err
		}

		validatorKey, err := svmkit.NewKeyPair(ctx, "validator-key", nil)

		if err != nil {
			return err
		}

		voteAccountKey, err := svmkit.NewKeyPair(ctx, "vote-account-key", nil)

		if err != nil {
			return err
		}

		securityGroup, err := ec2.NewSecurityGroup(ctx, "security-group", &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					FromPort: pulumi.Int(22),
					ToPort:   pulumi.Int(22),
					Protocol: pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
				&ec2.SecurityGroupIngressArgs{
					FromPort: pulumi.Int(8000),
					ToPort:   pulumi.Int(8020),
					Protocol: pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
				&ec2.SecurityGroupIngressArgs{
					FromPort: pulumi.Int(8000),
					ToPort:   pulumi.Int(8020),
					Protocol: pulumi.String("udp"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})

		if err != nil {
			return err
		}

		instance, err := ec2.NewInstance(ctx, "instance", &ec2.InstanceArgs{
			Ami:          pulumi.String(ami.Id),
			InstanceType: pulumi.String("m5.2xlarge"),
			KeyName:      keyPair.KeyName,
			VpcSecurityGroupIds: pulumi.StringArray{
				securityGroup.ID(),
			},
			EbsBlockDevices: &ec2.InstanceEbsBlockDeviceArray{
				&ec2.InstanceEbsBlockDeviceArgs{
					DeviceName: pulumi.String("/dev/sdf"),
					VolumeSize: pulumi.Int(500),
					VolumeType: pulumi.String("io2"),
					Iops:       pulumi.Int(16000),
				},
				&ec2.InstanceEbsBlockDeviceArgs{
					DeviceName: pulumi.String("/dev/sdg"),
					VolumeSize: pulumi.Int(1024),
					VolumeType: pulumi.String("io2"),
					Iops:       pulumi.Int(16000),
				},
			},
			UserData: pulumi.String(`#!/bin/bash
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
`),
		})

		if err != nil {
			return err
		}

		_, err = validator.NewAgave(ctx,
			"validator",
			&validator.AgaveArgs{
				Connection: &ssh.ConnectionArgs{
					Host:       instance.PublicDns,
					User:       pulumi.String("admin"),
					PrivateKey: sshKey.PrivateKeyOpenssh,
				},
				Version: pulumi.String("1.18.24-1"),
				KeyPairs: &agave.KeyPairsArgs{
					Identity:    validatorKey.Json,
					VoteAccount: voteAccountKey.Json,
				},
				Flags: &agave.FlagsArgs{
					EntryPoint: pulumi.StringArray{
						pulumi.String("entrypoint.testnet.solana.com:8001"),
						pulumi.String("entrypoint2.testnet.solana.com:8001"),
						pulumi.String("entrypoint3.testnet.solana.com:8001"),
					},
					KnownValidator: pulumi.StringArray{
						pulumi.String("5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on"),
						pulumi.String("7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY"),
						pulumi.String("Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN"),
						pulumi.String("9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv"),
					},
					ExpectedGenesisHash:          pulumi.String("4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY"),
					UseSnapshotArchivesAtStartup: pulumi.String("when-newest"),
					RpcPort:                      pulumi.Int(8899),
					PrivateRPC:                   pulumi.Bool(true),
					OnlyKnownRPC:                 pulumi.Bool(true),
					DynamicPortRange:             pulumi.String("8002-8020"),
					GossipPort:                   pulumi.Int(8001),
					RpcBindAddress:               pulumi.String("0.0.0.0"),
					WalRecoveryMode:              pulumi.String("skip_any_corrupted_record"),
					LimitLedgerSize:              pulumi.Int(50000000),
					BlockProductionMethod:        pulumi.String("central-scheduler"),
					FullSnapshotIntervalSlots:    pulumi.Int(1000),
					NoWaitForVoteToStartLeader:   pulumi.Bool(true),
				},
			},
			pulumi.DependsOn([]pulumi.Resource{instance}),
		)

		if err != nil {
			return err
		}
		
		ctx.Export("nodes_name", pulumi.StringArray{pulumi.String("instance")})
		ctx.Export("nodes_public_ip", pulumi.StringArray{instance.PublicIp})
		ctx.Export("nodes_private_key", pulumi.StringArray{sshKey.PrivateKeyOpenssh})

		return nil
	})
}
