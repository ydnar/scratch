package main

import (
	"fmt"
	"os"
)

// This requires TinyGo with WASI Preview 2 support
// https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

func main() {
	fmt.Println("Hello world from WebAssembly!")
	fmt.Print("os.Args: ")
	for _, arg := range os.Args {
		fmt.Print(arg, " ")
	}
	fmt.Print("\n\n")

	filename := os.Args[1]
	fmt.Printf("File.Stat() of %s\n", filename)
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
