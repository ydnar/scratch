package main

func main() {
	r := Result[int, int]{1, 2}
	ok, err := r.Result()
	_, _ = ok, err
}

type Result[OK, Err any] struct {
	ok  OK
	err Err
}

func (r Result[OK, Err]) Result() (*OK, *Err) {
	return &r.ok, &r.err
}
