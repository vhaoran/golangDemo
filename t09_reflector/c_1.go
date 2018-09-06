package main

import (
	"fmt"
	//"strconv"

	"reflect"

	//"time"
	//"strings"
	"log"
	"unsafe"
)

type User struct {
	ID   int
	Name string
}

func main() {

	a := User{}

	r := test_slice_type(a)

	log.Println(a)
	log.Println(r)

}

func test_slice_type(elementType interface{}) (result []interface{}) {
	log.Println("---------------------------------")

	//	v_a := reflect.ValueOf(elementType)
	t := reflect.ValueOf(elementType).Type()
	log.Println(t)

	log.Println("type:", t)
	log.Println("kind:", t.Kind())
	//	log.Println("Elem:", t.Elem())
	//	log.Println("Elem.Kind:", t.Elem().Kind())

	//---------new element ---------------
	//each := reflect.New(t).Interface()
	each := reflect.New(t).Interface()

	u := each.(*User)
	u.ID = 11
	u.Name = "whr test"

	log.Println("each", each)
	log.Println("u:", u)

	//	fmt.Println(*(unsafe.Pointer(each)))
	//	reflect.NewAt(typ reflect.Type, p unsafe.Pointer)

	//	fmt.Println(x)

	//	fmt.Println()

	result = append(result, *u)

	x := (unsafe.Pointer)(reflect.ValueOf(each).UnsafeAddr())

	result = append(result, x)

	v := reflect.ValueOf(each)

	log.Println(v)
	fmt.Println(unsafe.Pointer(u))

	return result
	//----------add to slice-----
	//temp := reflect.Append(v_a, reflect.ValueOf(User{11, "whr"}))

	//	v_a.SetCap(10)
	//v_a.Index(0).Set(reflect.ValueOf(User{1, "whr"}))

	//	log.Println(v_a)
	//	log.Println(temp)

	//	reflect.Copy(temp, v_a)

	//	log.Println("after copy")

	//	log.Println(v_a)
	//	log.Println(temp)

	//	log.Println("---end func-----")
	return nil
}
