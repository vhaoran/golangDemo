package main

import (
	"fmt" // A package in the Go standard library.
	//	"io/ioutil" // Implements some I/O utility functions.
	//	m "math"    // Math library with local alias m.
	//	"net/http"  // Yes, a web server!
	//	"os"        // OS functions like working with the file system
	//	"strconv"   // String conversions.
)

func main() {

	demoChan()

}

func inc(i int, c chan int) {
	c <- i + 1
}

func demoChan() {
	c := make(chan int)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)

	fmt.Println("fmt.Println(<-c, <-c, <-c)")
	fmt.Println(<-c, <-c, <-c)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)
	go inc(50, c)
	cs := make(chan string)
	ccs := make(chan string)

	go func() { c <- 84 }()
	go func() { cs <- "wordy" }()
	go func() { ccs <- "css value" }()

	go func() { c <- 88 }()

	fmt.Printf("%T\n\r", <-c)

	fmt.Println("----before for-------")

	for {
		bexit := false

		select {
		case i := <-c:
			{
				fmt.Println("c: ", i)
			}
		case str := <-cs:
			fmt.Println("cs:", str)
		case str2 := <-ccs:
			fmt.Println("css:", str2)
		default:
			{
				fmt.Println("no data")
				bexit = true
			}
		}

		fmt.Println("--------------")
		if bexit {
			break
		}

	}

}
