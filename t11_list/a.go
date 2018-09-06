package main

import (
	//"fmt"
	//"strconv"
	//"reflect"
	//"time"
	//"strings"
	"container/list"
	"log"
)

func main() {
	log.Println("--------------")
	lst := list.List{}
	lst.Init()

	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushFront(-1)
	log.Println(lst.Len())

	log.Println(lst)

	log.Println("----------")

	for v := range lst {
		log.Println(v)
	}

}
