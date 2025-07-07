package validator

import (
	"context"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/firedancer"
	"github.com/abklabs/svmkit/pkg/validator"
)

type Firedancer struct{}

type FiredancerArgs struct {
	utils.RunnerArgs
	firedancer.Firedancer
}

type FiredancerState struct {
	FiredancerArgs
	validator.Properties
}

func (Firedancer) Create(ctx context.Context, name string, input FiredancerArgs, preview bool) (string, FiredancerState, error) {
	state := FiredancerState{FiredancerArgs: input}

	if preview {
		return name, state, nil
	}

	fd := input.Firedancer

	command := fd.Install()

	if err := utils.RunnerHelper(ctx, input.RunnerArgs, command); err != nil {
		return "", FiredancerState{}, err
	}

	state.Properties = fd.Properties()

	return name, state, nil
}

func (Firedancer) Update(ctx context.Context, name string, state FiredancerState, newInput FiredancerArgs, preview bool) (FiredancerState, error) {
	state.FiredancerArgs = newInput

	if preview {
		return state, nil
	}

	fd := newInput.Firedancer

	command := fd.Install()

	if err := utils.RunnerHelper(ctx, newInput.RunnerArgs, command); err != nil {
		return FiredancerState{}, err
	}

	return state, nil
}

func (Firedancer) Delete(ctx context.Context, id string, props FiredancerState) error {
	fd := props.Firedancer

	command := fd.Uninstall()

	if err := utils.RunnerHelper(ctx, props.RunnerArgs, command); err != nil {
		return err
	}

	return nil
}
