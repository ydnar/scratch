package main

import "os"

// This requires TinyGo with WASI Preview 2 support
// https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

func main() {
	print("hello world\n")
	for _, arg := range os.Args {
		print(arg)
		print(" ")
	}
	print("\n")
}
