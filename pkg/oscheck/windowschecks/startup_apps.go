//go:build windows
// +build windows

package windowschecks

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func Startup_RegistryLM() (string, error) {
	//Check Startup Registry Locations
	var result string

	startupKeyLM, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("unable to query Local Machine Startup Registry: %v", err)
	}
	defer startupKeyLM.Close()

	names, err := startupKeyLM.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("unable to read Local Machine Startup value names: %v", err)
	}
	result += "Local Machine Run:\n"
	for _, name := range names {
		value, _, err := startupKeyLM.GetStringValue(name)
		if err != nil {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		} else {
			result += fmt.Sprintf("  %s: %s\n", name, value)
		}
	}

	startupKeyRONCE, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("unable to query Local Machine Startup Registry: %v", err)
	}
	defer startupKeyRONCE.Close()

	namesOnce, err := startupKeyRONCE.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("unable to read Local Machine Startup value names: %v", err)
	}
	result += "Local Machine RunOnce:\n"
	for _, name := range namesOnce {
		value, _, err := startupKeyRONCE.GetStringValue(name)
		if err != nil {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		} else {
			result += fmt.Sprintf("  %s: %s\n", name, value)
		}
	}
	return result, nil

}

func Startup_RegistryCU() (string, error) {
	//Check Startup Registry Locations
	var result string

	startupKeyCU, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("unable to query Current User Startup Registry: %v", err)
	}
	defer startupKeyCU.Close()

	namesOnce, err := startupKeyCU.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("unable to read Current User Startup value names: %v", err)
	}
	result += "Current User Run:\n"
	for _, name := range namesOnce {
		value, _, err := startupKeyCU.GetStringValue(name)
		if err != nil {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		} else {
			result += fmt.Sprintf("  %s: %s\n", name, value)
		}
	}

	startupKeyCUONCE, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("unable to query Local Machine Startup Registry: %v", err)
	}
	defer startupKeyCUONCE.Close()

	names, err := startupKeyCUONCE.ReadValueNames(0)
	if err != nil {
		return "", fmt.Errorf("unable to read Local Machine Startup value names: %v", err)
	}
	result += "Local Machine RunOnce:\n"
	for _, name := range names {
		value, _, err := startupKeyCUONCE.GetStringValue(name)
		if err != nil {
			result += fmt.Sprintf("  %s: error retrieving value\n", name)
		} else {
			result += fmt.Sprintf("  %s: %s\n", name, value)
		}
	}
	return result, nil
}
