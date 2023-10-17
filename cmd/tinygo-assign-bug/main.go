package main

/*
This program demonstrates a difference in behavior between Go and TinyGo.
In Go, the multiple assignment is evaluated after the call to toUpper.
In TinyGo, the assignment happens prior to the call to toUpper.

See: https://go.dev/ref/spec#Assignment_statements

$ go run ./cmd/tinygo-assign-bug
HELLO HELLO <nil>

$ tinygo run ./cmd/tinygo-assign-bug
HELLO hello <nil>
*/

import (
	"strings"
)

func main() {
	s1 := "hello"
	s2, err := s1, toUpper(&s1)
	println(s1, s2, err)
}

func toUpper(s *string) error {
	*s = strings.ToUpper(*s)
	return nil
}
