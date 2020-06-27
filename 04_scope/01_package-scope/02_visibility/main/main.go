package main

import (
	"fmt"
	"github.com/ptflp/golang_training_course/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
