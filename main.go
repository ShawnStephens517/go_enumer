package main


/*
TODO
[] Cobra for CMD flagging
[] Viper for Importing offline checking (Similar to Blood/Azurehound functionallity, but local machine)
[] Vendor any external dependancies for offline compile.
[] Ensure Garbleable for hiding :)
[] Pipelining (Action Build unobfuscated-Binary)
[] Look at Go-native implimentation of some of Win/LinPeas functionallity
[] Output to multi-format (CSV, TXT, HTML, xlsx?)
[] Comments and Readme update...
*/

func main(){
  fmt.Println("Call checker based on OS type")
  //Code Here for Logging output to a file & HTTP file

  osType := runtime.GOOS
  switch os {
  case "windows":
    fmt.Println("Running Windows Enumer")
    oscheck.WinCheck()

  case "linux":
    fmt.Println("Running Linux Enumer")
    oscheck.LinCheck()
  }
  
}
