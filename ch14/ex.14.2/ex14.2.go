package main

import (
	"fmt"
	"trucker/hello/ch14/ex.14.2/publicpkg"
)

func main() {
	fmt.Println("PI:", publicpkg.PI) // 공개되는 상수 접근
	publicpkg.PublicFunc()           // 공개되는 함수 호출

	var myInt publicpkg.MyInt = 10 // 공개되는 별칭 타입 사용
	fmt.Println("myInt:", myInt)

	var myStruct = publicpkg.MyStruct{Age: 18} // 구조체 사용
	fmt.Println("myStruct:", myStruct)

	p := publicpkg.NewMyStruct(3, `hop`)
	fmt.Println("p:", *p)
}
