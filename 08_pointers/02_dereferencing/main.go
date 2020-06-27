package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43
}
