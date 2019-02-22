package main

import (
	"fmt"
	"testing"

	"github.com/tebeka/selenium"
)

func TestChromeStatus(t *testing.T) {
	selenium.SetDebug(true)
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"platform":    "linux"}
	wd, e1 := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	fatal(e1)
	defer wd.Quit()
	st, e2 := wd.Status()
	fatal(e2)
	fmt.Printf("\nStatus of webdriver : %v\n", st)
}

func TestFirefoxStatus(t *testing.T) {

	//t.Log("The firefox browser generates a nil pointer panic in NewSession, I suspect firefox does not comply with new W3C capability format")
	//t.FailNow()

	selenium.SetDebug(true)
	caps := selenium.Capabilities{
		"browserName": "firefox",
		"platform":    "linux"}
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
