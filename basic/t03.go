package main

import (
	"fmt"
	"t02"
)

type UserX struct {
	t02.User
	CName string
}

func main() {

	user := UserX{}
	user.CName = "good"
	user.Name = "whr"
	user.Age = 30
	user.ID = 1

	fmt.Println(user.ToString())
	fmt.Println("break line", user.CName)

	fmt.Printss("aa")
	fmt.Println("bbc")
	
		
	



}
