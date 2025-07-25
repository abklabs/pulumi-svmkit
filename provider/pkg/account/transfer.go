package account

import (
	"context"
	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/solana"
)

type Transfer struct{}

type TransferArgs struct {
	utils.RunnerArgs
	solana.Transfer
}

type TransferState struct {
	TransferArgs
}

func (Transfer) Create(ctx context.Context, name string, input TransferArgs, preview bool) (string, TransferState, error) {
	state := TransferState{TransferArgs: input}

	if preview {
		return name, state, nil
	}

	client := input.Transfer

	command := client.Create()

	if err := utils.RunnerHelper(ctx, input.RunnerArgs, command); err != nil {
		return "", TransferState{}, err
	}

	return name, state, nil
}

func (Transfer) Update(ctx context.Context, name string, old TransferState, new TransferArgs, preview bool) (TransferState, error) {
	return old, fmt.Errorf("transfers are sent once, and may not be modified after creation")
}
