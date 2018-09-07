package main

/*
此例 demo 原理:
只要主线程不退出,则从其它函数中调用的go routine是不会退出的
*/

import (
	//	"github.com/gin-gonic/gin"
	//"myblog/g/xbind"
	//	"myblog/db"
	//"log"
	"fmt"
	"time"
	//"strconv"
	//"strings"
	//"reflect"
	//"time"
	//	"net/http"
	//	"bytes"
	//	"encoding/binary"
	//	"github.com/widuu/goini"
	//      _ "github.com/go-sql-driver/mysql"
	//      "database/sql/driver"
)

func main() {
	fmt.Println("----------begin----")
	a()
	fmt.Println("before sleep###################")
	time.Sleep(time.Second * 200)

}

func a() {
	for i := 0; i < 5; i++ {
		go f(i)
	}

}
func f(thread_id int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 10)
		s := fmt.Sprintf("--id:  %d ----index: %d", thread_id, i)
		fmt.Println(s)
	}
}
