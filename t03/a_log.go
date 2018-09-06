package main

import (
	"fmt"
	"log"
)

func main() {

	log.SetFlags(1)
	fmt.Println("fmt.Println")
	log.Println("log.println", "def")

	log.Panicln("PanicLn")

}
