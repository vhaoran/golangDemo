package main

import (
	//	"bufio"
	"fmt"
	"math"
	//	"database/sql"
)

func main() {
	var a = math.Abs(1)
	fmt.Println(a)

	fmt.Println("----------")
	lst := [3]int{1, 2, 3}
	for _, each := range lst {
		fmt.Println(each)
	}

	fmt.Println("----------")
	lst[0], lst[1] = lst[1], lst[0]
	for _, each := range lst {
		fmt.Println(each)
	}

}
