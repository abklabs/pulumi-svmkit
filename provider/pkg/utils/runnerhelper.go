package utils

import (
	"context"
	"fmt"

	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
	"github.com/abklabs/svmkit/pkg/runner"
)

type RunnerArgs struct {
	Connection ssh.Connection `pulumi:"connection"`
}

func RunnerHelper(ctx context.Context, runnerArgs RunnerArgs, command runner.Command) error {
	if err := command.Check(); err != nil {
		return fmt.Errorf("failed to check component config: %w", err)
	}

	client, err := runnerArgs.Connection.Dial(ctx)

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
