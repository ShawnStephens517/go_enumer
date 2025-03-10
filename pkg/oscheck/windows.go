// +build windows
package oscheck

import (
	"context"
	"fmt"
	"os/exec"
	"time"
  "golang.org/x/sys/windows"
)

func checkWSL() {
	fmt.Println("Capture if WSL Enabled???")
	cmd := exec.Command("wsl", "--list", "--verbose")
	err := cmd.Run()
	if err != nil {
		fmt.Println("WSL is not enabled")
	}
	//Code Me Here
}

func roastMe() {
	fmt.Println("Can I kerberoast this box???")
	//TODO...Code Me Here
}

func schTask() {
	fmt.Println("What Scheduled Tasks can I see???")
  	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "Get-ChildItem "c:\windows\system32\tasks"")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Unaccessible as current user: %v", err)
	}
	return string(output), nil
  
  
  
  //Code Me Here
}

func gettheBasics() (string, error) {
	fmt.Println("What is the Hostname, Scriptblock Logging, etc...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "hostname")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %v", err)
	}
	return string(output), nil
	//Code me Here
}

func WinCheck() {
	checkWSL()
	roastMe()
	schTask()

}
