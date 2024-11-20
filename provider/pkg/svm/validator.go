package svm

import (
	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
)

// ValidatorArgs represents the arguments required to configure a validator.
type ValidatorArgs struct {
	// Connection holds the SSH connection details needed to access the validator.
	Connection ssh.Connection `pulumi:"connection"`
}
