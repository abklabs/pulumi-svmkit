package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
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
	Name pulumi.StringOutput `pulumi:"name"`
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

	// Set ownership of the home directory to the user and allow root to write to the user directory
	_, err = remote.NewCommand(ctx, "setHomeOwnership", &remote.CommandArgs{
		Create:     pulumi.Sprintf("sudo chown %s:%s /home/%s && sudo chmod 770 /home/%s", args.Username, args.Username, args.Username, args.Username),
		Connection: args.Connection,
	}, pulumi.DependsOn([]pulumi.Resource{addUser}))
	if err != nil {
		return nil, err
	}

	component := &User{}
	err = ctx.RegisterComponentResource("svmkit:index:User", name, component, opts...)
	if err != nil {
		return nil, err
	}

	component.Name = args.Username.ToStringOutput()

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"name": component.Name,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
