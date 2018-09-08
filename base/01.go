package main

import (
	"fmt"
)

/* demo variables declare*/

func main() {
	var (
		a1 int32
		a2 string
	)
	a1, a2 = 3, "55"
	fmt.Println(a1, a2)

	fmt.Println("hello worldaaaaaaa")
	fmt.Println("step 1")

	//declare 1,classcial
	var a int32 = 4
	fmt.Println(a)
	//common usages
	b := 5
	fmt.Println(b)
	//declare mult
	var c, d int32 = 5, 4
	fmt.Println(d, c)
	//empty declare
	var e int32
	e = 5
	fmt.Println(e)

}
