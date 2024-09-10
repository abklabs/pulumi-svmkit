package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a User component resource.
type UserArgs struct {
	Connection remote.ConnectionInput `pulumi:"connection" provider:"secret"`
	Username   pulumi.StringInput     `pulumi:"username"`
}

// The User component resource.
type User struct {
	pulumi.ResourceState
	PubKey  pulumi.StringOutput `pulumi:"pubKey"`
	PrivKey pulumi.StringOutput `pulumi:"privKey" provider:"secret"`
	Name    pulumi.StringOutput `pulumi:"name"`
}

// NewUser creates a new user on the host, generates an SSH key for the user, and adds it to the box.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, fmt.Errorf("missing required arguments")
	}
	if args.Connection == nil {
		return nil, fmt.Errorf("missing required argument 'Connection'")
	}
	if args.Username == nil {
		return nil, fmt.Errorf("missing required argument 'Username'")
	}

	// Create the user on the host
	addUser, err := remote.NewCommand(ctx, "addUser", &remote.CommandArgs{
		Create:     pulumi.Sprintf("sudo adduser --disabled-password --gecos '' %s", args.Username),
		Connection: args.Connection,
	})
	if err != nil {
		return nil, err
	}

	// Generate an SSH key for the user using the tls.PrivateKey resource
	sshKey, err := tls.NewPrivateKey(ctx, "createSshKey", &tls.PrivateKeyArgs{
		Algorithm: pulumi.String("ED25519"),
	})
	if err != nil {
		return nil, err
	}

	// Set ownership of the home directory to the user
	_, err = remote.NewCommand(ctx, "setHomeOwnership", &remote.CommandArgs{
		Create:     pulumi.Sprintf("sudo chown %s:%s /home/%s", args.Username, args.Username, args.Username),
		Connection: args.Connection,
	}, pulumi.DependsOn([]pulumi.Resource{addUser}))
	if err != nil {
		return nil, err
	}

	// Create the .ssh directory and set the correct permissions
	setupSSHDir, err := remote.NewCommand(ctx, "setupSSHDir", &remote.CommandArgs{
		Create:     pulumi.Sprintf("sudo mkdir -p /home/%s/.ssh && sudo chmod 700 /home/%s/.ssh && sudo chown %s:%s /home/%s/.ssh", args.Username, args.Username, args.Username, args.Username, args.Username),
		Connection: args.Connection,
	}, pulumi.DependsOn([]pulumi.Resource{sshKey, addUser}))
	if err != nil {
		return nil, err
	}

	// Add the SSH key to the authorized_keys file
	_, err = remote.NewCommand(ctx, "addSSHKey", &remote.CommandArgs{
		Create:     pulumi.Sprintf("sudo sh -c \"echo '%s' >> /home/%s/.ssh/authorized_keys\"", sshKey.PublicKeyOpenssh, args.Username),
		Connection: args.Connection,
	}, pulumi.DependsOn([]pulumi.Resource{setupSSHDir}))
	if err != nil {
		return nil, err
	}

	component := &User{}
	err = ctx.RegisterComponentResource("svmkit:index:User", name, component, opts...)
	if err != nil {
		return nil, err
	}

	component.PubKey = pulumi.ToOutput(sshKey.PublicKeyOpenssh).(pulumi.StringOutput)
	component.PrivKey = pulumi.ToSecret(sshKey.PrivateKeyOpenssh).(pulumi.StringOutput)
	component.Name = args.Username.ToStringOutput()

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"pubKey":  component.PubKey,
		"privKey": component.PrivKey,
		"name":    component.Name,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
