package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("defer ----1")
	a := 1
	defer fmt.Println("defer ----2")
	a += 1
	defer fmt.Println("defer ----3")
	fmt.Println("a = ", a)
}
