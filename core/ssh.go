package core

import (
	"fmt"
	"os"
	"os/exec"
)

// ConnectSSH connects to the SSH server using the provided connection string
func ConnectSSH(connectionString string) error {
	cmd := exec.Command("ssh", connectionString)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute SSH command: %v", err)
	}

	return nil
}
