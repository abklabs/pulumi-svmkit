package validator

import (
	"context"

	_ "embed"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/firedancer"
)

type Firedancer struct{}

type FiredancerArgs struct {
	utils.RunnerArgs
	firedancer.Firedancer
}

type FiredancerState struct {
	FiredancerArgs
}

func (Firedancer) Create(ctx context.Context, name string, input FiredancerArgs, preview bool) (string, FiredancerState, error) {
	state := FiredancerState{FiredancerArgs: input}

	if preview {
		return name, state, nil
	}

	fd := input.Firedancer

	command := fd.Install()

	if err := utils.RunnerHelper(ctx, input.Connection, command); err != nil {
		return "", FiredancerState{}, err
	}

	return name, state, nil
}

func (Firedancer) Delete(ctx context.Context, id string, props FiredancerState) error {
	fd := props.Firedancer

	command := fd.Uninstall()

	if err := utils.RunnerHelper(ctx, props.Connection, command); err != nil {
		return err
	}

	return nil
}
