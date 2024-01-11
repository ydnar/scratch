package layout

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestAssumptions(t *testing.T) {
	var v1 struct {
		_   bool
		_   [0][7]byte
		u64 uint64
	}
	if want, got := uintptr(16), unsafe.Sizeof(v1); want != got {
		t.Errorf("expected unsafe.Sizeof(v1) == %d, got %d", want, got)
	}
	if want, got := uintptr(8), unsafe.Offsetof(v1.u64); want != got {
		t.Errorf("expected unsafe.Offsetof(v1.u64) == %d, got %d", want, got)
	}

	var v2 struct {
		_ bool
		_ [0][7]byte
		_ [0][51]float64
		_ [0]struct {
			uint64
			_ []byte
		}
		u64 uint64
	}
	if want, got := uintptr(16), unsafe.Sizeof(v2); want != got {
		t.Errorf("expected unsafe.Sizeof(v2) == %d, got %d", want, got)
	}
	if want, got := uintptr(8), unsafe.Offsetof(v2.u64); want != got {
		t.Errorf("expected unsafe.Offsetof(v2.u64) == %d, got %d", want, got)
	}

	// size 1
	var v3 struct {
		_ struct{}
		b bool // offset 0
	}
	if want, got := uintptr(1), unsafe.Sizeof(v3); want != got {
		t.Errorf("expected unsafe.Sizeof(v3) == %d, got %d", want, got)
	}
	if want, got := uintptr(0), unsafe.Offsetof(v3.b); want != got {
		t.Errorf("expected unsafe.Offsetof(v3.b) == %d, got %d", want, got)
	}

	// size 0
	var v4 struct {
		_ [0]uint32
		b bool // offset 0!
	}
	if want, got := uintptr(4), unsafe.Sizeof(v4); want != got {
		t.Errorf("expected unsafe.Sizeof(v4) == %d, got %d", want, got)
	}
	if want, got := uintptr(0), unsafe.Offsetof(v4.b); want != got {
		t.Errorf("expected unsafe.Offsetof(v4.b) == %d, got %d", want, got)
	}
}

func TestResultLayout(t *testing.T) {
	var r UnsizedResult[struct{}, struct{}]
	if want, got := uintptr(1), unsafe.Sizeof(r); want != got {
		t.Errorf("expected unsafe.Sizeof(UntypedResult) == %d, got %d", want, got)
	}

	// testResultLayout[Shape[struct{}], struct{}, struct{}](t, 1, 1)
	testResultLayout[Shape[uintptr], uintptr, bool](t, 16, 8)
	testResultLayout[Shape[uintptr], bool, uintptr](t, 16, 8)
	testResultLayout[Shape[string], struct{}, string](t, 24, 8)
	testResultLayout[Shape[string], string, struct{}](t, 24, 8)
	testResultLayout[Shape[string], string, string](t, 24, 8)

	// Alignment and size
	testResultLayout[Shape[[7]byte], [7]byte, uint64](t, 16, 8)
}

func testResultLayout[S Shape[OK] | Shape[Err], OK any, Err any](t *testing.T, size, offset uintptr) {
	var shape S
	var ok OK
	var err Err
	types := strings.ReplaceAll(typeName(shape)+", "+typeName(ok)+", "+typeName(err), "layout.", "")
	var r SizedResult[S, OK, Err]
	if want, got := size, unsafe.Sizeof(r); want != got {
		t.Errorf("expected unsafe.Sizeof(SizedResult[%s]) == %d, got %d", types, want, got)
	}
	if want, got := offset, unsafe.Offsetof(r.v); want != got {
		t.Errorf("expected unsafe.Offsetof(SizedResult[%s].v) == %d, got %d", types, want, got)
	}
}

func typeName(v any) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().String()
	} else {
		return t.String()
	}
}

type Shape[T any] [1]T

type UnsizedResult[OK any, Err any] uint8

func (r UnsizedResult[OK, Err]) IsOK() bool {
	return r == 0
}

func (r UnsizedResult[OK, Err]) IsErr() bool {
	return r == 1
}

type OKResult[OK any, Err any] struct {
	SizedResult[Shape[OK], OK, Err]
}

type ErrResult[OK any, Err any] struct {
	SizedResult[Shape[Err], OK, Err]
}

