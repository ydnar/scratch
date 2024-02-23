package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type S struct {
	_ [0]atomic.Uint64
	v uint32
}

func main() {
	var v S
	fmt.Printf("size: %d align: %d", unsafe.Sizeof(v), unsafe.Alignof(v))
	var i uintptr
	i = uintptr(unsafe.Alignof(v) << 29)
	_ = i
}
