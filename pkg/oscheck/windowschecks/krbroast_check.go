//go:build windows
// +build windows

package windowschecks

import (
	"fmt"
)

func RoastAble() (string, error) {
	fmt.Println("Can I kerberoast this box???")
	return "Kerberoasting check not implemented yet", nil
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
