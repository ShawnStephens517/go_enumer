//go:build linux
// +build linux

package oscheck

import (
	"fmt"
)

func LinCheck() {
	fmt.Println("Gathering basic Linux info")
}
