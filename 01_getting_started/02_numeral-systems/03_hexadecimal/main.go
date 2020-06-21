package main

import "fmt"

func showHexRepresent(n int) {
	fmt.Printf("%d \t %#X \n", n, n)
}

func main() {
	showHexRepresent(13)
}
