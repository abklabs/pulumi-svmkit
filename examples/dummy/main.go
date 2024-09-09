package main

import (
	svmkit "github.com/abklabs/pulumi-svmkit"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		sshKey, err := tls.NewPrivateKey(ctx, "sshKey", &tls.PrivateKeyArgs{
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
			Filters: []ec2.GetAmiFilter{
				{
					Name:   "name",
					Values: []string{"debian-12-*"},
				},
				{
					Name:   "architecture",
					Values: []string{"x86_64"},
				},
			},
			Owners:     []string{"136693071363"}, // Debian
			MostRecent: pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}

		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow SSH"),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(22),
					ToPort:     pulumi.Int(22),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					Protocol:   pulumi.String("-1"),
					FromPort:   pulumi.Int(0),
					ToPort:     pulumi.Int(0),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
		})
		if err != nil {
			return err
		}

		instance, err := ec2.NewInstance(ctx, "instance", &ec2.InstanceArgs{
			Ami:                 pulumi.String(ami.Id),
			InstanceType:        pulumi.String("t3.small"),
			KeyName:             keyPair.KeyName,
			VpcSecurityGroupIds: pulumi.StringArray{securityGroup.ID()},
		})
		if err != nil {
			return err
		}

		connection := remote.ConnectionArgs{
			Host:       instance.PublicDns,
			User:       pulumi.String("admin"),
			PrivateKey: sshKey.PrivateKeyOpenssh,
		}

		_, err = svmkit.NewValidator(ctx, "validator", &svmkit.ValidatorArgs{
			Connection: connection,
		})
		if err != nil {
			return err
		}

		ctx.Export("PUBLIC_DNS_NAME", instance.PublicDns)
		ctx.Export("SSH_PRIVATE_KEY", sshKey.PrivateKeyOpenssh)

		return nil
	})
}
