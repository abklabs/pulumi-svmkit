package firewall

import (
	"context"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/firewall"
)

// Firewall represents a Pulumi resource for managing an Firewall.
type Firewall struct{}

// FirewallArgs represents the input arguments required to create an Firewall resource.
type FirewallArgs struct {
	utils.RunnerArgs
	firewall.Firewall
}

// FirewallState represents the state of an Firewall resource.
type FirewallState struct {
	FirewallArgs
}

// Create is the method that Pulumi calls to create an Firewall resource.
// It sets up the Firewall on the specified machine using the provided connection and flags.
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
func (Firewall) Create(ctx context.Context, name string, input FirewallArgs, preview bool) (string, FirewallState, error) {
	state := FirewallState{FirewallArgs: input}

	if preview {
		return name, state, nil
	}

	inputFirewall := input.Firewall

	command := inputFirewall.Create()

	if err := utils.RunnerHelper(ctx, input.RunnerArgs, command); err != nil {
		return "", FirewallState{}, err
	}

	return name, state, nil
}

// Update is the method that Pulumi calls to update an Firewall resource.
// It sets up the Firewall validator on the specified machine using the provided connection and flags.
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
func (Firewall) Update(ctx context.Context, name string, oldInput, newInput FirewallArgs, preview bool) (string, FirewallState, error) {
	state := FirewallState{FirewallArgs: newInput}

	if preview {
		return name, state, nil
	}

	inputFirewall := newInput.Firewall

	command := inputFirewall.Create()

	if err := utils.RunnerHelper(ctx, newInput.RunnerArgs, command); err != nil {
		return "", FirewallState{}, err
	}

	return name, state, nil
}
