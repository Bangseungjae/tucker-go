package main

import "fmt"

func main() {
	str := "Hello 월드!" // 한영이 섞인 문자열
	arr := []rune(str)

	for _, v := range arr { // range를 이용한 순회
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", v, v, v) // 출력
	}
}
