package main

// demo gogo
//      loop
//
import (
	//	"github.com/gin-gonic/gin"
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
	//	"github.com/widuu/goini"
	//      _ "github.com/go-sql-driver/mysql"
	//      "database/sql/driver"
)

func main() {
	fmt.Println("--------------")
	goto LOOP
	fmt.Println("bbbbbbbbbb")

LOOP:
	fmt.Println("start loop")

	for i := 0; i < 3; i++ {
		switch i % 2 {
		case 1:
			goto OUTER
		case 0:
			fmt.Print("in loop!")

		}
		if i == 2 {
			goto OUTER
		}

	}
OUTER:
}
