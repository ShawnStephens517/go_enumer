//go:build windows
// +build windows

package windowschecks

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func SchTask() (string, error) {
	fmt.Println("What Scheduled Tasks can I see???")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "Get-ChildItem 'c:\\windows\\system32\\tasks'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Current user can't enumerate scheduled tasks: %v", err)
	}
	return string(output), nil
}
