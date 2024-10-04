//go:build never

package collide

import "unsafe"

func f() {
	_ = unsafe.Pointer(nil)
}
