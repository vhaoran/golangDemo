package main

import (

	//"myblog/g/xbind"
	//	"myblog/db"
	//"log"
	"fmt"
	//"strconv"
	//"strings"
	//"reflect"
	//"time"
	//	"net/http"
	//	"bytes"
	//	"encoding/binary"
	"encoding/json"
	//	"github.com/widuu/goini"
	//      _ "github.com/go-sql-driver/mysql"
	//      "database/sql/driver"
)

func main() {
	fmt.Println("--------------")
	s := "{\"id\":1,\"CName\":\"good\"}"
	m := &map[string]interface{}{}
	err := json.Unmarshal([]byte(s), m)

	fmt.Println(err)

	fmt.Println(m)
	fmt.Println(*m)

}
