package genericTree

import "strings"

type Orderable[T any] interface {
	Order(T) int
}

type OrderableInt int

func (oi OrderableInt) Order(val OrderableInt) int {
	return int(oi - val)
}

type OrderableString string

func (os OrderableString) Order(val OrderableString) int {
	return strings.Compare(string(os), string(val))
}

// Tree having Orderable val
type Tree[T Orderable[T]] struct {
	val         T
	left, right *Tree[T]
}

func (t *Tree[T]) Insert(val T) *Tree[T] {
	if t == nil {
		return &Tree[T]{val: val}
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

func (t *Tree[T]) Contains(val T) bool {
	if t.val.Order(val) == 0 {
		return true
	}
	if t.left != nil {
		if t.left.Contains(val) {
			return true
		}
	}
	if t.right != nil {
		if t.right.Contains(val) {
			return true
		}
	}
	return false
}
