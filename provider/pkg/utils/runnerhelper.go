package utils

import (
	"context"
	"fmt"

	"github.com/abklabs/svmkit/pkg/runner"
	"github.com/abklabs/svmkit/pkg/ssh"
)

func RunnerHelper(ctx context.Context, connection ssh.Connection, command runner.Command) error {
	if err := command.Check(); err != nil {
		return fmt.Errorf("failed to check component config: %w", err)
	}

	client, err := connection.Dial(ctx)

	if err != nil {
		return fmt.Errorf("failed to dial SSH connection to hosst: %w", err)
	}

	r := runner.NewRunner(client, command)

	handler := MakePulumiLogger(ctx)

	if err := r.Run(ctx, handler); err != nil {
		return err
	}

	return nil
}
