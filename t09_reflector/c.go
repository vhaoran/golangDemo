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

	aa(a)

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

	result = append(result, *u)
	x := (unsafe.Pointer)(reflect.ValueOf(each).Pointer())
	result = append(result, x)

	v := reflect.ValueOf(each)

	log.Println(v)
	fmt.Println(unsafe.Pointer(u))

	return result
	return nil
}

func aa(tp interface{}) {
	t := reflect.ValueOf(tp).Type()

	reflect.MakeSlice(typ reflect.Type, len int, cap int)

	for i := 0; i < 10; i++ {
		data := reflect.New(t)

		a := data.Interface()
		u := a.(*User)
		u.ID = i * 2
		u.Name = "whr"

		b := data.Convert(t)

		log.Println("b:", b)
	}

	log.Println("=----------end loop--------")
	log.Println("len:", t.Size())
	log.Println("-------------")

}
