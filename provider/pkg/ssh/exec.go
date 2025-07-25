package ssh

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
)

func Exec(ctx context.Context, client *ssh.Client, command string) (stdout string, stderr string, err error) {
	// Create a new session for running the command

	var execSession *ssh.Session
	execSession, err = client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create SSH session: %w", err)
	}

	defer func() {
		if closeErr := execSession.Close(); closeErr != io.EOF {
			err = errors.Join(closeErr)
		}
	}()

	var stdoutPipe io.Reader
	stdoutPipe, err = execSession.StdoutPipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	var stderrPipe io.Reader
	stderrPipe, err = execSession.StderrPipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutDone := make(chan struct{})
	stderrDone := make(chan struct{})

	go func() {
		if _, err = io.Copy(&stdoutBuf, stdoutPipe); err != nil {
			fmt.Printf("error copying stdout: %v\n", err)
		}
		close(stdoutDone)
	}()

	go func() {
		if _, err = io.Copy(&stderrBuf, stderrPipe); err != nil {
			fmt.Printf("error copying stderr: %v\n", err)
		}
		close(stderrDone)
	}()

	bashCommand := fmt.Sprintf("/bin/sh -c '%s'", command)
	if err = execSession.Start(bashCommand); err != nil {
		return "", "", fmt.Errorf("failed to start command: %w", err)
	}

	if err = execSession.Wait(); err != nil {
		return "", "", fmt.Errorf("command execution failed: %w", err)
	}

	<-stdoutDone
	<-stderrDone

	return stdoutBuf.String(), stderrBuf.String(), nil
}
