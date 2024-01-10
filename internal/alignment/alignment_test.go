package alignment

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestResultAlignment(t *testing.T) {
	testResultAlignment[struct{}, struct{}](t, 2, 1)
	testResultAlignment[uintptr, bool](t, 16, 8)
	testResultAlignment[bool, uintptr](t, 16, 8)
	testResultAlignment[struct{}, string](t, 24, 8)
	testResultAlignment[string, struct{}](t, 24, 8)
	testResultAlignment[string, string](t, 24, 8)
}

func testResultAlignment[OK any, Err any](t *testing.T, size, offset uintptr) {
	var ok OK
	var err Err
	types := typeName(ok) + ", " + typeName(err)
	var r Result[OK, Err]
	if got, want := unsafe.Sizeof(r), size; got != want {
		t.Errorf("expected unsafe.Sizeof(Result[%s]) == %d, want %d", types, got, want)
	}
	if got, want := unsafe.Offsetof(r.v), offset; got != want {
		t.Errorf("expected unsafe.Offsetof(Result[%s].v) == %d, want %d", types, got, want)
	}
}

func typeName(v any) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().String()
	} else {
		return t.String()
	}
}

// Result is a union that represents either the OK type or the Err type.
type Result[OK any, Err any] struct {
	isErr bool
	_     [0]OK
	v     [1]Err
}

func (r *Result[_, _]) IsOK() bool {
	return !r.isErr
}

func (r *Result[_, _]) IsErr() bool {
	return r.isErr
}

/*
func (r *Result[OK, Err]) SetOK(ok OK) {
	r.isErr = false
	*(*OK)(unsafe.Pointer(&r.ok)) = ok
}

func (r *Result[OK, Err]) SetErr(err Err) {
	r.isErr = true
	r.err[0] = err
}

// Result returns the OK value and error value for r.
// If r represents an error, then the zero value of OK is returned.
// If r represents an OK value, then the zero value of Err is returned.
func (r *Result[OK, Err]) Result() (OK, Err) {
	return r.OK(), r.Err()
}

// OK returns the OK value of r.
// If r is an error value, then the zero value of OK is returned.
func (r *Result[OK, Err]) OK() OK {
	if r.isErr {
		var zero OK
		return zero
	}
	return *(*OK)(unsafe.Pointer(&r.ok))
}

// OK returns the Err value of r.
// If r is an OK value, then the zero value of Err is returned.
func (r *Result[OK, Err]) Err() Err {
	if !r.isErr {
		var zero Err
		return zero
	}
	return r.err[0]
}
*/
