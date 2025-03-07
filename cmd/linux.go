package linux

import(
  "fmt"
  "os/exec"
  "log"
  "net"
  "os"
  "bufio"
  "path/filepath"
)

func CaptureSysInfo(){
  fmt.Println("Gathering basic Linux info")
}
