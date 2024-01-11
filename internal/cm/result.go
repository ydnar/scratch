package cm

type Result[OK, Err any] interface {
	StoreOK(OK)
	StoreErr(Err)
	Result() (ok OK, err Err, isOK bool)
}

type OKSizedResult[OK any, Err any] struct {
	SizedResult[Shape[OK], OK, Err]
}

type ErrSizedResult[OK any, Err any] struct {
	SizedResult[Shape[Err], OK, Err]
}

// SizedResult is a tagged union that represents either the OK type or the Err type.
// Either OK or Err must have non-zero size, e.g. both cannot be struct{} or a zero-length array.
// For results with two zero-length types, use UnsizedResult.
type SizedResult[S Shape[OK] | Shape[Err], OK any, Err any] struct {
	v SizedVariant2[S, OK, Err]
}

// StoreOK stores the OK value in r.
func (r *SizedResult[S, OK, Err]) StoreOK(ok OK) {
	r.v.Store0(ok)
}

// StoreErr stores the error value in r.
func (r *SizedResult[S, OK, Err]) StoreErr(err Err) {
	r.v.Store1(err)
}

// Result returns the OK value and error value for r.
// If r represents an error, then the zero value of OK is returned.
// If r represents an OK value, then the zero value of Err is returned.
func (r *SizedResult[S, OK, Err]) Result() (ok OK, err Err, isOK bool) {
	ok, isOK = r.v.Load0()
	err, _ = r.v.Load1()
	return ok, err, isOK
}

type UnsizedResult[OK any, Err any] struct {
	v UnsizedVariant2[OK, Err]
}

// StoreErr stores the OK value in r.
func (r *UnsizedResult[OK, Err]) StoreOK(ok OK) {
	r.v.Store0(ok)
}

// StoreErr stores the error value in r.
func (r *UnsizedResult[OK, Err]) StoreErr(err Err) {
	r.v.Store1(err)
}

// Result returns the OK value and error value for r.
// If r represents an error, then the zero value of OK is returned.
// If r represents an OK value, then the zero value of Err is returned.
func (r *UnsizedResult[OK, Err]) Result() (ok OK, err Err, isOK bool) {
	ok, isOK = r.v.Load0()
	err, _ = r.v.Load1()
	return ok, err, isOK
}
