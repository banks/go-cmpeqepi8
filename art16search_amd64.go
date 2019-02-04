//+build !noasm
//+build !appengine

//go:generate build-asm.sh

package cmpeqepi8

import (
	"unsafe"
)

//go:noescape
func _IndexOfByteIn16Bytes(n, arr, c, idx unsafe.Pointer)

func IndexOfByteIn16Bytes(n uint8, arr *[16]byte, c byte) int {
	idx := int8(-1)
	_IndexOfByteIn16Bytes(unsafe.Pointer(&n), unsafe.Pointer(arr), unsafe.Pointer(&c), unsafe.Pointer(&idx))
	return int(idx)
}
