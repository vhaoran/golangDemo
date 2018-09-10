package main

// deom for

import (
	"fmt"
)

func main() {
	fmt.Println("--------------")
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
	//
	fmt.Println("----------------")
	lst := [5]int{3, 4, 5, 6, 7}
	for i, each := range lst {
		fmt.Println(i, each)
	}
	//
	fmt.Println("----")
	for _, each := range lst {
		fmt.Println(each)
	}

	//
	i := 0
	for {
		fmt.Println("##")
		if i = i + 1; i > 5 {
			break
		}
	}
}
