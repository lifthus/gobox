package main

import (
	"fmt"
	"goIO/countIO"
	"goIO/gzipRead"
)

func main() {
	fmt.Println("CountLetters:")
	countIO.UsingCountLetters()
	fmt.Println("GzipReader")
	gzipRead.UsingBuildGZipReader()
}
