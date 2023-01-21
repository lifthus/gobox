package handlingBytes

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

type Data struct {
	Value  uint32   // 4 bytes ( big-endian integer )
	Label  [10]byte // 10 bytes ( ASCII name )
	Active bool     // 1 byte whether the item is active
	// and Go padded this with 1 byte to make it align
}

// get data without unsafe
func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4]) // first four bytes
	copy(d.Label[:], b[4:14])                // 10 bytes of label
	d.Active = b[14] != 0                    // whether active
	return d
}

// get data with usnafe
var isLE bool

func init() { // checking little endian
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}
func DataFromBytesUnsafe(b [16]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE { // if it is little endian,
		data.Value = bits.ReverseBytes32(data.Value) // reverse it
	}
	return data
}

// get bytes without unsafe
func BytesFromData(d Data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

// get bytes with unsafe
func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}
