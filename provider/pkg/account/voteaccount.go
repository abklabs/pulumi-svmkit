package account

import (
	"context"

	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/svmkit/pkg/runner"
	"github.com/abklabs/svmkit/pkg/solana"
)

type VoteAccount struct{}

type VoteAccountArgs struct {
	svm.ValidatorArgs
	solana.VoteAccount
}

type VoteAccountState struct {
	VoteAccountArgs
}

func (VoteAccount) Create(ctx context.Context, name string, input VoteAccountArgs, preview bool) (string, VoteAccountState, error) {
	state := VoteAccountState{VoteAccountArgs: input}

	if preview {
		return name, state, nil
	}

	client := input.VoteAccount

	command := client.Create()

	if err := command.Check(); err != nil {
		return "", VoteAccountState{}, fmt.Errorf("failed to check config: %w", err)
	}

	r := runner.NewRunner(input.Connection, command)

	if err := r.Run(ctx); err != nil {
		return "", VoteAccountState{}, fmt.Errorf("failed to install: %w", err)
	}

	return name, state, nil
}

func (VoteAccount) Update(ctx context.Context, name string, old VoteAccountState, new VoteAccountArgs, preview bool) (VoteAccountState, error) {
	if preview {
		return VoteAccountState{VoteAccountArgs: new}, nil
	}

	old.VoteAccountArgs = new

	return old, nil
}

func (VoteAccount) Delete(ctx context.Context, name string, output VoteAccountState) error {
	client := output.VoteAccount

	command := client.Delete()

	if err := command.Check(); err != nil {
		return fmt.Errorf("failed to check config: %w", err)
	}

	r := runner.NewRunner(output.Connection, command)

	if err := r.Run(ctx); err != nil {
		return fmt.Errorf("failed to install: %w", err)
	}

	return nil
}
