package main

import (
	//	"github.com/gin-gonic/gin"
	//"myblog/g/xbind"
	//	"myblog/db"
	"log"
	//	"fmt"
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

type User struct {
	ID    string
	CName string
}

func main() {
	log.Println("--------------")
	s := `{"ID":"aaa","CName":"whr"}`
	u := &User{}
	json.Unmarshal([]byte(s), u)

	log.Println(u)

}
