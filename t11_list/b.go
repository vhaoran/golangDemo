package main

import (
	//	"fmt"
	//"strconv"
	//"reflect"
	//"time"
	//"strings"
	"log"
)

func main() {
	p := []int{}

	p = append(p, 1)
	show(p)
	//log.Println("------------------")
	log.Println("")

	p = append(p, 2)
	show(p)

}

// name ...
func show(a interface{}) {
	//log.Println("------------------")
	log.Println("", a)

}
