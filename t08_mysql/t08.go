package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"unsafe"
)

func main() {
	//	test1()

	fmt.Println("-------2-----")
	test2()
	fmt.Println("-------3-----")
	//	test3()
}

func test3() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from t limit 100 ")
	CheckErr(err)

	//other ver's getRow data
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)

		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}

		fmt.Println(record)
	}
}

func test2() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/myblog?charset=utf8")
	CheckErr(err)

	id := int64(1)
	rows, err := db.Query("select * from testUser where id > ?  limit 10 ", &id)
	CheckErr(err)

	//other ver's getRow data
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		err = rows.Scan(scanArgs...)

		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

		fmt.Println(record)
	}

}

func test1() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from t limit 10 ")
	CheckErr(err)

	fmt.Println("a", "b")

	for rows.Next() {
		var a int
		var b string

		err = rows.Scan(&a, &b)

		CheckErr(err)
		fmt.Println(a, b)
	}

	db.Close()

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
