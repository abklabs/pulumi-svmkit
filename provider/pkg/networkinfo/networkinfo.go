package networkinfo

import (
	"context"

	"github.com/abklabs/svmkit/pkg/solana"
)

type GetNetworkInfoInput struct {
	NetworkName solana.NetworkName `pulumi:"networkName"`
}

type GetNetworkInfoOutput struct {
	GetNetworkInfoInput
	solana.NetworkInfo
}

type GetNetworkInfo struct{}

func (GetNetworkInfo) Call(ctx context.Context, input GetNetworkInfoInput) (GetNetworkInfoOutput, error) {
	info, err := solana.LookupNetworkInfo(input.NetworkName)

	if err != nil {
		return GetNetworkInfoOutput{}, err
	}

	output := GetNetworkInfoOutput{
		GetNetworkInfoInput: input,
		NetworkInfo:         *info,
	}

	return output, nil
}
