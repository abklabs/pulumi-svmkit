package account

import (
	"context"
	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/pulumi-svmkit/pkg/utils"
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

	if err := utils.RunnerHelper(ctx, input.Connection, command); err != nil {
		return "", VoteAccountState{}, err
	}

	return name, state, nil
}

func (VoteAccount) Update(ctx context.Context, name string, old VoteAccountState, new VoteAccountArgs, preview bool) (VoteAccountState, error) {
	return old, fmt.Errorf("Vote accounts may not be modified after creation!")
}

func (VoteAccount) Delete(ctx context.Context, name string, output VoteAccountState) error {
	client := output.VoteAccount

	command := client.Delete()

	if err := utils.RunnerHelper(ctx, output.Connection, command); err != nil {
		return err
	}

	return nil
}
