package agave

import (
	"fmt"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	// Path to the sysctl configuration file
	sysctlConfPath = "/etc/sysctl.d/21-solana-validator.conf"

	// Template for sysctl configuration settings
	sysctlTemplate = `# Increase UDP buffer sizes
net.core.rmem_default = 134217728
net.core.rmem_max = 134217728
net.core.wmem_default = 134217728
net.core.wmem_max = 134217728
# Increase memory mapped files limit
vm.max_map_count = 1000000
# Increase number of allowed open file descriptors
fs.nr_open = 1000000
vm.swappiness=1
`
)

// The set of arguments for creating a key pair component resource.
type KeyPairArgs struct {
	// The identification name of the key pair within pulumi
	Name string
	// The private key to write to the file
	Key string
	// The user to own the file
	User string
	// The name of the file to write the key to
	FileName string
}

// WriteKeyPair writes a key pair to a specified file on a remote machine and sets the file ownership.
//
// This function creates a command that writes the provided private key to a file in the user's home directory
// and changes the ownership of the file to the specified user. The command is then executed on the remote machine.
//
// Parameters:
//   - ctx: The Pulumi context to use for resource registration and management.
//   - conn: The remote connection information to use for executing the command.
//   - args: The arguments required to create the key pair, including the name, key, user, and file name.
//
// Returns:
//   - *remote.Command: The remote command resource that was created to write the key pair.
//   - error: An error object if there was an issue creating the remote command.
func WriteKeyPair(ctx *pulumi.Context, conn remote.ConnectionInput, args KeyPairArgs) (*remote.Command, error) {
	targetDir := fmt.Sprintf("/home/%s", args.User)
	filePath := fmt.Sprintf("%s/%s", targetDir, args.FileName)
	cmd := fmt.Sprintf("echo %s | sudo tee %s > /dev/null && sudo chown %s:%s %s", args.Key, filePath, args.User, args.User, filePath)

	// Create the command to write the key pair to the file and set the ownership
	return remote.NewCommand(ctx, args.Name, &remote.CommandArgs{
		Create:     pulumi.String(cmd),
		Connection: conn,
	})
}

// ConfigureSysctl configures the sysctl settings on a remote machine.
//
// This function creates and executes commands to write sysctl configuration settings to a file
// and then applies those settings on the remote machine.
//
// Parameters:
//   - ctx: The Pulumi context to use for resource registration and management.
//   - conn: The remote connection information to use for executing the commands.
//
// Returns:
//   - error: An error object if there was an issue creating or executing the remote commands.
func ConfigureSysctl(ctx *pulumi.Context, conn remote.ConnectionInput) error {
	// Create the command to write the template content to the file
	writeCommand := fmt.Sprintf("echo '%s' | sudo tee %s > /dev/null", sysctlTemplate, sysctlConfPath)

	// Execute the command to write the file
	_, err := remote.NewCommand(ctx, "writeSysctlConfig", &remote.CommandArgs{
		Create:     pulumi.String(writeCommand),
		Connection: conn,
	})
	if err != nil {
		return fmt.Errorf("failed to write sysctl config: %w", err)
	}

	// Create the command to apply the sysctl settings
	applyCommand := fmt.Sprintf("sudo sysctl -p %s", sysctlConfPath)

	// Execute the command to apply the sysctl settings
	_, err = remote.NewCommand(ctx, "applySysctlConfig", &remote.CommandArgs{
		Create:     pulumi.String(applyCommand),
		Connection: conn,
	})
	if err != nil {
		return fmt.Errorf("failed to apply sysctl config: %w", err)
	}

	return nil
}

// SetUfwRules configures the UFW firewall rules on a remote machine.
//
// This function creates and executes commands to install UFW and set UFW firewall rules on the remote machine.
//
// Parameters:
//   - ctx: The Pulumi context to use for resource registration and management.
//   - conn: The remote connection information to use for executing the commands.
//
// Returns:
//   - error: An error object if there was an issue creating or executing the remote commands.
func SetUfwRules(ctx *pulumi.Context, conn remote.ConnectionInput) error {
	// Command to install UFW
	installUfwCmd := "sudo apt-get update && sudo apt-get install -y ufw"

	// Execute the command to install UFW
	installUfw, err := remote.NewCommand(ctx, "installUfw", &remote.CommandArgs{
		Create:     pulumi.String(installUfwCmd),
		Connection: conn,
	})
	if err != nil {
		return fmt.Errorf("failed to install UFW: %w", err)
	}

	commands := []string{
		"sudo ufw allow 53",
		"sudo ufw allow ssh",
		"sudo ufw allow 8000:8020/tcp",
		"sudo ufw allow 8000:8020/udp",
		"sudo ufw deny 8899",
	}

	var ufwRules []pulumi.Resource
	ufwRules = append(ufwRules, installUfw)
	for i, cmd := range commands {
		ufwRule, err := remote.NewCommand(ctx, fmt.Sprintf("ufwRule%d", i), &remote.CommandArgs{
			Create:     pulumi.String(cmd),
			Connection: conn,
		}, pulumi.DependsOn(ufwRules))
		if err != nil {
			return fmt.Errorf("failed to execute command '%s': %w", cmd, err)
		}
		ufwRules = append(ufwRules, ufwRule)
	}

	// Enable UFW after all rules are set
	_, err = remote.NewCommand(ctx, "enableUfw", &remote.CommandArgs{
		Create:     pulumi.String("sudo ufw --force enable"),
		Connection: conn,
	}, pulumi.DependsOn(ufwRules))
	if err != nil {
		return fmt.Errorf("failed to enable UFW: %w", err)
	}

	return nil
}
