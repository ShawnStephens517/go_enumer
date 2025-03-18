//go:build windows
// +build windows

package windowschecks

import (
	"fmt"
	"os"
)

// Initial. Need to Iterate over all users to find out read access
func CheckUserDirs() (string, error) {
	var result string
	files, err := os.ReadDir("C:\\Users\\")
	if err != nil {
		return "", fmt.Errorf("Unable to enumerate User Directories: %v", err)
	}
	result += "Local Machine RunOnce:\n"

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return result, nil
}
