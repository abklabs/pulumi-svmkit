package svm

import (
	"github.com/abklabs/svmkit/pkg/ssh"
	"github.com/abklabs/svmkit/pkg/validator"
)

// ValidatorArgs represents the arguments required to configure a validator.
type ValidatorArgs struct {
	// Connection holds the SSH connection details needed to access the validator.
	Connection ssh.Connection `pulumi:"connection"`
	// Version specifies the version of the validator to be used.
	// It is an optional field and can be omitted if not needed.
	// If omitted, the latest version of the validator will be installed.
	Version validator.Version `pulumi:"version,optional"`
}
