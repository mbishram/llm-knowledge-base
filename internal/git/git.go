package git

import (
	"fmt"
	"os/exec"
)

func Commit(dataDir, message string) error {
	// Add changes
	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = dataDir
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to git add: %v", err)
	}

	// Commit changes
	commitCmd := exec.Command("git", "commit", "-m", message)
	commitCmd.Dir = dataDir
	// We don't necessarily error if there's nothing to commit
	_ = commitCmd.Run()

	return nil
}
