package svm

import (
	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
)

// FaucetArgs represents the arguments required to configure a Faucet.
type FaucetArgs struct {
	// Connection holds the SSH connection details needed to access the Faucet.
	Connection ssh.Connection `pulumi:"connection"`
}
