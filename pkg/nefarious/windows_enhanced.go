//go:build windows
// +build windows

package nefarious

import (
	"fmt"
)

/*
Since most of these actions require quite a bit of code to perform,
This will serve as the handle/middleware between the logic and the call
at runtime.
*/

func DLLInject() {
	fmt.Println("Attempting an DLL Inject")
	//Code Me Here
}

func ShellPopCalc() {
	fmt.Println("Hooking the Windows API to pop a Calc Instance..\n Check your Logs!!!")
	//Code me Here
}

func ProcessInject(){
	fmt.Println("Attempt Process Injection")
	//Code Here
}

func Voltorb(){
	// Self-Removing function. Silly Pokemon Reference to Self-Destruct
	fmt.Println("Removes itself..."
	//Code Here
}

func  NewDLL() {
	fmt.Println("Creating New DLL")
	//Code Here
}

