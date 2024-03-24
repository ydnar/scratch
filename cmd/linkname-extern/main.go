package main

import (
	_ "unsafe"

	"github.com/ydnar/scratch/cmd/linkname-extern/hello"
)

func main() {
	println(hello.World())
	println(hello.String("hello string").String())
}

//go:linkname World github.com/ydnar/scratch/cmd/linkname-extern/hello.World
func World() string {
	return "hello world"
}

//go:linkname StringString github.com/ydnar/scratch/cmd/linkname-extern/hello.String.String
func StringString(s hello.String) string {
	return string(s)
}
