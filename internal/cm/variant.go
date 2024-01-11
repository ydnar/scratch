package cm

import "unsafe"

type Shape[T any] [1]T

type Variant2[T0, T1 any] interface {
	Disc() uint
	Store0(T0)
	Store1(T1)
	Load0() (T0, bool)
	Load1() (T1, bool)
}

// SizedVariant2 represents a variant with 2 associated types, at least one of which has a non-zero size.
// Use UnsizedVariant2 if both T0 or T1 are zero-sized.
// The memory layout will have additional padding if both T0 and T1 are zero-sized.
type SizedVariant2[S Shape[T0] | Shape[T1], T0, T1 any] struct {
	disc uint8
	_    [0]T0
	_    [0]T1
	val  S
}

func (v *SizedVariant2[S, T0, T1]) Disc() uint {
	return uint(v.disc)
}

func (v *SizedVariant2[S, T0, T1]) Store0(val T0) {
	store(&v.disc, 0, &v.val, val)
}

func (v *SizedVariant2[S, T0, T1]) Store1(val T1) {
	store(&v.disc, 1, &v.val, val)
}

func (v *SizedVariant2[S, T0, T1]) Load0() (val T0, ok bool) {
	return load[T0](&v.disc, 0, &v.val)
}

func (v *SizedVariant2[S, T0, T1]) Load1() (val T1, ok bool) {
	return load[T1](&v.disc, 0, &v.val)
}

// UnsizedVariant2 represents a variant with 2 zero-sized associated types, e.g. struct{} or [0]T.
// Use SizedVariant2 if either T0 or T1 has a non-zero size.
// Loads and stores may panic if T0 or T1 has a non-zero size.
type UnsizedVariant2[T0, T1 any] struct {
	val  struct{} // first to avoid padding of zero-sized trailing field
	disc uint8
}

func (v *UnsizedVariant2[T0, T1]) Disc() uint {
	return uint(v.disc)
}

func (v *UnsizedVariant2[T0, T1]) Store0(val T0) {
	store(&v.disc, 0, &v.val, val)
}

func (v *UnsizedVariant2[T0, T1]) Store1(val T1) {
	store(&v.disc, 1, &v.val, val)
}

func (v *UnsizedVariant2[T0, T1]) Load0() (val T0, ok bool) {
	return load[T0](&v.disc, 0, &v.val)
}

func (v *UnsizedVariant2[T0, T1]) Load1() (val T1, ok bool) {
	return load[T1](&v.disc, 0, &v.val)
}

func store[T any, S any, Disc uint8 | uint16 | uint32](disc *Disc, n Disc, ptr *S, val T) {
	*(*T)(unsafe.Pointer(ptr)) = val
	*disc = n
}

func load[T any, S any, Disc uint8 | uint16 | uint32](disc *Disc, n Disc, ptr *S) (val T, ok bool) {
	if *disc != n {
		return val, false
	}
	return *(*T)(unsafe.Pointer(ptr)), true
}
