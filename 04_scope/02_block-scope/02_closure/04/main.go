package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func wrapperFib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a
		b = a + b
		return a
	}
}
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fib := wrapperFib()
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
}
