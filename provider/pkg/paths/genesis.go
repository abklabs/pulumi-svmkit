package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/solana/genesis"
)

type GenesisPathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type GenesisPathsOutput struct {
	GenesisPathsInput
	genesis.GenesisPaths
}

type GetDefaultGenesisPaths struct{}

func (GetDefaultGenesisPaths) Call(ctx context.Context, input GenesisPathsInput) (GenesisPathsOutput, error) {
	gp, err := genesis.NewDefaultGenesisPaths(input.BasePaths)
	if err != nil {
		return GenesisPathsOutput{}, err
	}
	return GenesisPathsOutput{
		GenesisPathsInput: input,
		GenesisPaths:      *gp,
	}, nil
}
