package countIO

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func CountLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func UsingCountLetters() {
	s := "Go will controll the whole world"
	sr := strings.NewReader(s)
	counts, err := CountLetters(sr)
	if err != nil {
		log.Fatal("Failed")
	}
	fmt.Println(counts)
}
