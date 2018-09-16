package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	h := 10000 * 80
	for i := 0; i < h; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			if index%10000 == 0 {
				fmt.Println("---", index)
			}
			time.Sleep(time.Second * 50)
			//			fmt.Println("###  ", index)
		}(i)

	}
	wg.Wait()
	fmt.Println("ok")
}
