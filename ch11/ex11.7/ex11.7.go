package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1 byte
	B int8 // 1 byte
	C int8 // 1 byte
	D int  // 8 byte
	E int  // 8 byte
}

func main() {
	user := User{1, 2, 3, 4, 5}

	fmt.Printf("%d\n", unsafe.Sizeof(user))
}
