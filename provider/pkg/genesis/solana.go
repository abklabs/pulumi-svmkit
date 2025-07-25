package genesis

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/solana/genesis"
	p "github.com/pulumi/pulumi-go-provider"
)

// Solana represents a Pulumi resource for building the genesis ledger for the Solana network.
type Solana struct{}

// SolanaArgs represents the input arguments required to create a Solana genesis resource.
type SolanaArgs struct {
	utils.RunnerArgs
	genesis.Genesis
	IgnoreChangesOnUpdate bool `pulumi:"ignoreChangesOnUpdate,optional"`
}

// SolanaState represents the state of a Solana genesis resource.
type SolanaState struct {
	SolanaArgs
	// Hash is the genesis hash of the Solana ledger.
	Hash string `pulumi:"genesisHash"`
}

// Create sets up the Solana genesis ledger on the specified machine using the provided connection and flags.
//
// Parameters:
// - ctx: The context for the creation operation.
// - name: The name of the resource.
// - input: The input arguments for creating the resource.
// - preview: A boolean indicating whether this is a preview operation.
//
// Returns:
// - The name of the created resource.
// - The state of the created resource.
// - An error if the creation fails.
func (Solana) Create(ctx context.Context, name string, input SolanaArgs, preview bool) (retName string, retState SolanaState, retErr error) {
	state := SolanaState{SolanaArgs: input}

	if preview {
		return name, state, nil
	}

	genesis := input.Genesis
	command := genesis.Create()

	if err := utils.RunnerHelper(ctx, input.RunnerArgs, command); err != nil {
		return "", SolanaState{}, err
	}

	// Establish SSH connection
	connection, err := input.Connection.Dial(ctx)
	if err != nil {
		return "", SolanaState{}, fmt.Errorf("failed to establish SSH connection: %w", err)
	}
	defer func() {
		if closeErr := connection.Close(); closeErr != io.EOF {
			retErr = errors.Join(closeErr)
		}
	}()

	// Execute the command on the remote machine
	stdout, stderr, err := ssh.Exec(ctx, connection, "sudo -i -u sol agave-ledger-tool genesis-hash")
	if err != nil {
		return "", SolanaState{}, fmt.Errorf("failed to execute Solana genesis command: %w, stderr: %s", err, stderr)
	}

	state.Hash = strings.TrimSpace(stdout)

	return name, state, nil
}

func (Solana) Update(ctx context.Context, name string, oldState SolanaState, newInput SolanaArgs, preview bool) (SolanaState, error) {
	logger := p.GetLogger(ctx)

	msg := "genesis configuration may not be modified after initial creation"
	if !newInput.IgnoreChangesOnUpdate {
		return oldState, errors.New(msg)
	} else {
		logger.WarningStatus(msg)
	}

	return oldState, nil
}

func (Solana) Delete(ctx context.Context, name string, props SolanaState) error {
	genesis := props.Genesis
	command := genesis.Delete()

	if err := utils.RunnerHelper(ctx, props.RunnerArgs, command); err != nil {
		return err
	}

	return nil
}
