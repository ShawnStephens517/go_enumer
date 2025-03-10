//go:build linux
// +build linux

package linux

import (
	"fmt"
)

func CaptureSysInfo() {
	fmt.Println("Gathering basic Linux info")
}
