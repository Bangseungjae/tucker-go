package main

import "fmt"

func main() {
	str := "Hello 월드"  // 한영이 섞인 문자열
	arr := []rune(str) // 문자열 []rune으로 타입 변환

	for i := 0; i < len(arr); i++ { // 문자열 []rune으로 타입 변환
		// 바이트 단위로 출력
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}
