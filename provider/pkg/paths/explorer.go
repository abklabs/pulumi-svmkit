package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/solana/explorer"
)

type ExplorerPathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type ExplorerPathsOutput struct {
	ExplorerPathsInput
	explorer.ExplorerPaths
}

type GetDefaultExplorerPaths struct{}

func (GetDefaultExplorerPaths) Call(ctx context.Context, input ExplorerPathsInput) (ExplorerPathsOutput, error) {
	ep, err := explorer.NewDefaultExplorerPaths(input.BasePaths)
	if err != nil {
		return ExplorerPathsOutput{}, err
	}
	return ExplorerPathsOutput{
		ExplorerPathsInput: input,
		ExplorerPaths:      *ep,
	}, nil
}
