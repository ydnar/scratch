package main

import (
	"time"
	"unsafe"
)

func main() {
	var t uint64
	errno := clock_time_get(0, 1e3, unsafe.Pointer(&t))
	if errno != 0 {
		panic(errno)
	}
	sec := int64(t / 1e9)
	nsec := int64(t % 1e9)
	println("clock_time_get:", time.Unix(sec, nsec).String())
	println("time.Now():", time.Now().String())
}

//go:wasmimport wasi_snapshot_preview1 clock_time_get
func clock_time_get(clockid uint32, precision uint64, time unsafe.Pointer) (errno uint32)
