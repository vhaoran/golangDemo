package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	k := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(z *int, index int) {
			defer wg.Done()

			inc(z, index)
		}(&k, i)

	}
	wg.Wait()
}

func inc(i *int, index int) {
	defer m.Unlock()
	m.Lock()
	time.Sleep(time.Second * 1)
	fmt.Println("-------------index is:", index)
	fmt.Println("value before:", *i)
	*i += 3
	fmt.Println("value after:", *i)
}