// SizedResult is a tagged union that represents either the OK type or the Err type.
// Either OK or Err must have non-zero size, e.g. both cannot be struct{} or a zero-length array.
// For results with two zero-length types, use UnsizedResult.
type SizedResult[S Shape[OK] | Shape[Err], OK any, Err any] struct {
	disc uint8
	v    S
}

func (r *SizedResult[S, OK, Err]) IsOK() bool {
	return r.disc == 0
}

func (r *SizedResult[S, OK, Err]) IsErr() bool {
	return r.disc == 1
}

func (r *SizedResult[S, OK, Err]) SetOK(ok OK) {
	r.disc = 0
	*(*OK)(unsafe.Pointer(&r.v)) = ok
}

func (r *SizedResult[S, OK, Err]) SetErr(err Err) {
	r.disc = 1
	*(*Err)(unsafe.Pointer(&r.v)) = err
}

// Result returns the OK value and error value for r.
// If r represents an error, then the zero value of OK is returned.
// If r represents an OK value, then the zero value of Err is returned.
func (r *SizedResult[S, OK, Err]) Result() (ok OK, err Err, isOK bool) {
	ok, isOK = r.OK()
	err, _ = r.Err()
	return ok, err, isOK
}

// OK returns the OK value of r.
// If r is an error value, then the zero value of OK is returned.
func (r *SizedResult[S, OK, Err]) OK() (ok OK, isOK bool) {
	if !r.IsOK() {
		return ok, false
	}
	return *(*OK)(unsafe.Pointer(&r.v)), true
}

// OK returns the Err value of r.
// If r is an OK value, then the zero value of Err is returned.
func (r *SizedResult[S, OK, Err]) Err() (err Err, isErr bool) {
	if !r.IsErr() {
		return err, false
	}
	return *(*Err)(unsafe.Pointer(&r.v)), true
}

// SizedVariant2 represents a variant with 2 associated types, at least one of which has a non-zero size.
// Use UnsizedVariant2 if both T0 or T1 are zero-sized.
// The memory layout will have additional padding if both T0 and T1 are zero-sized.
type SizedVariant2[Shape, T0, T1 any] struct {
	disc uint8
	_    [0]T0
	_    [0]T1
	val  Shape
}

func (v *SizedVariant2[Shape, T0, T1]) Store0(val T0) {
	store(&v.disc, 0, &v.val, val)
}

func (v *SizedVariant2[Shape, T0, T1]) Store1(val T1) {
	store(&v.disc, 1, &v.val, val)
}

func (v *SizedVariant2[Shape, T0, T1]) Load0() (val T0, ok bool) {
	return load[T0](&v.disc, 0, &v.val)
}

func (v *SizedVariant2[Shape, T0, T1]) Load1() (val T1, ok bool) {
	return load[T1](&v.disc, 0, &v.val)
}

// UnsizedVariant2 represents a variant with 2 zero-sized associated types, e.g. struct{} or [0]T.
// Use SizedVariant2 if either T0 or T1 has a non-zero size.
// Loads and stores may panic if T0 or T1 has a non-zero size.
type UnsizedVariant2[Shape, T0, T1 any] struct {
	val  struct{} // first to avoid padding of zero-sized trailing field
	disc uint8
}

func (v *UnsizedVariant2[Shape, T0, T1]) Store0(val T0) {
	store(&v.disc, 0, &v.val, val)
}

func (v *UnsizedVariant2[Shape, T0, T1]) Store1(val T1) {
	store(&v.disc, 1, &v.val, val)
}

func (v *UnsizedVariant2[Shape, T0, T1]) Load0() (val T0, ok bool) {
	return load[T0](&v.disc, 0, &v.val)
}

func (v *UnsizedVariant2[Shape, T0, T1]) Load1() (val T1, ok bool) {
	return load[T1](&v.disc, 0, &v.val)
}

func store[T any, Shape any, Disc uint8 | uint16 | uint32](disc *Disc, n Disc, ptr *Shape, val T) {
	*(*T)(unsafe.Pointer(ptr)) = val
	*disc = n
}

func load[T any, Shape any, Disc uint8 | uint16 | uint32](disc *Disc, n Disc, ptr *Shape) (val T, ok bool) {
	if *disc != n {
		return val, false
	}
	return *(*T)(unsafe.Pointer(ptr)), true
}
