package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/firedancer"
)

type FiredancerPathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type FiredancerPathsOutput struct {
	FiredancerPathsInput
	firedancer.FiredancerPaths
}

type GetDefaultFiredancerPaths struct{}

func (GetDefaultFiredancerPaths) Call(ctx context.Context, input FiredancerPathsInput) (FiredancerPathsOutput, error) {
	fp, err := firedancer.NewDefaultFiredancerPaths(input.BasePaths)
	if err != nil {
		return FiredancerPathsOutput{}, err
	}
	return FiredancerPathsOutput{
		FiredancerPathsInput: input,
		FiredancerPaths:      *fp,
	}, nil
}
