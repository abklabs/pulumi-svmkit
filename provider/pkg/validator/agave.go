package validator

import (
	"context"

	_ "embed"
	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/svmkit/pkg/agave"
	"github.com/abklabs/svmkit/pkg/runner"
)

// Agave represents a Pulumi resource for managing an Agave validator.
type Agave struct{}

// AgaveArgs represents the input arguments required to create an Agave resource.
type AgaveArgs struct {
	svm.ValidatorArgs
	// Flags contains the configuration flags for the Agave validator.
	Flags    agave.Flags    `pulumi:"flags"`
	KeyPairs agave.KeyPairs `pulumi:"keyPairs"`
	Metrics  *agave.Metrics `pulumi:"metrics,optional"`
}

// AgaveState represents the state of an Agave resource.
type AgaveState struct {
	AgaveArgs
}

// Create is the method that Pulumi calls to create an Agave resource.
// It sets up the Agave validator on the specified machine using the provided connection and flags.
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
func (Agave) Create(ctx context.Context, name string, input AgaveArgs, preview bool) (string, AgaveState, error) {
	state := AgaveState{AgaveArgs: input}

	if preview {
		return name, state, nil
	}

	client := &agave.Agave{
		Flags:    input.Flags,
		KeyPairs: input.KeyPairs,
		Version:  input.Version,
		Metrics:  input.Metrics,
	}
	command := client.Install()

	r := runner.Machine(input.Connection).
		Env(command.Env()).
		Script(command.Script())

	if err := r.Run(ctx); err != nil {
		return "", AgaveState{}, fmt.Errorf("failed to install validator: %w", err)
	}
	return name, state, nil
}

// Update is the method that Pulumi calls to update an Agave resource.
// It sets up the Agave validator on the specified machine using the provided connection and flags.
//
// Parameters:
// - ctx: The context for the update operation.
// - name: The name of the resource.
// - oldInput: The previous input arguments for the resource.
// - newInput: The new input arguments for updating the resource.
// - preview: A boolean indicating whether this is a preview operation.
//
// Returns:
// - The name of the updated resource.
// - The state of the updated resource.
// - An error if the update fails.
func (Agave) Update(ctx context.Context, name string, oldInput, newInput AgaveArgs, preview bool) (string, AgaveState, error) {
	state := AgaveState{AgaveArgs: newInput}

	if preview {
		return name, state, nil
	}

	client := &agave.Agave{
		Flags:    newInput.Flags,
		KeyPairs: newInput.KeyPairs,
	}
	command := client.Install()

	r := runner.Machine(newInput.Connection).
		Env(command.Env()).
		Script(command.Script())

	if err := r.Run(ctx); err != nil {
		return "", AgaveState{}, fmt.Errorf("failed to update validator: %w", err)
	}

	return name, state, nil
}
