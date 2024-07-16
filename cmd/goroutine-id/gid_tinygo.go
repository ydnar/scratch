//go:build tinygo

package main

import (
	"unsafe"
)

var (
	goroutineID  int64
	goroutineIDs = make(map[uintptr]int64)
)

func GoroutineID() int64 {
	task := uintptr(currentTask())
	if id, ok := goroutineIDs[task]; ok {
		return id
	}
	goroutineID++
	goroutineIDs[task] = goroutineID
	return goroutineIDs[task]
}

//go:linkname currentTask internal/task.Current
func currentTask() unsafe.Pointer
