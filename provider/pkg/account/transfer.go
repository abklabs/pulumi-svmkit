package account

import (
	"context"

	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/svmkit/pkg/runner"
	"github.com/abklabs/svmkit/pkg/solana"
)

type Transfer struct{}

type TransferArgs struct {
	svm.ValidatorArgs
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

	if err := command.Check(); err != nil {
		return "", TransferState{}, fmt.Errorf("failed to check validator config: %w", err)
	}

	r := runner.NewRunner(input.Connection, command)

	if err := r.Run(ctx); err != nil {
		return "", TransferState{}, fmt.Errorf("failed to install validator: %w", err)
	}

	return name, state, nil
}
