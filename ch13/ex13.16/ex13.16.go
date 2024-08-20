package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Printf("str:\t%p\n", unsafe.StringData(str))
	fmt.Printf("slice:\t%p\n", unsafe.SliceData(slice))
}
