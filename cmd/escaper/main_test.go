package main

import "testing"

func BenchmarkResultInlines(b *testing.B) {
	r := Result[int, int]{1, 2}
	var ok, err *int
	for i := 0; i < b.N; i++ {
		ok, err = r.Result()
	}
	_, _ = ok, err
}
