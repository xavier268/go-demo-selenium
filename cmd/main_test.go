package main

import (
	"fmt"
	"testing"

	"github.com/tebeka/selenium"
)

func TestChromeStatus(t *testing.T) {
	selenium.SetDebug(false)
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"platform":    "LINUX"}
	wd, e1 := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	fatal(e1)
	defer wd.Quit()
	st, e2 := wd.Status()
	fatal(e2)
	fmt.Printf("\nStatus of webdriver : %v\n", st)
}

func TestFirefoxStatus(t *testing.T) {
	selenium.SetDebug(false)
	caps := selenium.Capabilities{
		"browserName": "firefox",
		"platform":    "LINUX"}
	wd, e1 := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	fatal(e1)
	defer wd.Quit()
	st, e2 := wd.Status()
	fatal(e2)
	fmt.Printf("\nStatus of webdriver : %v\n", st)
}

func fatal(e error) {
	if e != nil {
		panic(e)
	}
}
