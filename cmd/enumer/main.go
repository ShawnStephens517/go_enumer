package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shawnstephens517/go_enumer/pkg/oscheck"
	"github.com/shawnstephens517/go_enumer/pkg/logging"
	"github.com/spf13/cobra"
)

/*
TODO
[x] Cobra for CMD flagging
[x] Viper for Importing offline checking (Similar to Blood/Azurehound functionallity, but local machine)
[] Vendor any external dependancies for offline compile.
[] Ensure Garbleable for hiding :)
[] Pipelining (Action Build unobfuscated-Binary)
[] Look at Go-native implimentation of some of Win/LinPeas functionallity
[] Output to multi-format (CSV, TXT, HTML, xlsx?)
[] Comments and Readme update...
*/

var rootCmd = &cobra.Command{
	Use:   "enumtool",
	Short: "A cross-platform enumeration tool for Purple-Teaming",
	Run: func(cmd *cobra.Command, args []string) {
		// Call your enumeration functions here based on OS
	},
}

func main() {
	fmt.Println("Call checker based on OS type")
	//Code Here for Logging output to a file & HTTP file
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	osType := runtime.GOOS
	switch osType {
	case "windows":
		fmt.Println("Running Windows Enumer")
		oscheck.WinCheck()

	case "linux":
		fmt.Println("Running Linux Enumer")
		oscheck.LinCheck()
	default:
		fmt.Printf("Unsupported OS: %s\n", osType)
	}

}
