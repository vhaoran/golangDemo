package main

import (
	//	"fmt"
	//"strconv"
	"reflect"
	//"time"
	//"strings"
	"log"
)

type User struct {
	ID   int
	Name string
}

func main() {
	log.Println("--------------")

	t := &User{}
	//log.Println("------------------")
	log.Println("type.kind", reflect.TypeOf(t).Kind())
	log.Println("type.kind", reflect.TypeOf(t).Kind())
	//	log.Println("type.type", reflect.TypeOf(t).Type())

	log.Println("v.Kind", reflect.ValueOf(t).Kind().String())
	log.Println("v.Type", reflect.ValueOf(t).Type())

}
