package main

import (
	"fmt"

	"github.com/ballpit/depthree/v2"
	"github.com/ballpit/vgo-dep1"
	"github.com/ballpit/vgo-dep2"
)

func main() {
	fmt.Printf(dep1.TestFunc() + "\n")
	fmt.Printf(depthree.Dep3func() + "\n")
	fmt.Printf(dep2.RootFunc() + "\n")
	fmt.Printf("Hello, world.\n")
}
