//go:build linux
// +build linux

package oscheck

import (
	"fmt"
)

func CaptureSysInfo() {
	fmt.Println("Gathering basic Linux info")
}
