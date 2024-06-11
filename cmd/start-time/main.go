package main

import (
	"fmt"
	"time"
	"unsafe"
	_ "unsafe"
)

func main() {
	start := monoTime(runtimeInitTime)
	// start := time.Unix(runtimeInitTime/1_000_000_000, runtimeInitTime%1_000_000_000)
	fmt.Printf("runtimeInitTime: %d \n", runtimeInitTime)
	fmt.Printf("timeStartNano:   %d\n", timeStartNano)
	fmt.Printf("since init:      %v\n", time.Since(start))
}

//go:linkname runtimeInitTime runtime.runtimeInitTime
var runtimeInitTime int64

//go:linkname timeStartNano time.startNano
var timeStartNano int64

func monoTime(mono int64) time.Time {
	var t time.Time
	p := (*struct {
		wall uint64
		ext  int64
		loc  *time.Location
	})(unsafe.Pointer(&t))
	p.wall |= 1 << 63
	p.ext = mono - timeStartNano // mirror mono -= startNano in time.Now()
	return t
}

func init() {
	for i := 0; i < 1_000_000; i++ {
		j += i
	}
}

var j int
