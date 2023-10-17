package main

import (
	"fmt"
	"os"
	_ "runtime"
	_ "unsafe"
)

func main() {
	fmt.Println("USER:", os.Getenv("USER"))
	fmt.Println("environ:", os.Environ())
}

//go:linkname getenv runtime.Getenv
func getenv(key string) (string, bool) {
	return "fake", true
}

//go:linkname environ runtime.Environ
func environ() []string {
	return []string{"USER=fake"}
}
