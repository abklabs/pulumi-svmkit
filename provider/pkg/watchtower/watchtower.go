package watchtower

import (
	"context"

	_ "embed"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/solana/watchtower"
)

// Watchtower represents a Pulumi resource for managing an Watchtower.
type Watchtower struct{}

// WatchtowerArgs represents the input arguments required to create an Watchtower resource.
type WatchtowerArgs struct {
	utils.RunnerArgs
	watchtower.Watchtower
}

// WatchtowerState represents the state of an Watchtower resource.
type WatchtowerState struct {
	WatchtowerArgs
}

// Create is the method that Pulumi calls to create an Watchtower resource.
// It sets up the Watchtower on the specified machine using the provided connection and flags.
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
func (Watchtower) Create(ctx context.Context, name string, input WatchtowerArgs, preview bool) (string, WatchtowerState, error) {
	state := WatchtowerState{WatchtowerArgs: input}

	if preview {
		return name, state, nil
	}

	watchtower := input.Watchtower

	command := watchtower.Install()

	if err := utils.RunnerHelper(ctx, input.Connection, command); err != nil {
		return "", WatchtowerState{}, err
	}

	return name, state, nil
}
