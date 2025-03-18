package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shawnstephens517/go_enumer/pkg/logging"
	"github.com/shawnstephens517/go_enumer/pkg/oscheck"
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
[x] Output to multi-format (CSV, TXT, HTML, xlsx?)
[] Comments and Readme update...
[x] Add a Timer Function that staggers execution within a set min max range
*/

var rootCmd = &cobra.Command{
	Use:   "goenumer",
	Short: "A cross-platform enumeration tool for Purple-Teaming",
	Long: `goenumer is a lightweight tool that mimics some functionallity found in WinPeas/LinPeas.

Flags:
-G, --exfilGit: Exfil results of the enumeration to a Git Repo.
-g, --gitPort: Specify a port to push to if not HTTP/S based.
-E, --exfilHTTP: Exfil to a Web Server.
-P, --httpPort: Specify port for the receiving Web Server.
--ChariZarD: Runs all nefarious functions on the machine. This is a designed to be extremely noisy.
-O, --outputall: Outputs results to CSV, HTML, JSON files.
-H, --outputHTML: Outputs results to HTML.
-J, --outputJSON: Outputs results to JSON.
-C, --outputCSV: Outputs results to CSV.
-f, --filename: Base name for the results file/s.
-T, --timeroff: Use timer function to randomly stagger checks
--sourceIP: Reverse Shell connect back IP
--sourcePort: Port to connect back too. Defaults 9090
--reverse: Calls the shell function to initiate a generic Reverse Shell. Source IP flag required!
`,

	Run: func(cmd *cobra.Command, args []string) {
		// Call your enumeration functions here based on OS
		fmt.Println("Enumerating the Systems...")
		var entries []logging.LogEntry

		switch runtime.GOOS {
		case "windows":
			fmt.Println("Running Windows Enumer")
			entries = oscheck.WinCheck()

		case "linux":
			fmt.Println("Running Linux Enumer")
			entries = oscheck.LinCheck()

		//case "darwin":
		//fmt.Println("Running MacOS Enumer")
		//entries = oscheck.MacCheck()

		default:
			fmt.Printf("Unsupported OS: %s\n", runtime.GOOS)
			os.Exit(1)
		}
		//Write the results to multiple formats concurrently
		logging.WriteAllFormats(entries, "Posible PrivEsc")
	},
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Displays help info")
	rootCmd.Flags().BoolP("reverse", "", false, "Use the Reverse Shell Function")
	rootCmd.Flags().BoolP("timeroff, T", "", true, "Use the Timer Function")
	rootCmd.Flags().StringP("exfilGit", "G", "", "Exfil results to Git Repo")
	rootCmd.Flags().StringP("exfilHTTP", "E", "", "Exfil results to Web Server")
	rootCmd.Flags().StringP("outputall", "A", "", "Export to CSV, HTML, JSON")
	rootCmd.Flags().StringP("outputHTML", "H", "", "Export to HTML Only")
	rootCmd.Flags().StringP("outputJSON", "J", "", "Export to JSON Only")
	rootCmd.Flags().StringP("outputCSV", "C", "", "Export to CSV Only")
	rootCmd.Flags().StringP("filename", "f", "enumerresults"+time.Now().Format("20060102150405"), "Base name for the Output files")
	rootCmd.Flags().StringP("sourceIP", "", "", "Reverse Shell call back IP.\nEnsure Listener Started on callback host to receive the connection")
	rootCmd.Flags().IntP("gitPort", "g", 443, "Non Standard port for Git operations. EX:5000")
	rootCmd.Flags().IntP("httpPort", "P", 80, "Specify Web Server receiving the results. EX: 443 or 8080")
	rootCmd.Flags().IntP("sourcePort", "", 9090, "Specify the port the reverse shell should connect")
}
func main() {
	fmt.Println("Call checker based on OS type")
	//Code Here for Logging output to a file & HTTP file
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
