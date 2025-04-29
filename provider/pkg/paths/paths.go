package paths

import (
	"context"

	"github.com/abklabs/svmkit/pkg/paths"
)

type PathsInput struct {
}

type PathsOutput struct {
    PathsInput
	paths.Paths
}

type GetDefaultPaths struct{}

func (GetDefaultPaths) Call(ctx context.Context, input PathsInput) (PathsOutput, error) {
    p, err := paths.NewDefaultPaths()
    if err != nil {
        return PathsOutput{}, err
    }

    output := PathsOutput{
        PathsInput: input,
		Paths: *p,
    }

    return output, nil
}
