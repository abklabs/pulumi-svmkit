package utils

import (
	"context"
	"fmt"

	"github.com/abklabs/svmkit/pkg/runner"
	"github.com/abklabs/svmkit/pkg/ssh"
)

func RunnerHelper(ctx context.Context, connection ssh.Connection, command runner.Command) error {
	if err := command.Check(); err != nil {
		return fmt.Errorf("failed to check validator config: %w", err)
	}

	r := runner.NewRunner(connection, command)

	if err := r.Run(ctx); err != nil {
		return fmt.Errorf("failed to install validator: %w", err)
	}

	return nil
}
