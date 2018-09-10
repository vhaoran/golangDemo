package main

/*
wroing application

*/
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------")
	//mult chan not read,and memory raise
	limit := 10000 * 100
	for i := 0; i < limit; i++ {
		c := make(chan bool)
		go func(cc chan<- bool) {
			cc <- true
		}(c)
	}
	fmt.Println("before sleep")
	time.Sleep(time.Second * 20)
}
