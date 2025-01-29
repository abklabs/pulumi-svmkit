package svm

import (
	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
)

// TunerArgs represents the arguments required to configure a Tuner.
type TunerArgs struct {
	// Connection holds the SSH connection details needed to access the Tuner.
	Connection ssh.Connection `pulumi:"connection"`
}
