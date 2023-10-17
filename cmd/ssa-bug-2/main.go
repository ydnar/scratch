// https://github.com/golang/go/issues/48105#issuecomment-944611043

package main

var x = 0

func f() int {
	x = 3
	return x
}

func main() {
	x = 0
	a, _ := x, f()

	x = 0
	var b, _ = x, f()
	println(a, b)
}
