package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a Ready component resource.
type ReadyArgs struct {
	Connection    remote.ConnectionInput `pulumi:"connection" provider:"secret"`
	CustomTimeout string                 `pulumi:"customTimeout"`
}

// The Ready component resource.
type Ready struct {
	pulumi.ResourceState
}

// NewReady waits for an SSH connection to be active.
func NewReady(ctx *pulumi.Context,
	name string, args *ReadyArgs, opts ...pulumi.ResourceOption) (*Ready, error) {
	if args == nil {
		args = &ReadyArgs{}
	}

	// Set default values if not provided
	if args.CustomTimeout == "" {
		args.CustomTimeout = "10m" // Default timeout
	}

	ready := &Ready{}
	err := ctx.RegisterComponentResource("svmkit:index:Ready", name, ready, opts...)
	if err != nil {
		return nil, err
	}

	_, err = remote.NewCommand(ctx, "isReady", &remote.CommandArgs{
		Create:     pulumi.String("echo connected"),
		Connection: args.Connection,
	}, pulumi.Parent(ready), pulumi.Timeouts(&pulumi.CustomTimeouts{Create: args.CustomTimeout}))
	if err != nil {
		return nil, err
	}

	if err := ctx.RegisterResourceOutputs(ready, pulumi.Map{}); err != nil {
		return nil, err
	}

	return ready, nil
}
