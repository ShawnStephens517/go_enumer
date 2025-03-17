//go:build windows
// +build windows

package oscheck

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/shawnstephens517/go_enumer/pkg/logging"
	"github.com/shawnstephens517/go_enumer/pkg/oscheck/windowschecks"
	"github.com/shawnstephens517/go_enumer/pkg/utils"
)

func checkWSL() (string, error) {
	fmt.Println("Capture if WSL Enabled???")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "wsl", "--list", "--verbose")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("WSL check failed: %v", err)
	}
	return string(output), nil
}

func getBasicInfo() (string, error) {
	fmt.Println("What is the Hostname, Scriptblock Logging, etc...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "hostname")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %v", err)
	}
	//Code me Here
	return string(output), nil
}

func WinCheck() []logging.LogEntry {
	var entries []logging.LogEntry
	now := time.Now()

	// Basic Info
	if output, err := getBasicInfo(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Basic Info",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Basic Info",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// WSL Check
	if output, err := checkWSL(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "WSL Check",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "WSL Check",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Kerberoasting (RoastAble)
	if output, err := windowschecks.RoastAble(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Kerberoasting Check",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Kerberoasting Check",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Scheduled Tasks
	if output, err := windowschecks.SchTask(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Scheduled Tasks",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Scheduled Tasks",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Protection Checks
	if output, err := windowschecks.ProtectionChecks(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Protection Checks",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Protection Checks",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Accounting
	if output, err := windowschecks.Accounting(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Accounting",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Accounting",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Logging Info
	if output, err := windowschecks.CheckLoggingInfo(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Logging Info",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Logging Info",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Startup Apps Registry HKLM
	if output, err := windowschecks.Startup_RegistryLM(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Logging Info",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Logging Info",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()

	// Startup Apps Registry HKCU
	if output, err := windowschecks.Startup_RegistryCU(); err != nil {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Logging Info",
			Message:   fmt.Sprintf("Error: %v", err),
			Data:      "",
		})
	} else {
		entries = append(entries, logging.LogEntry{
			Timestamp: now,
			CheckName: "Logging Info",
			Message:   "Captured",
			Data:      output,
		})
	}
	utils.WaitForNextCheck()
	// Additional checks can be added here in the same manner.
	return entries
}
