package svm

import (
	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
)

// WatchtowerArgs represents the arguments required to configure a Watchtower.
type WatchtowerArgs struct {
	// Connection holds the SSH connection details needed to access the Watchtower.
	Connection ssh.Connection `pulumi:"connection"`
}
