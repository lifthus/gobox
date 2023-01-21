package main

type Data struct {
	Value  uint32   // 4 bytes ( big-endian integer )
	Label  [10]byte // 10 bytes ( ASCII name )
	Active bool     // 1byte whether the item is active
	// and Go padded this with 1 byte to make it align
}

func main() {
}
