package main

import (
	"fmt"
	"goGeneric/interfaceTree"
	sG "goGeneric/stackGeneric"
	sNG "goGeneric/stackNonGeneric"
)

func main() {
	var it *interfaceTree.Tree
	it = it.Insert(interfaceTree.OrderableInt(5))
	it = it.Insert(interfaceTree.OrderableInt(3))
	fmt.Println(it)

	var s sNG.Stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	v, ok := s.Pop()
	fmt.Println(v, ok)

	var s sG.Stack[int]
	s.Push(10) // compile error when trying pushing non-int type data.
	s.Push(20)
	s.Push(30)
	v, ok = s.Pop()
	fmt.Println(v, ok)
}
