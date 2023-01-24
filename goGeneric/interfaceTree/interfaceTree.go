package interfaceTree

import "strings"

type Orderable interface {
	// Order returns:
	// given value < 0, when the given value is bigger than Orderable.
	// given value > 0, when the given value is smaller than Orderable
	// 0 when boths are same.
	Order(interface{}) int
}

// Now, every type having Orderable which has Order method can be inserted to tree.

type Tree struct {
	val         Orderable
	left, right *Tree
}

func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil {
		return &Tree{val: val}
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

// OderableInt
type OrderableInt int

func (oi OrderableInt) Order(val interface{}) int {
	return int(oi - val.(OrderableInt))
}

// OrderableString
type OrderableString string

func (os OrderableString) Order(val interface{}) int {
	return strings.Compare(string(os), val.(string))
}
