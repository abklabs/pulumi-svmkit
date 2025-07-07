package validator

import (
	"context"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/agave"
	"github.com/abklabs/svmkit/pkg/validator"
)

type Agave struct{}

type AgaveArgs struct {
	utils.RunnerArgs
	agave.Agave
}

type AgaveState struct {
	AgaveArgs
	validator.Properties
}

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

func (Agave) Update(ctx context.Context, name string, state AgaveState, newInput AgaveArgs, preview bool) (AgaveState, error) {
	state.AgaveArgs = newInput

	if preview {
		return state, nil
	}

	agave := newInput.Agave
	command := agave.Install()

	if err := utils.RunnerHelper(ctx, newInput.RunnerArgs, command); err != nil {
		return AgaveState{}, err
	}

	return state, nil
}

func (Agave) Delete(ctx context.Context, id string, props AgaveState) error {
	agave := props.Agave
	command := agave.Uninstall()

	if err := utils.RunnerHelper(ctx, props.RunnerArgs, command); err != nil {
		return err
	}

	return nil
}
