//go:build windows
// +build windows

package windowschecks

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func CheckLoggingInfo() (string, error) {
	var result string

	// --- Powershell ScriptBlock Logging in CURRENT_USER ---
	key1, err := registry.OpenKey(registry.CURRENT_USER, `Software\Policies\Microsoft\Windows\Powershell`, registry.QUERY_VALUE)
	if err != nil {
		result += "Unable to open CURRENT_USER Powershell key\n"
	} else {
		defer key1.Close()
		if value, _, err := key1.GetStringValue("PowershellScriptBlocking"); err != nil {
			result += "PowershellScriptBlocking not set in CURRENT_USER\n"
		} else {
			result += fmt.Sprintf("CURRENT_USER PowershellScriptBlocking: %s\n", value)
		}
	}

	// --- WoW6432Node Powershell Logging in CURRENT_USER ---
	key2, err := registry.OpenKey(registry.CURRENT_USER, `WoW6432Node\Software\Policies\Microsoft\Windows\Powershell`, registry.QUERY_VALUE)
	if err != nil {
		result += "Unable to open WoW6432Node Powershell key in CURRENT_USER\n"
	} else {
		defer key2.Close()
		if value, _, err := key2.GetStringValue("PowershellScriptBlocking"); err != nil {
			result += "PowershellScriptBlocking not set in WoW6432Node CURRENT_USER\n"
		} else {
			result += fmt.Sprintf("WoW6432Node PowershellScriptBlocking: %s\n", value)
		}
	}

	// --- Event Forwarding Subscription Manager (LOCAL_MACHINE) ---
	key3, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\EventLog\EventForwarding\SubscriptionManager`, registry.QUERY_VALUE)
	if err != nil {
		result += "Unable to open EventForwarding SubscriptionManager key\n"
	} else {
		defer key3.Close()
		if value, _, err := key3.GetStringValue(""); err != nil {
			result += "No default value in EventForwarding SubscriptionManager\n"
		} else {
			result += fmt.Sprintf("EventForwarding SubscriptionManager: %s\n", value)
		}
	}

	return result, nil
}
