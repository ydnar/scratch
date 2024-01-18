package main

import (
	"fmt"
	_ "net/http"
	"os"
	"time"
)

// This requires TinyGo with WASI Preview 2 support
// https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

func main() {
	fmt.Print("Hello world from WebAssembly!\n\n")

	start := time.Now()
	fmt.Printf("time.Now: %v\n", start)
	fmt.Println("sleeping for 10ms...")
	time.Sleep(10 * time.Millisecond)
	end := time.Now()
	fmt.Printf("elapsed: %v\n\n", end.Sub(start))

	fmt.Println("os.Environ: ")
	for _, e := range os.Environ() {
		fmt.Print(e, "\n")
	}
	fmt.Print("\n\n")

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
