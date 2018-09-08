package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------------aa----------")
	w := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i :", i)
			time.Sleep(time.Second * 1)
		}
		w <- 1
	}()

	a, x := <-w

	if x {
		fmt.Println("------end----- x =", x, "a=", a)
	}
	//再次访问时将产生一个错误,

	select {
	case z, ok := <-w:
		fmt.Println("read ok,result:z=", z, "ok", ok)
	default:
		fmt.Println("read fail")
	}

	fmt.Println("   end ")
}
