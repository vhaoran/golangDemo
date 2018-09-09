package main

import (
	"fmt"
)

/*
 demo variables declare
*/

func main() {
	//type 1
	var (
		a1 int32
		a2 string
	)
	a1, a2 = 3, "55"
	fmt.Println(a1, a2)

	fmt.Println("hello worldaaaaaaa")
	fmt.Println("step 1")

	//type 2
	var a int32 = 4
	fmt.Println(a)
	//type 3
	b := 5
	fmt.Println(b)
	//type 4
	var c, d int32 = 5, 4
	fmt.Println(d, c)
	//type 5
	var e int32
	e = 5
	fmt.Println(e)
	//type 6
	var c5 = 5
	fmt.Println(c5)
}
