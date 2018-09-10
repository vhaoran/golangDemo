package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------begin-----")
	limit := 10000 * 1000
	for i := 0; i < limit; i++ {
		lst := []*string{}
		s := fmt.Sprintf("%d--x", i)
		lst = append(lst, &s)
	}
	fmt.Println("before waitting....")
	time.Sleep(time.Second * 20)
}
