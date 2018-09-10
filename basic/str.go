package main

import (
	"fmt"
)

func main() {
	var a string

	s1 := "china"
	fmt.Println(s1, "---s1 len:", len(s1))
	s2 := "查询到的密码"
	fmt.Println(s2, "---  s2 len:", len(s2))
	s3 := t(s2)
	fmt.Println(s3, "---  s2 len:", len(s3))

}
