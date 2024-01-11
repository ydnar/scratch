package cm

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
