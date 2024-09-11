package provider

import (
	"errors"

	"github.com/abklabs/pulumi-svmkit/pkg/agave"
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a validator component resource.
type ValidatorArgs struct {
	Connection   remote.ConnectionInput `pulumi:"connection" provider:"secret"`
	Identity     pulumi.StringInput     `pulumi:"identity" provider:"secret"`
	StakeAccount *pulumi.StringInput    `pulumi:"stakeAccount,optional" provider:"secret"`
	VoteAccount  pulumi.StringInput     `pulumi:"voteAccount" provider:"secret"`
	User         pulumi.StringInput     `pulumi:"user"`
}

// The Validator component resource.
type Validator struct {
	pulumi.ResourceState
}

const (
	identityFile     = "validator-keypair.json"
	stakeAccountFile = "stake-keypair.json"
	voteAccountFile  = "vote-account-keypair.json"
)

// NewValidator setups a solana validator
func NewValidator(ctx *pulumi.Context,
	name string, args *ValidatorArgs, opts ...pulumi.ResourceOption) (*Validator, error) {
	if args == nil {
		return nil, errors.New("missing required arguments")
	}
	if args.Connection == nil {
		return nil, errors.New("missing required argument 'Connection'")
	}

	// Write the key pairs to the filesystem
	args.User.ToStringOutput().ApplyT(func(user string) error {
		args.Identity.ToStringOutput().ApplyT(func(pk string) (*remote.Command, error) {
			return agave.WriteKeyPair(ctx, args.Connection, agave.KeyPairArgs{
				Name:     "writeIdentityKeyPair",
				Key:      pk,
				User:     user,
				FileName: identityFile,
			})
		})

		args.VoteAccount.ToStringOutput().ApplyT(func(pk string) (*remote.Command, error) {
			return agave.WriteKeyPair(ctx, args.Connection, agave.KeyPairArgs{
				Name:     "writeVoteAccountKeyPair",
				Key:      pk,
				User:     user,
				FileName: voteAccountFile,
			})
		})

		if args.StakeAccount != nil {
			(*args.StakeAccount).ToStringOutput().ApplyT(func(pk string) (*remote.Command, error) {
				return agave.WriteKeyPair(ctx, args.Connection, agave.KeyPairArgs{
					Name:     "writeStakeAccountKeyPair",
					Key:      pk,
					User:     user,
					FileName: stakeAccountFile,
				})
			})
		}

		return nil
	})

	// Configure sysctl
	agave.ConfigureSysctl(ctx, args.Connection)

	// Set up UFW rules
	agave.SetUfwRules(ctx, args.Connection)

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
