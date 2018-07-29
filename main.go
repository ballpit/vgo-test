package main

import (
	"fmt"

	dep1 "github.com/ballpit/vgo-dep1"
	dep2 "github.com/ballpit/vgo-dep2"
)

func main() {
	fmt.Printf(dep1.TestFunc() + "\n")
	fmt.Printf(dep2.RootFunc() + "\n")
	fmt.Printf("Hello, world.\n")
}
