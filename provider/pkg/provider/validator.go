package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a StaticPage component resource.
type ValidatorArgs struct {
	Connection remote.ConnectionArgs `pulumi:"connection"`
}

// The Validator component resource.
type Validator struct {
	pulumi.ResourceState
}

// NewValidator setups a solana validator
func NewValidator(ctx *pulumi.Context,
	name string, args *ValidatorArgs, opts ...pulumi.ResourceOption) (*Validator, error) {
	if args == nil {
		args = &ValidatorArgs{}
	}

	remote.NewCommand(ctx, "hostnameCmd", &remote.CommandArgs{
		Create:     pulumi.String("hostname"),
		Connection: &args.Connection,
	})

	component := &Validator{}
	err := ctx.RegisterComponentResource("svmkit:index:Validator", name, component, opts...)
	if err != nil {
		return nil, err
	}

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}
