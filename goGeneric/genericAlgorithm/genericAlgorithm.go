package genericAlgorithm

import (
	"fmt"
)

// Map converts []T1 to []T2 using mapping function.
// this function has two type parameters, T1 and T2.
// it works with slices of every type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces []T1 to single value using reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters a slice using filter function.
// it returns a new slice that is consisted of the elements for whcih f returns true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func UsingAbove() {
	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)
	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)
	sum := Reduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum)
}
