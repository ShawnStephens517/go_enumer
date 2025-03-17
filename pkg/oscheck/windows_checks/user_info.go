// +build windows

package windows_checks

import (
	"context"
	"fmt"
  "os"
	"os/exec"
	"time"
  "path/filepath"

  "golang.org/x/sys/windows/registry"
)
//Initial. Need to Iterate over all users to find out read access
func CheckUserDirs() (string, error){
  files, err := os.ReadDir("C:\Users\")
   if err != nil {
		  return "", fmt.Errorf("Unable to enumerate User Directories: %v", err)
	  }

  for _, file := range files{
    fmt.Println(file.Name())
    }
}
