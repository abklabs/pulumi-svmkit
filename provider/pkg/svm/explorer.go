package svm

import (
	"github.com/abklabs/pulumi-svmkit/pkg/ssh"
)

// ExplorerArgs represents the arguments required to configure a Explorer.
type ExplorerArgs struct {
	// Connection holds the SSH connection details needed to access the Explorer.
	Connection ssh.Connection `pulumi:"connection"`
}
