//go:build tinygo

package main

import "unsafe"

func GoroutineID() uintptr {
	return uintptr(currentTask())
}

//go:linkname currentTask internal/task.Current
func currentTask() unsafe.Pointer
