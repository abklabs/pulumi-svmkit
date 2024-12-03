package validator

import (
	"context"

	_ "embed"

	"github.com/abklabs/pulumi-svmkit/pkg/svm"
	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/firedancer"
)

type Firedancer struct{}

type FiredancerArgs struct {
	svm.ValidatorArgs
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
