//go:build windows
// +build windows

package windowschecks

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func ProtectionChecks() (string, error) {
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
