package main

import (
	//	"github.com/gin-gonic/gin"
	//"myblog/g/xbind"
	//	"myblog/db"
	//"log"
	"fmt"
	//"strconv"
	//"strings"
	//"reflect"
	"time"
	//	"net/http"
	//	"bytes"
	//	"encoding/binary"
	//	"github.com/widuu/goini"
	//      _ "github.com/go-sql-driver/mysql"
	//      "database/sql/driver"
)

func main() {
	fmt.Println("--------------")
	limit := 100000
	c := make(chan int64, limit)
	fmt.Println("before waiting...")

	go func() {
		for {
			select {
			case z := <-c:
				fmt.Println("---", z, "--len", len(c))
			}
		}
	}()
	go func() {
		for z := range c {
			fmt.Println("##########", z, "--len", len(c))
		}
	}()

	for i := 0; i < limit; i++ {
		var k int64 = int64(i)
		c <- k
		fmt.Println("len:", len(c))
	}

	time.Sleep(time.Second * 10)
}
