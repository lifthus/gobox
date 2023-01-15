package benchmark

import (
	"fmt"
	"testing"
)

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 65204 {
		t.Error("Expectd 65204, got", result)
	}
}

var blackhole int

func BenchmarkFileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
