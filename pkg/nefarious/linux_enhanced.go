//go:build linux
// +build linux

package nefarious

import (
	"fmt"
	"github.com/shawnstephens517/go_enumer/nefarious/linuxactions
)

func linux_Enchaned() {
	fmt.Println("Linux Enhanced Functionality")
	fmt.Print("This is a placeholder for the Linux Nasty Checks. Not safe for OSCP use.\n Discretion used in production tests!")
}

func Linux_Nefarious() {
	linux_Enchaned()
	linuxactions.Reciever()
}
