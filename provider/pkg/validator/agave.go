package validator

import (
	"context"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/agave"
	"github.com/abklabs/svmkit/pkg/validator"
)

// Agave represents a Pulumi resource for managing an Agave validator.
type Agave struct{}

// AgaveArgs represents the input arguments required to create an Agave resource.
type AgaveArgs struct {
	utils.RunnerArgs
	agave.Agave
}

// AgaveState represents the state of an Agave resource.
type AgaveState struct {
	AgaveArgs
	validator.Properties
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

	agave := input.Agave

	command := agave.Install()

	if err := utils.RunnerHelper(ctx, input.RunnerArgs, command); err != nil {
		return "", AgaveState{}, err
	}

	state.Properties = agave.Properties()

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

	agave := newInput.Agave
	command := agave.Install()

	if err := utils.RunnerHelper(ctx, newInput.RunnerArgs, command); err != nil {
		return "", AgaveState{}, err
	}

	return name, state, nil
}

func (Agave) Delete(ctx context.Context, id string, props AgaveState) error {
	agave := props.Agave
	command := agave.Uninstall()

	if err := utils.RunnerHelper(ctx, props.RunnerArgs, command); err != nil {
		return err
	}

	return nil
}
