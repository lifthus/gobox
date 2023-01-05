// You can edit this code!
// Click here and start typing.
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
