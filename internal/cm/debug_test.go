package cm

import (
	"reflect"
	"strings"
	"unsafe"
)

func typeName(v any) string {
	var name string
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		name = "*" + t.Elem().String()
	} else {
		name = t.String()
	}
	return strings.ReplaceAll(name, " ", "")
}

// VariantDebug is an interface used in tests to validate layout of variant types.
type VariantDebug interface {
	Size() uintptr
	ValAlign() uintptr
	ValOffset() uintptr
}

func (v *SizedVariant2[Shape, T0, T1]) Size() uintptr {
	return unsafe.Sizeof(*v)
}

func (v *SizedVariant2[Shape, T0, T1]) ValAlign() uintptr {
	return unsafe.Alignof(v.val)
}

func (v *SizedVariant2[Shape, T0, T1]) ValOffset() uintptr {
	return unsafe.Offsetof(v.val)
}

func (v *UnsizedVariant2[T0, T1]) Size() uintptr {
	return unsafe.Sizeof(*v)
}

func (v *UnsizedVariant2[T0, T1]) ValAlign() uintptr {
	return unsafe.Alignof(v.val)
}

func (v *UnsizedVariant2[T0, T1]) ValOffset() uintptr {
	return unsafe.Offsetof(v.val)
}

// ResultDebug is an interface used in tests to validate layout of result types.
type ResultDebug interface {
	VariantDebug
}

func (r *SizedResult[S, OK, Err]) Size() uintptr {
	return unsafe.Sizeof(*r)
}

func (r *SizedResult[S, OK, Err]) ValAlign() uintptr {
	return r.v.ValAlign()
}

func (r *SizedResult[S, OK, Err]) ValOffset() uintptr {
	return r.v.ValOffset()
}

func (r *UnsizedResult[OK, Err]) Size() uintptr {
	return unsafe.Sizeof(*r)
}

func (r *UnsizedResult[OK, Err]) ValAlign() uintptr {
	return r.v.ValAlign()
}

func (r *UnsizedResult[OK, Err]) ValOffset() uintptr {
	return r.v.ValOffset()
}
