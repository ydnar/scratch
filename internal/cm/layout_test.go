package cm

import (
	"testing"
	"unsafe"
)

func TestLayoutAssumptions(t *testing.T) {
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

func TestVariantLayout(t *testing.T) {
	// 8 on 64-bit, 4 on 32-bit
	ptrSize := unsafe.Sizeof(uintptr(0))

	tests := []struct {
		v      VariantDebug
		size   uintptr
		offset uintptr
	}{
		{&UnsizedVariant2[struct{}, struct{}]{}, 1, 0},
		{&UnsizedVariant2[[0]byte, struct{}]{}, 1, 0},
		{&SizedVariant2[Shape[string], string, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedVariant2[Shape[string], bool, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedVariant2[Shape[string], string, struct{}]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedVariant2[Shape[string], struct{}, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedVariant2[Shape[uint64], uint64, uint64]{}, 16, alignOf[uint64]()},
		{&SizedVariant2[Shape[uint64], uint32, uint64]{}, 16, alignOf[uint64]()},
		{&SizedVariant2[Shape[uint64], uint64, uint32]{}, 16, alignOf[uint64]()},
		{&SizedVariant2[Shape[uint64], uint8, uint64]{}, 16, alignOf[uint64]()},
		{&SizedVariant2[Shape[uint64], uint64, uint8]{}, 16, alignOf[uint64]()},
		{&SizedVariant2[Shape[uint32], uint8, uint32]{}, 8, alignOf[uint32]()},
		{&SizedVariant2[Shape[uint32], uint32, uint8]{}, 8, alignOf[uint32]()},
		{&SizedVariant2[Shape[[9]byte], [9]byte, uint64]{}, 24, alignOf[uint64]()},
	}

	for _, tt := range tests {
		name := typeName(tt.v)
		t.Run(name, func(t *testing.T) {
			if got, want := tt.v.Size(), tt.size; got != want {
				t.Errorf("(%s).Size() == %v, expected %v", name, got, want)
			}
			if got, want := tt.v.ValOffset(), tt.offset; got != want {
				t.Errorf("(%s).ValOffset() == %v, expected %v", name, got, want)
			}
		})
	}
}

func TestResultLayout(t *testing.T) {
	// 8 on 64-bit, 4 on 32-bit
	ptrSize := unsafe.Sizeof(uintptr(0))

	tests := []struct {
		r      ResultDebug
		size   uintptr
		offset uintptr
	}{
		{&UnsizedResult[struct{}, struct{}]{}, 1, 0},
		{&UnsizedResult[[0]byte, struct{}]{}, 1, 0},

		{&SizedResult[Shape[string], string, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedResult[Shape[string], bool, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedResult[Shape[string], string, struct{}]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedResult[Shape[string], struct{}, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&SizedResult[Shape[uint64], uint64, uint64]{}, 16, alignOf[uint64]()},
		{&SizedResult[Shape[uint64], uint32, uint64]{}, 16, alignOf[uint64]()},
		{&SizedResult[Shape[uint64], uint64, uint32]{}, 16, alignOf[uint64]()},
		{&SizedResult[Shape[uint64], uint8, uint64]{}, 16, alignOf[uint64]()},
		{&SizedResult[Shape[uint64], uint64, uint8]{}, 16, alignOf[uint64]()},
		{&SizedResult[Shape[uint32], uint8, uint32]{}, 8, alignOf[uint32]()},
		{&SizedResult[Shape[uint32], uint32, uint8]{}, 8, alignOf[uint32]()},
		{&SizedResult[Shape[[9]byte], [9]byte, uint64]{}, 24, alignOf[uint64]()},

		{&OKSizedResult[string, struct{}]{}, sizePlusAlignOf[string](), ptrSize},
		{&OKSizedResult[string, bool]{}, sizePlusAlignOf[string](), ptrSize},
		{&OKSizedResult[[9]byte, uint64]{}, sizePlusAlignOf[string](), alignOf[uint64]()},

		{&ErrSizedResult[struct{}, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&ErrSizedResult[bool, string]{}, sizePlusAlignOf[string](), ptrSize},
		{&ErrSizedResult[uint64, [9]byte]{}, sizePlusAlignOf[string](), alignOf[uint64]()},
	}

	for _, tt := range tests {
		name := typeName(tt.r)
		t.Run(name, func(t *testing.T) {
			if got, want := tt.r.Size(), tt.size; got != want {
				t.Errorf("(%s).Size() == %v, expected %v", name, got, want)
			}
			if got, want := tt.r.ValOffset(), tt.offset; got != want {
				t.Errorf("(%s).ValOffset() == %v, expected %v", name, got, want)
			}
		})
	}
}

func sizePlusAlignOf[T any]() uintptr {
	var v T
	return unsafe.Sizeof(v) + unsafe.Alignof(v)
}

func alignOf[T any]() uintptr {
	var v T
	return unsafe.Alignof(v)
}
