package cm

type Result[OK, Err any] interface {
	IsOK() bool
	IsErr() bool // TODO: remove this method everywhere
	StoreOK(OK)
	StoreErr(Err)
	LoadOK() (OK, bool)
	LoadErr() (Err, bool)
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

// IsOK returns true if r represents the OK value.
func (r *SizedResult[S, OK, Err]) IsOK() bool {
	return r.v.disc == 0
}

// IsErr returns true if r represents the error value.
func (r *SizedResult[S, OK, Err]) IsErr() bool {
	return r.v.disc == 1
}

// StoreOK stores the OK value in r.
func (r *SizedResult[S, OK, Err]) StoreOK(ok OK) {
	r.v.Store0(ok)
}

// StoreErr stores the error value in r.
func (r *SizedResult[S, OK, Err]) StoreErr(err Err) {
	r.v.Store1(err)
}

// LoadOK returns the OK value of r.
// If r is an error value, then the zero value of OK is returned.
func (r *SizedResult[S, OK, Err]) LoadOK() (ok OK, isOK bool) {
	return r.v.Load0()
}

// LoadErr returns the Err value of r.
// If r is an OK value, then the zero value of Err is returned.
func (r *SizedResult[S, OK, Err]) LoadErr() (err Err, isErr bool) {
	return r.v.Load1()
}

// Result returns the OK value and error value for r.
// If r represents an error, then the zero value of OK is returned.
// If r represents an OK value, then the zero value of Err is returned.
func (r *SizedResult[S, OK, Err]) Result() (ok OK, err Err, isOK bool) {
	ok, isOK = r.LoadOK()
	err, _ = r.LoadErr()
	return ok, err, isOK
}

type UnsizedResult[OK any, Err any] struct {
	v UnsizedVariant2[OK, Err]
}

// IsOK returns true if r represents the OK value.
func (r UnsizedResult[OK, Err]) IsOK() bool {
	return r.v.disc == 0
}

// IsErr returns true if r represents the error value.
func (r UnsizedResult[OK, Err]) IsErr() bool {
	return r.v.disc == 1
}

// StoreErr stores the OK value in r.
func (r *UnsizedResult[OK, Err]) StoreOK(ok OK) {
	r.v.Store0(ok)
}

// StoreErr stores the error value in r.
func (r *UnsizedResult[OK, Err]) StoreErr(err Err) {
	r.v.Store1(err)
}

// LoadOK returns the OK value of r.
func (r *UnsizedResult[OK, Err]) LoadOK() (ok OK, isOK bool) {
	return r.v.Load0()
}

// LoadErr returns the Err value of r.
func (r *UnsizedResult[OK, Err]) LoadErr() (err Err, isErr bool) {
	return r.v.Load1()
}
