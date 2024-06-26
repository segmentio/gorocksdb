package gorocksdb

import "C"
import (
	"unsafe"
)

// btoi converts a bool value to int.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// boolToChar converts a bool value to C.uchar.
func boolToChar(b bool) C.uchar {
	if b {
		return 1
	}
	return 0
}

// charToByte converts a *C.char to a byte slice.
func charToByte(data *C.char, len C.size_t) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(data)), int(len))
}

// byteToChar returns *C.char from byte slice.
func byteToChar(b []byte) *C.char {
	var c *C.char
	if len(b) > 0 {
		c = (*C.char)(unsafe.Pointer(&b[0]))
	}
	return c
}

// Go []byte to C string
// The C string is allocated in the C heap using malloc.
func cByteSlice(b []byte) *C.char {
	var c *C.char
	if len(b) > 0 {
		cData := C.malloc(C.size_t(len(b)))
		copy((*[1 << 24]byte)(cData)[0:len(b)], b)
		c = (*C.char)(cData)
	}
	return c
}

// charSlice converts a C array of *char to a []*C.char.
func charSlice(data **C.char, len C.int) []*C.char {
	return unsafe.Slice(data, len)
}

// sizeSlice converts a C array of size_t to a []C.size_t.
func sizeSlice(data *C.size_t, len C.int) []C.size_t {
	return unsafe.Slice(data, len)
}
