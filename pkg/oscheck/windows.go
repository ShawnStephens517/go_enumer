//go:build windows
// +build windows

package oscheck

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/shawnstephens517/go_enumer/pkg/logging"
	"github.com/shawnstephens517/go_enumer/pkg/utils"
	"golang.org/x/sys/windows/registry"
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
func protectionChecks() (string, error) {
	var result string

	//Check LSA Protections
	//Check Credential Guard
	lsaKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\LSA`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("unable to query LSA protections: %v", err)
	}
	defer lsaKey.Close()

	names, err := lsaKey.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("unable to read LSA value names: %v", err)
	}
	result += "LSA Protections:\n"
	for _, name := range names {
		value, _, err := lsaKey.GetStringValue(name)
		if err != nil {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		} else {
			result += fmt.Sprintf("  %s: %s\n", name, value)
		}
	}

	//UAC Settings. We want to capture the Values here, and also specifically check for EnableLUA =1. If so, display that there are some UAC settings enabled. TODO
	uacKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("unable to query UAC settings: %v", err)
	}
	defer uacKey.Close()

	names, err = uacKey.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("unable to read UAC key value names: %v", err)
	}
	result += "\nUAC Settings:\n"
	for _, name := range names {
		if value, _, err := uacKey.GetIntegerValue(name); err == nil {
			result += fmt.Sprintf("  %s: %d\n", name, value)
		} else if svalue, _, err := uacKey.GetStringValue(name); err == nil {
			result += fmt.Sprintf("  %s: %s\n", name, svalue)
		} else {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		}
	}
	return result, nil
}
func getUserRunMRU() (string, error) {
	var result string

	// Open the USERS hive
	usersKey, err := registry.OpenKey(registry.USERS, "", registry.READ)
	if err != nil {
		return "", fmt.Errorf("cannot open USERS key: %w", err)
	}
	defer usersKey.Close()

	// Read all subkey names (typically SIDs)
	sids, err := usersKey.ReadSubKeyNames(-1)
	if err != nil {
		return "", fmt.Errorf("cannot read subkeys for USERS: %w", err)
	}

	// Iterate over each SID subkey
	for _, sid := range sids {
		runMRUPath := fmt.Sprintf("%s\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\RunMRU", sid)
		runMRUKey, err := registry.OpenKey(registry.USERS, runMRUPath, registry.QUERY_VALUE)
		if err != nil {
			// Skip if the RunMRU key is not present for this SID.
			continue
		}
		// Read all value names in the RunMRU key
		valNames, err := runMRUKey.ReadValueNames(0)
		if err != nil {
			runMRUKey.Close()
			continue
		}
		result += fmt.Sprintf("SID %s RunMRU:\n", sid)
		// Iterate over each value and capture its contents
		for _, name := range valNames {
			value, _, err := runMRUKey.GetStringValue(name)
			if err != nil {
				result += fmt.Sprintf("  %s: error retrieving value\n", name)
			} else {
				result += fmt.Sprintf("  %s: %s\n", name, value)
			}
		}
		runMRUKey.Close()
	}

	return result, nil
}
func getRDPSavedConnections() (string, error) {
	var result string

	// Open the HKEY_USERS hive.
	usersKey, err := registry.OpenKey(registry.USERS, "", registry.READ)
	if err != nil {
		return "", fmt.Errorf("cannot open HKEY_USERS: %v", err)
	}
	defer usersKey.Close()

	sids, err := usersKey.ReadSubKeyNames(-1)
	if err != nil {
		return "", fmt.Errorf("cannot read HKEY_USERS subkeys: %v", err)
	}

	result += "RDP Saved Connections (HKU):\n"
	for _, sid := range sids {
		// Build the path to the RDP settings for this SID.
		rdpPath := fmt.Sprintf(`%s\Software\Microsoft\Terminal Server Client\Default`, sid)
		rdpKey, err := registry.OpenKey(registry.USERS, rdpPath, registry.QUERY_VALUE)
		if err != nil {
			result += fmt.Sprintf("  %s: RDP key not found or inaccessible\n", sid)
			continue
		}
		value, _, err := rdpKey.GetStringValue("MRU0")
		if err != nil {
			result += fmt.Sprintf("  %s: MRU0 not found\n", sid)
		} else {
			result += fmt.Sprintf("  %s: %s\n", sid, value)
		}
		rdpKey.Close()
	}
	return result, nil
}

func checkSSHKeys() (string, error) {
	var result string

	// --- PuTTY SSH Host Keys ---
	puttyKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\SshHostKeys`, registry.QUERY_VALUE)
	if err == nil {
		defer puttyKey.Close()
		props, err := puttyKey.ReadValueNames(0)
		if err == nil && len(props) > 0 {
			result += "PuTTY SSH Host Keys found:\n"
			for _, prop := range props {
				result += fmt.Sprintf("  %s\n", prop)
			}
		} else {
			result += "No PuTTY SSH Host Keys found.\n"
		}
	} else {
		result += "No PuTTY SSH Host Keys found (error opening key).\n"
	}

	// --- OpenSSH Keys ---
	opensshKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\OpenSSH\Agent\Keys`, registry.QUERY_VALUE)
	if err == nil {
		defer opensshKey.Close()
		result += "OpenSSH keys found. See: https://github.com/ropnop/windows_sshagent_extract\n"
	} else {
		result += "No OpenSSH keys found.\n"
	}

	return result, nil
}
func accounting() (string, error) {
	var result string

	//Password Policy Check
	//Cached Credentials
	winlogonKey, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Winlogon", registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("Unable to query WinLogon Cached Creds: %v", err)
	}
	defer winlogonKey.Close()

	names, err := winlogonKey.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("Unable to determine Winlogon Cached Creds %v", err)
	}
	result += "Winlogon Cached Credentials:\n"

	for _, name := range names {
		value, _, err := winlogonKey.GetStringValue(name)
		if err != nil {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		} else {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	//Saved RDP Connection Info
	rdpKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Terminal Server Client\Default`, registry.QUERY_VALUE)
	if err == nil {
		defer rdpKey.Close()
		rdpNames, err := rdpKey.ReadValueNames(0)
		if err == nil && len(rdpNames) > 0 {
			result += "\nSaved RDP Connections:\n"
			for _, name := range rdpNames {
				value, _, err := rdpKey.GetStringValue(name)
				if err != nil {
					result += fmt.Sprintf("  %s: error retrieving value\n", name)
				} else {
					result += fmt.Sprintf("  %s: %s\n", name, value)
				}
			}
		} else {
			result += "\nSaved RDP Connections: no values found or error reading values\n"
		}
	} else {
		result += "\nSaved RDP Connections: key not found or error opening\n"
	}

	rdpConnections, err := getRDPSavedConnections()
	if err != nil {
		result += fmt.Sprintf("\nError retrieving RDP saved connections: %v\n", err)
	} else {
		result += "\n" + rdpConnections
	}

	//Stored Putty Creds/Sessions
	puttyKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\Sessions`, registry.QUERY_VALUE)
	if err == nil {
		defer puttyKey.Close()
		sessions, err := puttyKey.ReadSubKeyNames(-1)
		if err == nil && len(sessions) > 0 {
			result += "\nPuTTY Sessions:\n"
			for _, session := range sessions {
				result += fmt.Sprintf("  Session: %s\n", session)
			}
		} else {
			result += "\nPuTTY Sessions: no sessions found or error reading sessions\n"
		}
	} else {
		result += "\nPuTTY Sessions: key not found or error opening\n"
	}

	//SSH Keys & Known Hosts
	sshKeys, err := checkSSHKeys()
	if err != nil {
		result += fmt.Sprintf("\nError retrieving SSH key info: %v\n", err)
	} else {
		result += "\n" + sshKeys
	}
	//cmdkey.exe /list
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "cmdkey.exe", "/list")
	cmdOutput, err := cmd.CombinedOutput()
	if err == nil {
		result += "\nStored Credentials (cmdkey.exe):\n" + string(cmdOutput)
	} else {
		result += fmt.Sprintf("\nStored Credentials (cmdkey.exe): error executing command: %v\n", err)
	}

	//Recently Run actions. Determine if system may have been pre-compromised using Win + R or if there are some sketchy actions ran previously that may assist in our efforts.
	userRunMRU, err := getUserRunMRU()
	if err != nil {
		return "", fmt.Errorf("failed to get RunMRU details: %w", err)
	}
	result += "\nRecently Run Actions:\n" + userRunMRU

	//Check 2
	RunMRUkeycu, err := registry.OpenKey(registry.CURRENT_USER, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\RunMRU", registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("Unable to query Recently Ran through Win+R. HKCU %v", err)
	}
	defer RunMRUkeycu.Close()
	result += "\nRecently Run Actions:\n"

	return result, nil
}

func roastAble() (string, error) {
	fmt.Println("Can I kerberoast this box???")
	return "Kerberoasting check not implemented yet", nil
	//TODO: Code Me Here
	/*
	 	fmt.Println("Retrieving SPNs")
	  	spns, err := exec.Command ("powershell.exe", "get-adobject", "-filter", "{serviceprincipalname -like &serviceaccountfromConfigFile(viper library)}", "-properties", "serviceprincipalname")
	   	output, err := spns.Output()
	    	if err != nil {
	     		return "", fmt.Errorf("Unable to obtain Service Principal Names using the supplied Service Account", err)
	      	}
	        return string(output), nil
	*/
}

func schTask() (string, error) {
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

func checkLoggingInfo() (string, error) {
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
	if output, err := roastAble(); err != nil {
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
	if output, err := schTask(); err != nil {
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
	if output, err := protectionChecks(); err != nil {
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
	if output, err := accounting(); err != nil {
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
	if output, err := checkLoggingInfo(); err != nil {
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
