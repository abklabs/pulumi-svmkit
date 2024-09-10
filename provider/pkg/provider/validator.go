package provider

import (
	"errors"
	"fmt"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a validator component resource.
type ValidatorArgs struct {
	Connection   remote.ConnectionInput `pulumi:"connection" provider:"secret"`
	Identity     pulumi.StringInput     `pulumi:"identity" provider:"secret"`
	StakeAccount *pulumi.StringInput    `pulumi:"stakeAccount,optional" provider:"secret"`
	VoteAccount  pulumi.StringInput     `pulumi:"voteAccount" provider:"secret"`
}

// The Validator component resource.
type Validator struct {
	pulumi.ResourceState
}

const identityFile = "validator-keypair.json"
const stakeAccountFile = "stake-keypair.json"
const voteAccountFile = "vote-account-keypair.json"

// NewValidator setups a solana validator
func NewValidator(ctx *pulumi.Context,
	name string, args *ValidatorArgs, opts ...pulumi.ResourceOption) (*Validator, error) {
	if args == nil {
		return nil, errors.New("missing required arguments")
	}
	if args.Connection == nil {
		return nil, errors.New("missing required argument 'Connection'")
	}

	args.Identity.ToStringOutput().ApplyT(func(pk string) (*remote.Command, error) {
		cmd := fmt.Sprintf("echo %s > %s", pk, identityFile)

		return remote.NewCommand(ctx, "identity", &remote.CommandArgs{
			Create:     pulumi.String(cmd),
			Connection: args.Connection,
		})
	})

	args.VoteAccount.ToStringOutput().ApplyT(func(pk string) (*remote.Command, error) {
		cmd := fmt.Sprintf("echo %s > %s", pk, voteAccountFile)

		return remote.NewCommand(ctx, "voteAccount", &remote.CommandArgs{
			Create:     pulumi.String(cmd),
			Connection: args.Connection,
		})
	})

	if args.StakeAccount != nil {
		(*args.StakeAccount).ToStringOutput().ApplyT(func(pk string) (*remote.Command, error) {
			cmd := fmt.Sprintf("echo %s > %s", pk, stakeAccountFile)

			return remote.NewCommand(ctx, "stakeAccount", &remote.CommandArgs{
				Create:     pulumi.String(cmd),
				Connection: args.Connection,
			})
		})
	}

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
