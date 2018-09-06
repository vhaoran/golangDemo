package main

import "fmt"
import "reflect"

func main() {
	test2()

}

func test3() {
	s := "this is string"
	y := reflect.TypeOf(s)

	fmt.Println(y)
	fmt.Println("-------------------")

}

func test2() {
	user := &User{22, "whr", 33}

	listv := reflect.ValueOf(user).Elem()
	fmt.Println("------------")

	listT := reflect.TypeOf(user).Elem()

	for i := 0; i < listv.NumField(); i++ {
		f := listv.Field(i)
		if f.CanSet() {
			if f.Kind() == reflect.Int32 {
				f.SetInt(f.Int() + 1)
			}
		}

		fmt.Printf("%d: %s %s %v \n", i, listT.Field(i).Name, f.Type(), f.Interface())
	}

}

type User struct {
	ID   int32
	Name string
	Age  int32
}

//------------------------------------------------------
func test1() {
	hl := hello
	fv := reflect.ValueOf(hl)

	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)

	fmt.Println("Addr: ?", fv.Addr)
	fmt.Println("Bool: ?", fv.Bool)

	fv.Call(nil)

}

func hello() {
	fmt.Println("hello world")
}

//-----------------------------------------------
