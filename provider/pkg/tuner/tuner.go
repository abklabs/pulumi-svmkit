package tuner

import (
	"context"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/tuner"
)

// Tuner represents a Pulumi resource for managing an Tuner.
type Tuner struct{}

// TunerArgs represents the input arguments required to create an Tuner resource.
type TunerArgs struct {
	utils.RunnerArgs
	tuner.Tuner
}

// TunerState represents the state of an Tuner resource.
type TunerState struct {
	TunerArgs
}

// Create is the method that Pulumi calls to create an Tuner resource.
// It sets up the Tuner on the specified machine using the provided connection and flags.
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
func (Tuner) Create(ctx context.Context, name string, input TunerArgs, preview bool) (string, TunerState, error) {
	state := TunerState{TunerArgs: input}

	if preview {
		return name, state, nil
	}

	inputTuner := input.Tuner

	command := inputTuner.Create()

	if err := utils.RunnerHelper(ctx, input.Connection, command); err != nil {
		return "", TunerState{}, err
	}

	return name, state, nil
}

// Update is the method that Pulumi calls to update an Tuner resource.
// It sets up the Tuner validator on the specified machine using the provided connection and flags.
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
func (Tuner) Update(ctx context.Context, name string, oldInput, newInput TunerArgs, preview bool) (string, TunerState, error) {
	state := TunerState{TunerArgs: newInput}

	if preview {
		return name, state, nil
	}

	inputTuner := newInput.Tuner

	command := inputTuner.Create()

	if err := utils.RunnerHelper(ctx, newInput.Connection, command); err != nil {
		return "", TunerState{}, err
	}

	return name, state, nil
}
