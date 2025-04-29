package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
	"github.com/abklabs/svmkit/pkg/solana/faucet"
)

type FaucetPathsInput struct {
	BasePaths *paths.Paths `pulumi:"basePaths,optional"`
}

type FaucetPathsOutput struct {
	FaucetPathsInput
	faucet.FaucetPaths
}

type GetDefaultFaucetPaths struct{}

func (GetDefaultFaucetPaths) Call(ctx context.Context, input FaucetPathsInput) (FaucetPathsOutput, error) {
	fp, err := faucet.NewDefaultFaucetPaths(input.BasePaths)
	if err != nil {
		return FaucetPathsOutput{}, err
	}
	return FaucetPathsOutput{
		FaucetPathsInput: input,
		FaucetPaths:      *fp,
	}, nil
}
