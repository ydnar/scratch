package main

import (
	"fmt"
	"io"
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

	// wd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "os.Getwd(): %v", err)
	// 	return
	// }
	// fmt.Printf("os.Getwd: %s\n", wd)

	fmt.Println("os.Environ: ")
	for _, e := range os.Environ() {
		fmt.Print(e, "\n")
	}
	fmt.Print("\n")

	fmt.Print("os.Args: ")
	for _, arg := range os.Args {
		fmt.Print(arg, " ")
	}
	fmt.Print("\n")

	filename := os.Args[1]
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Stat(): %v", err)
		return
	}
	fmt.Printf("os.Stat size: %d\n", info.Size())
	fmt.Print("\n")

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open(%q): %v", filename, err)
		return
	}
	defer f.Close()
	info, err = f.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "(*os.File).Stat(): %v", err)
		return
	}
	fmt.Printf("(*os.File).Stat size: %d\n", info.Size())
	fmt.Print("\n")

	var buf [256]byte
	for {
		b := buf[:]
		n, err := os.Stdin.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		b = b[:n]
		fmt.Print(string(b))
	}
	fmt.Print("\n\n")
}
