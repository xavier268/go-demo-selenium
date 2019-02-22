package main

import "github.com/tebeka/selenium"

// NewWebDriver creates the default WebDriver
func NewWebDriver() (wd selenium.WebDriver) {
	selenium.SetDebug(debug)
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"platform":    "linux"}
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	return wd
}
