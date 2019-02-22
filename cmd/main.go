package main

import (
	"fmt"

	"github.com/tebeka/selenium"
)

var debug = false

func main() {
	selenium.SetDebug(debug)
	fmt.Println("Demo of using the GO selenium client")

	wd := NewWebDriver()
	defer wd.Quit()

	wd.Get("http://www.google.fr")
	t, err := wd.Title()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nPage title from google : %v\n", t)

	we, e2 := wd.FindElement(selenium.ByTagName, "html")
	if e2 != nil {
		panic(e2)
	}
	c, _ := we.Text()
	fmt.Printf("\nPage content : %v\n", c)
}
