package main

import (
	"fmt"
	"handleJSON/streamEncoder"
	"handleJSON/tmpEncoder"
	"handleJSON/userDefinedJSONparser"
)

func main() {
	tmpEncoder.UsingTmpEncoder()
	fmt.Println("-----")
	streamEncoder.UsingStreamEncoder()
	fmt.Println("-----")
	userDefinedJSONparser.UsingUserDefinedJSONParser()
}
