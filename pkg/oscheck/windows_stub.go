//go:build !windows
// +build !windows

package oscheck

import "github.com/shawnstephens517/go_enumer/pkg/logging"

// WinCheck is a stub implementation for non-Windows builds.
func WinCheck() []logging.LogEntry {
	return []logging.LogEntry{}
}
