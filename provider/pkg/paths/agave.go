package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/agave"
)

type AgavePathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type AgavePathsOutput struct {
	AgavePathsInput
	agave.AgavePaths
}

type GetDefaultAgavePaths struct{}

func (GetDefaultAgavePaths) Call(ctx context.Context, input AgavePathsInput) (AgavePathsOutput, error) {
	ap, err := agave.NewDefaultAgavePaths(input.BasePaths)
	if err != nil {
		return AgavePathsOutput{}, err
	}
	return AgavePathsOutput{
		AgavePathsInput: input,
		AgavePaths:      *ap,
	}, nil
}
