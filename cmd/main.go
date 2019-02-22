package main

import (
  "fmt"
 "time"
 "github.com/tebeka/selenium"
)

var debug = true

func main() {
  selenium.SetDebug(debug)
	fmt.Println("Demo of using the GO selenium client")
  printStatus()

}

func printStatus() {
  caps := selenium.Capabilities{
    "browserName":"firefox",
    "platform":"LINUX"  }
  wd, e1 := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
  fatal(e1)
  defer wd.Quit()
  st, e2 := wd.Status()
  fatal(e2)
  fmt.Printf("\nStatus of webdriver : %v\n",st)
  time.Sleep(time.Second)
}


func fatal(e error) {
  if(e != nil) {
    panic(e)
  }
}
