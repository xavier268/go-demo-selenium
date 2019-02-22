package main

import (
	"fmt"

	"github.com/tebeka/selenium"
)

var debug = true

func main() {
	selenium.SetDebug(debug)
	fmt.Println("Demo of using the GO selenium client")
}
