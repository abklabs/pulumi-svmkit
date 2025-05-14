package explorer

import (
	"context"

	_ "embed"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/solana/explorer"
)

// Explorer represents a Pulumi resource for managing an Explorer.
type Explorer struct{}

// ExplorerArgs represents the input arguments required to create an Explorer resource.
type ExplorerArgs struct {
	utils.RunnerArgs
	explorer.Explorer
}

// ExplorerState represents the state of an Explorer resource.
type ExplorerState struct {
	ExplorerArgs
}

// Create is the method that Pulumi calls to create an Explorer resource.
// It sets up the Explorer on the specified machine using the provided connection and flags.
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
func (Explorer) Create(ctx context.Context, name string, input ExplorerArgs, preview bool) (string, ExplorerState, error) {
	state := ExplorerState{ExplorerArgs: input}

	if preview {
		return name, state, nil
	}

	explorer := input.Explorer

	command := explorer.Install()

	if err := utils.RunnerHelper(ctx, input.Connection, command); err != nil {
		return "", ExplorerState{}, err
	}

	return name, state, nil
}
