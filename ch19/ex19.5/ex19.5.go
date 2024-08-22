package main

import "fmt"

/*
*
함수 리터럴은 필요한 변수를 내부 상태로 가질 수 있습니다.
함수 리터럴 내부에서 사용되는 외부 변수는 자동으로 함수 내부 상태로 저장됩니다.
*/
func main() {
	i := 0

	f := func() {
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}
