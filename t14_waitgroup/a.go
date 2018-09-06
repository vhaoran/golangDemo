package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			inc(n)
			fmt.Println("-----------i=", n)
			time.Sleep(time.Second * 1)
			wg.Add(-1)
		}(i)
	}
	fmt.Println("---before wg.wait----")
	wg.Wait()
	fmt.Println("---after wg.wait----")
}

func inc(i int) {
	time.Sleep(time.Second * 1)
	fmt.Println(i)
}
