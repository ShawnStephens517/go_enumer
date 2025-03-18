//go:build windows
// +build windows

package windowschecks

import (
	"fmt"
)

// Check for credentials in files
func Possible_Pass() (string, error) {
	//Loop through drives looking for files with names like User or Password
	//Also, if xls, xlsm, xlsx; look for user or pass in content
	return "Feature not Implemented", nil

}

// Check for Mcaffee SiteList.xml
func Mcaffee_Check() (string, error) {
	//Look for Mcaffee site list. .*SiteList.xml
	fmt.Println("Site List may be decrypted using \nhttps://github.com/funoverip/mcafee-sitelist-pwd-decryption")
	return "Feature not Implemented", nil
}
