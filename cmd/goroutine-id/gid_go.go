//go:build !tinygo

package main

import "unsafe"

func GoroutineID() uintptr {
	return 0
}

//go:linkname currentTask internal/task.Current
func currentTask() unsafe.Pointer
