package alignment

import (
	"testing"
	"unsafe"
)

func TestAlignment(t *testing.T) {
	var r1 Result[uintptr, bool]
	t.Log("sizeof r1", unsafe.Sizeof(r1))
	t.Log("offset r1.a", unsafe.Offsetof(r1.ok))
	t.Log("offset r1.b", unsafe.Offsetof(r1.err))

	var r2 Result[bool, uint8]
	t.Log("sizeof r2", unsafe.Sizeof(r2))
	t.Log("offset r2.a", unsafe.Offsetof(r2.ok))
	t.Log("offset r2.b", unsafe.Offsetof(r2.err))
}

type Result[OK any, Err any] struct {
	isErr bool
	ok    [0]OK
	err   [1]Err
}

func (r *Result[OK, Err]) Result() (OK, Err) {
	return r.OK(), r.Err()
}

func (r *Result[OK, Err]) OK() OK {
	if r.isErr {
		var zero OK
		return zero
	}
	return *(*OK)(unsafe.Pointer(&r.ok))
}

func (r *Result[OK, Err]) Err() Err {
	if !r.isErr {
		var zero Err
		return zero
	}
	return r.err[0]
}
