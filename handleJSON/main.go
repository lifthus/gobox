package main

import (
	"fmt"
	"handleJSON/streamEncoder"
	"handleJSON/tmpEncoder"
)

func main() {
	tmpEncoder.UsingTmpEncoder()
	fmt.Println("------")
	streamEncoder.UsingStreamEncoder()
}
