package main

import (
	"fmt"
	"sharedRW/chanScoreboard"
	"sharedRW/mutexScoreboard"
)

func main() {
	sbChan, done := chanScoreboard.NewChannelScoreboardManger()
	sbChan.Update("a", 5)
	fmt.Println(sbChan.Read("a"))
	fmt.Println(sbChan.Read("b"))
	done()

	sbMutex := mutexScoreboard.NewMutexScoreboardManager()
	sbMutex.Update("a", 10)
	fmt.Println(sbMutex.Read("a"))
	fmt.Println(sbMutex.Read("b"))
	sbMutex.Update("c", 159)

	// !!CAUTION!! still can access to the data without token.
	fmt.Println(sbMutex.ReadDangerously("c"))
}
