package cm

import "unsafe"

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
