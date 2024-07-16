package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("Main goroutine: %v\n", GoroutineID())
	wg.Add(1)
	go func() {
		fmt.Printf("Other goroutine: %v\n", GoroutineID())
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("Main goroutine: %v\n", GoroutineID())
}
