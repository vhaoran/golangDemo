package main

/*
演示一个读的chan和一个写的chan
并处理了超时
chan的读一般不要使用<-之类的操作
只考虑for each := range c
与select
*/
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("--------------")
	c := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go chan_in(&wg, c)
	wg.Add(1)
	go chan_out(&wg, c)
	wg.Wait()
}

func chan_in(wg *sync.WaitGroup, c chan<- int) {
	defer wg.Add(-1)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("-----in successed,i = ", i)
	}
}
func chan_out(wg *sync.WaitGroup, c <-chan int) {
	defer wg.Add(-1)
	for {
		timeout := make(chan bool)
		go func() {
			time.Sleep(time.Second * 5)
			timeout <- true
		}()

		select {
		case i := <-c:
			fmt.Println("######---i = ", i)
		case <-timeout:
			goto OUTER
		}
	}
OUTER:
}
