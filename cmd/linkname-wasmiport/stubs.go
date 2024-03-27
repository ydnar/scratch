// //go:build !wasip1

package main

import "unsafe"

type (
	errno     = uint32
	clockid   = uint32
	timestamp = uint64
)

//go:linkname clock_time_get runtime.clock_time_get
func clock_time_get(clock_id clockid, precision timestamp, time unsafe.Pointer) errno {
	return 0
}
