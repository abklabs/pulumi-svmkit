package account

import (
	"context"
	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/pulumi-svmkit/pkg/utils"
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

// getStakeClient is a helper method to establish an SSH connection and create a StakeAccountClient
func (StakeAccount) getStakeClient(ctx context.Context, args svm.ValidatorArgs) (*solana.StakeAccountClient, error) {
	sshClient, err := args.Connection.Dial(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to dial SSH connection to host: %w", err)
	}

	logger := utils.MakePulumiLogger(ctx)
	stakeOperator := solana.NewCliStakeOperator(
		sshClient,
		logger,
		ctx,
	)

	return solana.NewStakeAccountClient(stakeOperator), nil
}

func (s StakeAccount) Create(ctx context.Context, name string, input StakeAccountArgs, preview bool) (string, StakeAccountState, error) {
	if preview {
		state := StakeAccountState{StakeAccountArgs: input}
		return name, state, nil
	}

	client, err := s.getStakeClient(ctx, input.ValidatorArgs)
	if err != nil {
		return "", StakeAccountState{}, err
	}

	newStakeAccount, err := client.Create(input.StakeAccount)
	if err != nil {
		return "", StakeAccountState{}, err
	}

	newArgs := StakeAccountArgs{input.ValidatorArgs, newStakeAccount}
	state := StakeAccountState{newArgs}

	return name, state, nil
}

func (s StakeAccount) Update(ctx context.Context, name string, old StakeAccountState, newArgs StakeAccountArgs, preview bool) (StakeAccountState, error) {
	if preview {
		state := StakeAccountState{StakeAccountArgs: newArgs}
		return state, nil
	}

	client, err := s.getStakeClient(ctx, newArgs.ValidatorArgs)
	if err != nil {
		return StakeAccountState{}, err
	}

	newStakeAccount, err := client.Update(old.StakeAccount, newArgs.StakeAccount)
	if err != nil {
		return StakeAccountState{}, err
	}

	updatedArgs := StakeAccountArgs{newArgs.ValidatorArgs, newStakeAccount}
	state := StakeAccountState{updatedArgs}

	return state, nil
}

func (s StakeAccount) Delete(ctx context.Context, name string, state StakeAccountState) error {
	client, err := s.getStakeClient(ctx, state.ValidatorArgs)
	if err != nil {
		return err
	}

	err = client.Delete(state.StakeAccount)
	if err != nil {
		return err
	}

	return nil
}
