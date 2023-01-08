package main

import (
	"fmt"
	sharedRW "sharedRW/chanScoreboard"
)

func main() {
	sbSharedRW, done := sharedRW.NewChannelScoreboardManger()
	sbSharedRW.Update("a", 5)
	fmt.Println(sbSharedRW.Read("a"))
	fmt.Println(sbSharedRW.Read("b"))
	done()
}
