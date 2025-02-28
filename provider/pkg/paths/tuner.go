package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/tuner"
)

type TunerPathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type TunerPathsOutput struct {
	TunerPathsInput
	tuner.TunerPaths
}

type GetDefaultTunerPaths struct{}

func (GetDefaultTunerPaths) Call(ctx context.Context, input TunerPathsInput) (TunerPathsOutput, error) {
	tp, err := tuner.NewDefaultTunerPaths(input.BasePaths)
	if err != nil {
		return TunerPathsOutput{}, err
	}
	return TunerPathsOutput{
		TunerPathsInput: input,
		TunerPaths:      *tp,
	}, nil
}
