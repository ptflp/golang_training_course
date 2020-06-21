package main

import "fmt"

func showBinaryRepresentation(n int) {
	fmt.Printf("%d - %b \n", n, n)
}

func main() {
	showBinaryRepresentation(55)
}
