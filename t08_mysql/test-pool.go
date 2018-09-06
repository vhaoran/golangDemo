package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	for i  := 0; i < 900; i++{
	go test1(i)

	}
	time.Sleep(time.Second*100)
	
	
}

 
func test1(id int) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from t limit 10 ")
	CheckErr(err)

	fmt.Println("a", "b")

	for rows.Next() {
		var a int
		var b string

		err = rows.Scan(&a, &b)
		time.Sleep(time.Second)
		

		CheckErr(err)
		fmt.Println("-------",id,"---------",a,",", b)
	}

	db.Close()

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
