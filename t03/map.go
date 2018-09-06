package main

import "fmt"

func main() {
	a := make(map[int32]interface{})
	a[1] = 2

	b := a[3]

	if b == nil {
		fmt.Println("b is nil")
	}

	fmt.Println("end")

}
