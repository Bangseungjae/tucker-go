package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var c int = 20
	var d int64 = 20
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
}
