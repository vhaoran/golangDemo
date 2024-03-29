package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func main() {
	b := []byte{0x00, 0x00, 0x03, 0xe8}
	b_buf := bytes.NewBuffer(b)

	fmt.Println(b_buf.Bytes())

	var x int32
	binary.Read(b_buf, binary.BigEndian, &x)
	fmt.Println(x)

	fmt.Println(strings.Repeat("-", 100))

	x = 15
	b_buf = bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, x)
	fmt.Println(b_buf.Bytes())
}
