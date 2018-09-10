package t02

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (user *User) ToString() string {

	fmt.Println("enter toString func")
	if user == nil {
		return ""
	}
	
	return user.Name
}

