package main

import (
	"fmt"
	"goGeneric/interfaceTree"
)

func main() {
	var it *interfaceTree.Tree
	it = it.Insert(interfaceTree.OrderableInt(5))
	it = it.Insert(interfaceTree.OrderableInt(3))
	fmt.Println(it)
}
