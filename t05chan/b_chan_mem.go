package main

/*
边写边读成的chan
用于作 消息缓冲池使用,用以解决kafka连接过多问题
此为demo,
make(chan int,len)时不占用内容
只有在 c<-<data>操作时,才会战胜内容,弹出出,也不再占用内容
使用len判断chan中元素数量(但并不准确,只能作大要的数据),
只能用于判断是否接近满值的情况
demo中显示,取到数据后,len(c)和<-c那一廖的值是不同步 的,
此后,可以已进行了n步操作.

*/
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
