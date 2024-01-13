package cm

type Option[T any] struct {
	disc uint8
	v    T
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		disc: 1,
		v:    v,
	}
}

func None[T any]() Option[T] {
	return Option[T]{disc: 0}
}

func (o Option[T]) None() bool {
	return o.disc == 0
}

func (o Option[T]) Some() (T, bool) {
	return o.v, o.disc == 1
}
