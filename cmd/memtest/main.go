package main

import (
	"math/rand"
	"runtime"
	"unsafe"
)

func main() {
	t1 := g()[0]
	t2 := g()[0]
	t3 := f()
	trash := garbage(1000)
	runtime.GC()
	println(t1.P)
	println(t2.P)
	println(t3.P)
	runtime.KeepAlive(trash)
}

type T struct {
	P *string
}

func f() T {
	var foo int
	_ = foo
	v := T{}
	v.P = new(string)
	*(*uintptr)(unsafe.Pointer(&v.P)) = uintptr(unsafe.Add(unsafe.Pointer(&foo), 1))
	return v
}

func g() []T {
	var t [100000]int // a large stack frame to trigger stack growing
	_ = t
	var out []T
	for i := 0; i < 1000; i++ {
		out = append(out, f())
	}
	return out
}

func garbage(n int) []unsafe.Pointer {
	out := make([]unsafe.Pointer, n)
	for i := 0; i < n; i++ {
		*(*int64)(unsafe.Pointer(&out[i])) = rand.Int63()
	}
	return out
}
