package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	a := make(chan int)

	a <- 3

	fmt.Println("prog blocked,this line is not show!")
	<-a
}
