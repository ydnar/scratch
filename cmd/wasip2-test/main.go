package main

import (
	"fmt"
	"os"
)

// This requires TinyGo with WASI Preview 2 support
// https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open(%q): %v", filename, err)
		return
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "File.Stat(): %v", filename, err)
		return
	}
	fmt.Printf("file info: %#v\n", info)
}

func main2() {
	print("hello world\n")
	for _, arg := range os.Args {
		print(arg)
		print(" ")
	}
	print("\n")
}
