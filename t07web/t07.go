package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/test/", onTest)
	debug("server is running...")

	http.ListenAndServe(":80", nil)
}

func Hello(response http.ResponseWriter, request *http.Request) {
//	debug("enter handler -->hello")


	fmt.Fprintf(response, "Hello, Welcome to go web programming...")
}
func onTest(response http.ResponseWriter, request *http.Request) {
	// debug("enter handler -->onTest")
	fmt.Fprintf(response, "test is called!\n\r")
}

func debug(lst ...interface{}) {

	fmt.Println(lst)
}
