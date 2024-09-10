package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a Wait component resource.
type WaitArgs struct {
	Connection    remote.ConnectionInput `pulumi:"connection" provider:"secret"`
	CustomTimeout string                 `pulumi:"customTimeout"`
}

// The Wait component resource.
type Wait struct {
	pulumi.ResourceState
}

// NewWait waits for an SSH connection to be active.
func NewWait(ctx *pulumi.Context,
	name string, args *WaitArgs, opts ...pulumi.ResourceOption) (*Wait, error) {
	if args == nil {
		args = &WaitArgs{}
	}

	// Set default values if not provided
	if args.CustomTimeout == "" {
		args.CustomTimeout = "10m" // Default timeout
	}

	component := &Wait{}
	err := ctx.RegisterComponentResource("svmkit:index:Wait", name, component, opts...)
	if err != nil {
		return nil, err
	}

	_, err = remote.NewCommand(ctx, "connected", &remote.CommandArgs{
		Create:     pulumi.String("echo connected"),
		Connection: args.Connection,
	}, pulumi.Parent(component), pulumi.Timeouts(&pulumi.CustomTimeouts{Create: args.CustomTimeout}))
	if err != nil {
		return nil, err
	}

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}
