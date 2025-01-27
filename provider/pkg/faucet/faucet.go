package faucet

import (
	"context"

	_ "embed"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/solana/faucet"
)

// Faucet represents a Pulumi resource for managing an Faucet.
type Faucet struct{}

// FaucetArgs represents the input arguments required to create an Faucet resource.
type FaucetArgs struct {
	svm.FaucetArgs
	faucet.Faucet
}

// FaucetState represents the state of an Faucet resource.
type FaucetState struct {
	FaucetArgs
}

// Create is the method that Pulumi calls to create an Faucet resource.
// It sets up the Faucet on the specified machine using the provided connection and flags.
//
// Parameters:
// - ctx: The context for the creation operation.
// - name: The name of the resource.
// - input: The input arguments for creating the resource.
// - preview: A boolean indicating whether this is a preview operation.
//
// Returns:
// - The name of the created resource.
// - The state of the created resource.
// - An error if the creation fails.
func (Faucet) Create(ctx context.Context, name string, input FaucetArgs, preview bool) (string, FaucetState, error) {
	state := FaucetState{FaucetArgs: input}

	if preview {
		return name, state, nil
	}

	faucet := input.Faucet

	command := faucet.Install()

	if err := utils.RunnerHelper(ctx, input.Connection, command); err != nil {
		return "", FaucetState{}, err
	}

	return name, state, nil
}
