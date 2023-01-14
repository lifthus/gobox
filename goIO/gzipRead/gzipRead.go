package gzipRead

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"

	"goIO/countIO"
)

func BuildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func UsingBuildGZipReader() {
	r, closer, err := BuildGZipReader("/Users/ijeonghun/gitRepos/gobox/goIO/gzipRead/my_data.txt.gz")
	if err != nil {
		log.Fatal("Building reader failed")
	}
	defer closer()
	counts, err := countIO.CountLetters(r)
	if err != nil {
		log.Fatal("Counting failed")
	}
	fmt.Println(counts)
}
