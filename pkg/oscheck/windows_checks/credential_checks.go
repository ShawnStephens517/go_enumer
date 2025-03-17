// +build windows

package windows_checks

import (
	"context"
	"fmt"
  "os"
	"os/exec"
	"time"
  "path/filepath"

  "golang.org/x/sys/windows/registry"
)

func Possible_Pass() (string, error){
  //Loop through drives looking for files with names like User or Password
  //Also, if xls, xlsm, xlsx; look for user or pass in content

}

func Mcaffee_Check() (string, error){
  //Look for Mcaffee site list. .*SiteList.xml
  fmt.Println("Site List may be decrypted using \nhttps://github.com/funoverip/mcafee-sitelist-pwd-decryption")
}
