package firewall

import (
	"context"

	"github.com/abklabs/svmkit/pkg/firewall"
)

type FirewallParamsInput struct {
	Variant firewall.FirewallVariant `pulumi:"variant"`
}

type FirewallParamsOutput struct {
	FirewallParamsInput
	firewall.FirewallParams
}

type GetDefaultFirewallParams struct{}

func (GetDefaultFirewallParams) Call(ctx context.Context, input FirewallParamsInput) (FirewallParamsOutput, error) {
	params, err := firewall.GetDefaultFirewallParams([]firewall.FirewallVariant{input.Variant}...)
	if err != nil {
		return FirewallParamsOutput{}, err
	}

	output := FirewallParamsOutput{
		FirewallParamsInput: input,
		FirewallParams:      *params,
	}

	return output, nil
}
