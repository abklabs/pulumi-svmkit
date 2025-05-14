package machine

import (
	"context"

	"github.com/abklabs/pulumi-svmkit/pkg/utils"
	"github.com/abklabs/svmkit/pkg/machine"
)

type Machine struct{}

type MachineArgs struct {
	utils.RunnerArgs
	machine.Machine
}

type MachineState struct {
	MachineArgs
}

func (Machine) Create(ctx context.Context, name string, input MachineArgs, preview bool) (string, MachineState, error) {
	state := MachineState{
		MachineArgs: input,
	}

	if preview {
		return name, state, nil
	}

	command := &machine.CreateCommand{
		Machine: input.Machine,
	}

	if err := utils.RunnerHelper(ctx, input.RunnerArgs, command); err != nil {
		return "", MachineState{}, err
	}

	return name, state, nil
}
