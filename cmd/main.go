package main

import (
	"fmt"
	"os"
	"runtime"

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
[] Add a Timer Function that staggers execution within a set min max range
*/

var rootCmd = &cobra.Command{
	Use:   "goenumer",
	Short: "A cross-platform enumeration tool for Purple-Teaming",
	Long: `goenumer is a lightweight tool that mimics some functionallity found in WinPeas/LinPeas.

Flags:
-eg, --exfilGit: Exfil results of the enumeration to a Git Repo.
-egP, --gitPort: Specify a port to push to if not HTTP/S based.
--eh, --exfilHTTP: Exfil to a Web Server.
--ehP, --httpPort: Specify port for the receiving Web Server.
CZD, --ChariZarD: Runs all nefarious functions on the machine. This is a designed to be extremely noisy.
-OA, --outputall: Outputs results to CSV, HTML, JSON files.
-OH, --outputHTML: Outputs results to HTML.
-OJ, --outputJSON: Outputs results to JSON.
-OC, --outputCSV: Outputs results to CSV.
-fn, --filename: Base name for the results file/s.
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
	rootCmd.Flags().StringP("exfilGit", "eg", "", "Exfil results to Git Repo")
	rootCmd.Flags().StringP("exfilHTTP", "eh", "", "Exfil results to Web Server")
	rootCmd.Flags().StringP("outputall", "OA", "", "Export to CSV, HTML, JSON")
	rootCmd.Flags().StringP("outputHTML", "OH", "", "Export to HTML Only")
	rootCmd.Flags().StringP("outputJSON", "OJ", "", "Export to JSON Only")
	rootCmd.Flags().StringP("outputCSV", "OC", "", "Export to CSV Only")
	rootCmd.Flags().StringP("filename", "fn", "enumerresults"+time.now(), "Base name for the Output files")
	rootCmd.Flags().IntP("gitPort", "egP", 443, "Non Standard port for Git operations. EX:5000")
	rootCmd.Flags().IntP("httpPort", "ehP", 80, "Specify Web Server receiving the results. EX: 443 or 8080")
}
func main() {
	fmt.Println("Call checker based on OS type")
	//Code Here for Logging output to a file & HTTP file
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
