package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("ss")
	s := "25.4"
	f, ok := strconv.ParseFloat(s, 32)
	fmt.Println("", f)

}
