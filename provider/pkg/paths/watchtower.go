package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/solana/watchtower"
)

type WatchtowerPathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type WatchtowerPathsOutput struct {
	WatchtowerPathsInput
	watchtower.WatchtowerPaths
}

type GetDefaultWatchtowerPaths struct{}

func (GetDefaultWatchtowerPaths) Call(ctx context.Context, input WatchtowerPathsInput) (WatchtowerPathsOutput, error) {
	wp, err := watchtower.NewDefaultWatchtowerPaths(input.BasePaths)
	if err != nil {
		return WatchtowerPathsOutput{}, err
	}
	return WatchtowerPathsOutput{
		WatchtowerPathsInput: input,
		WatchtowerPaths:      *wp,
	}, nil
}
