package utils

import (
	"context"
	"fmt"
	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
	"github.com/abklabs/svmkit/pkg/runner"
	p "github.com/pulumi/pulumi-go-provider"
	"time"
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

	pcb := func(filename string, copied int, size int, start time.Time) {
		logger := p.GetLogger(ctx)
		elapsed := time.Since(start).Seconds()
		speed := float64(copied) / elapsed // bytes/sec
		if size != 0 {
			msg := "%s: %2.2d%% (%d bytes - %0.3f MB/s)\n"
			logger.InfoStatus(
				fmt.Sprintf(
					msg,
					filename,
					100*copied/size,
					copied,
					speed/1024/1024))
		} else {
			msg := "%s: (%d bytes - %0.3f MB/s)\n"
			logger.InfoStatus(
				fmt.Sprintf(
					msg,
					filename,
					copied,
					speed/1024/1024))
		}
	}

	r := runner.NewRunner(client, command)

	handler := MakePulumiLogger(ctx)

	if err := r.Run(ctx, handler, pcb); err != nil {
		return err
	}

	return nil
}
