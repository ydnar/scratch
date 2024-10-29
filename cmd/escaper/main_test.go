package main

import "testing"

var ok, err *int

func BenchmarkResultInlines(b *testing.B) {
	r := Result[int, int]{1, 2}
	for i := 0; i < b.N; i++ {
		ok, err = r.Result()
	}
}
