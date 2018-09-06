package main

import (
	"fmt"
	"strings"
)

func main() {
	callSafe()

	s := strings.Repeat("-", 50)
	fmt.Println("", s)

	fmt.Println("after safe call", "")

}

func callSafe() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)

		}
		fmt.Println("enter defer")
	}()
	func() {
		x, y := 0, 1
		z := y / x
		fmt.Println("", z)

	}()
}
