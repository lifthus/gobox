// This shows how you can use file system in go playground.
package main

import "play.ground/hello"

func main() {
	hello.SayHello()
}
-- go.mod --
module play.ground
-- hello/hello.go --
package hello

import "fmt"

func SayHello() {
	fmt.Println("Hello!")
}
