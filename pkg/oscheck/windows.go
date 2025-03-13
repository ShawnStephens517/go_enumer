// +build windows
package oscheck

import (
	"context"
	"fmt"
	"os/exec"
	"time"
  	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func checkWSL() {
	fmt.Println("Capture if WSL Enabled???")
	cmd := exec.Command("wsl", "--list", "--verbose")
	err := cmd.Run()
	if err != nil {
		fmt.Println("WSL is not enabled")
	}
	//Code Me Here
}
func protectionChecks(){
	//Check LSA

	lkey, err := registry.OpenKey(registry.LOCAL_MACHINE, 'SYSTEM\CurrentControlSet\Control\LSA', registry.QUERY_VALUE)
	if err != nil{
		return "", fmt.Errorf("Unable to query LSA Protections", err)
	}
	defer. lkey.Close()

	names, err := 1key.ValueNames()
	if err != nil {
		return "", fmt.Errorf("Unable to determine LSA Protections", err)
	}
	
	for _, name := range names {
		sv, _, err := lkey.GetStringValue(name)
		if err != nil {
			return "", fmt.Errorf("Can't determine if LSA protections enabled!", err)
			continue
	}
		fmt.Printf("%s: %s\n", name, value)
	}
	
	//Check Credential Guard should be caught in the previous check
	
	//UAC Settings
}

func accounting(){
	//Password Policy Check
	//Cached Credentials
	cakey, err := registry.OpenKey(registry.LOCAL_MACHINE, 'SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon', registry.QUERY_VALUE)
	if err != nil{
		return "", fmt.Errorf("Unable to query Cached Creds", err)
	}
	defer. cakey.Close()

	names, err := cakey.ValueNames()
	if err != nil {
		return "", fmt.Errorf("Unable to determine Cached Creds", err)
	}
	
	for _, name := range names {
		sv, _, err := cakey.GetStringValue(name)
		if err != nil {
			return "", fmt.Errorf("Can't determine if Cached Creds present", err)
			continue
	}
		fmt.Printf("%s: %s\n", name, value)
	}
	
	//Winlogon Credential Check
	//Saved RDP Connection Info
	//Stored Putty Creds
	//SSH Keys & Known Hosts
	//cmdkey.exe /list
}

func roastAble() {
	fmt.Println("Can I kerberoast this box???")
	//TODO: Code Me Here
	/*
 	fmt.Println("Retrieving SPNs")
  	spns, err := exec.Command ("powershell.exe", "get-adobject", "-filter", "{serviceprincipalname -like &serviceaccountfromConfigFile(viper library)}", "-properties", "serviceprincipalname")
   	output, err := spns.Output()
    	if err != nil {
     		return "", fmt.Errorf("Unable to obtain Service Principal Names using the supplied Service Account", err)
      	}
        return string(output), nil
 	*/
}

func schTask() {
	fmt.Println("What Scheduled Tasks can I see???")
  	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "Get-ChildItem "c:\windows\system32\tasks"")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Unaccessible as current user: %v", err)
	}
	return string(output), nil
  
  
  
  //Code Me Here
}

func checkLoggingInfo(){
	k1, err := registry.OpenKey(registry.CURRENT_USER, 'Software\Policies\Microsoft\Windows\Powershell', registry.QUERY_VALUE)
	if err != nil{
		return "", fmt.Errorf("Unable to query Powershell User key", err)
	}
	defer. k1.Close()

	sv, _, err := k1.GetStringValue("PowershellScriptBlocking")
	if err != nil {
		return "", fmt.Errorf("Unable to query ScriptBlock Value. Is the key present?", err)
	}

	k2, err := registry.OpenKey(registry.CURRENT_USER, 'WoW6432Node\Software\Policies\Microsoft\Windows\Powershell', registry.QUERY_VALUE)
	if err != nil{
		return "", fmt.Errorf("Unable to query Powershell User key", err)
	}
	defer. k2.Close()

	sv, _, err := k2.GetStringValue("PowershellScriptBlocking")
	if err != nil {
		return "", fmt.Errorf("Unable to query ScriptBlock Value. Is the key present?", err)
	}


	k3, err := registry.OpenKey(registry.LOCAL_MACHINE, 'SOFTWARE\Policies\Microsoft\Windows\EventLog\EventForwarding\SubscriptionManager', registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("Logs not being forwarded and/or registry not found", err)
	}
	defer.k3.Close()

	sv, _, err := k3.GetStringValue("")
	if err != nil {
		return "", fmt.Errorf("Unable to query Event Forward. Is this present?", err)
	}
	
	//TODO: return k(x) value or error message

	
}

func gettheBasics() (string, error) {
	fmt.Println("What is the Hostname, Scriptblock Logging, etc...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "hostname")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %v", err)
	}
	return string(output), nil
	//Code me Here
}

func WinCheck() {
	gettheBasics()
	checkWSL()
	roastAble()
	schTask()

}
