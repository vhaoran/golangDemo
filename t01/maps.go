package main

import "fmt"

func main() {
	x := map[int]string{}
	x[0] = "aaaa"
	x[4] = "def"

	fmt.Println("map len:", len(x))
	fmt.Println("map value:", x)

}
