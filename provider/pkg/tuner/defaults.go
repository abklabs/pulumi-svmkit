package tuner

import (
	"context"

	"github.com/abklabs/svmkit/pkg/tuner"
)

type TunerParamsInput struct {
    Variant tuner.TunerVariant `pulumi:"variant"`
}

type TunerParamsOutput struct {
    TunerParamsInput
	tuner.TunerParams
}

type GetDefaultTunerParams struct{}

func (GetDefaultTunerParams) Call(ctx context.Context, input TunerParamsInput) (TunerParamsOutput, error) {
    params, err := tuner.GetDefaultTunerParams([]tuner.TunerVariant{input.Variant}...)
    if err != nil {
        return TunerParamsOutput{}, err
    }

    output := TunerParamsOutput{
        TunerParamsInput: input,
		TunerParams: *params,
    }

    return output, nil
}
