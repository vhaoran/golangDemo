package main

import (
	"fmt"
	"reflect"
	//	"strings"
)

func main() {
	//---------int type-------------------
	var a int = 10
	t := reflect.TypeOf(a)
	fmt.Println("i typeof:", t)

	//---interface type---------------
	var b interface{}
	b = 5
	fmt.Println("b interface{} type:", reflect.TypeOf(b))

	v := reflect.ValueOf(b)
	fmt.Println("b interface{} type:", reflect.TypeOf(v))
	fmt.Println(v.Int())

	x := []byte{49, 53}

	y := []byte(x)
	fmt.Println("b interface{} type:", reflect.TypeOf(x))
	//	x_v := reflect.ValueOf(x)
	fmt.Println("x:", string(x))
	fmt.Println("", string(y))

	//------------------------
	b_byte := b.([]byte)
	fmt.Println("", string(b_byte))

	//-------------------------------
	i3_byte := []byte{20, 30}
	i3 := int32(i3_byte)

}
