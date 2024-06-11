package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	start := monoTime(runtimeInitTime)
	fmt.Printf("runtimeInitTime: %d \n", runtimeInitTime)
	fmt.Printf("timeStartNano:   %d\n", timeStartNano)
	fmt.Printf("since init:      %v\n", time.Since(start))
}

//go:linkname runtimeInitTime runtime.runtimeInitTime
var runtimeInitTime int64

//go:linkname timeStartNano time.startNano
var timeStartNano int64

// monoTime converts a monotonic nanotime to a [time.Time].
// The resulting Time has a zero-value wall clock,
// so is only useful for comparing against another Time value
// to create a Duration.
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
