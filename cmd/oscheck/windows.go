package oscheck

import (
  "fmt"
  "os/exec"
  "log"
  "net"
  "os"
  "bufio"
  "path/filepath"
  )

func checkWSL(){
  fmt.Println("Capture if WSL Enabled???")
  //Code Me Here
}

func roastMe(){
  fmt.Println("Can I kerberoast this box???")
  //Code Me Here
}

func schTask(){
  fmt.Println("What Scheduled Tasks can I see???")
  //Code Me Here
}

func gettheBasics(){
  fmt.Println("What is the Hostname, Scriptblock Logging, etc...")
  //Code me Here
}

func WinCheck(){
  checkWSL()
  roastMe()
  schTask()

}