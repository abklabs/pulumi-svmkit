package account

import (
	"context"

	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/svmkit/pkg/runner"
	"github.com/abklabs/svmkit/pkg/solana"
)

type StakeAccount struct{}

type StakeAccountArgs struct {
	svm.ValidatorArgs
	solana.StakeAccount
}

type StakeAccountState struct {
	StakeAccountArgs
}

func (StakeAccount) Create(ctx context.Context, name string, input StakeAccountArgs, preview bool) (string, StakeAccountState, error) {
	state := StakeAccountState{StakeAccountArgs: input}

	if preview {
		return name, state, nil
	}

	client := input.StakeAccount

	command := client.Create()

	if err := command.Check(); err != nil {
		return "", StakeAccountState{}, fmt.Errorf("failed to check config: %w", err)
	}

	r := runner.NewRunner(input.Connection, command)

	if err := r.Run(ctx); err != nil {
		return "", StakeAccountState{}, fmt.Errorf("failed to install: %w", err)
	}

	return name, state, nil
}

func (StakeAccount) Update(ctx context.Context, name string, old StakeAccountState, new StakeAccountArgs, preview bool) (StakeAccountState, error) {
	if preview {
		return StakeAccountState{StakeAccountArgs: new}, nil
	}

	old.StakeAccountArgs = new

	return old, nil
}
