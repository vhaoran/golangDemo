package main

import (
	"fmt"
)

func main() {
	a := []int32{10, 20}
	b := a

	fmt.Println("a len:", len(a))
	fmt.Println("-----------")
	for _, v := range a {
		fmt.Println("v = ", v)
	}

	fmt.Println("------------")

	a = append(a, 101, 201)
	for _, v := range a {
		fmt.Println("v = ", v)
	}

	fmt.Println("b-------------")

	for _, v := range b {
		fmt.Println("v = ", v)
	}

}
