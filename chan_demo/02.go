package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	a := make(chan int)
	fmt.Println("before go func()")

	go func() {
		fmt.Println("go routine---->read from chan")
		var i int
		i = <-a
		fmt.Println("go routine---->i is :", i)
	}()

	fmt.Println("before read")
	fmt.Println("prog blocked,untile routine read data!\n\r")
	a <- 3
	fmt.Println("after read")
}
