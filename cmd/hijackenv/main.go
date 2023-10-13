package main

import (
	"fmt"
	"os"
	_ "unsafe"
)

func main() {
	fmt.Println("USER:", os.Getenv("USER"))
	fmt.Println("environ:", os.Environ())
}

//go:linkname getenv os.Getenv
func getenv(key string) string {
	return "fake"
}

//go:linkname environ syscall.Environ
func environ() []string {
	return []string{"USER=fake"}
}
