package cm

import "unsafe"

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