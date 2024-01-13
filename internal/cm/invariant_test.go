package cm

import (
	"testing"
	"unsafe"
)

func TestFieldAlignment(t *testing.T) {
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

func TestBool(t *testing.T) {
	var b bool
	if got, want := unsafe.Sizeof(b), uintptr(1); got != want {
		t.Errorf("unsafe.Sizeof(b) == %d, expected %d", got, want)
	}

	// uint8(false) == 0
	b = false
	if got, want := *(*uint8)(unsafe.Pointer(&b)), uint8(0); got != want {
		t.Errorf("uint8(b) == %d, expected %d", got, want)
	}

	// uint8(true) == 1
	b = true
	if got, want := *(*uint8)(unsafe.Pointer(&b)), uint8(1); got != want {
		t.Errorf("uint8(b) == %d, expected %d", got, want)
	}

	// low bit 1 == true
	*(*uint8)(unsafe.Pointer(&b)) = 1
	if got, want := b, true; got != want {
		t.Errorf("b == %t, expected %t", got, want)
	}

	// low bit 0 == false
	*(*uint8)(unsafe.Pointer(&b)) = 2
	if got, want := b, false; got != want {
		t.Errorf("b == %t, expected %t", got, want)
	}

	// low bit 1 == true
	*(*uint8)(unsafe.Pointer(&b)) = 3
	if got, want := b, true; got != want {
		t.Errorf("b == %t, expected %t", got, want)
	}

	// low bit 0 == false
	*(*uint8)(unsafe.Pointer(&b)) = 254
	if got, want := b, false; got != want {
		t.Errorf("b == %t, expected %t", got, want)
	}

	// low bit 1 == true
	*(*uint8)(unsafe.Pointer(&b)) = 255
	if got, want := b, true; got != want {
		t.Errorf("b == %t, expected %t", got, want)
	}
}
