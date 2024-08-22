package main

import "fmt"

type onFunc func(int, int) int

func getOperator(op string) onFunc {
	if op == "+" {
		// 함수 리터럴을 사용해서 더하기 함수를 정의하고 반환
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		// 함수 리터럴을 사용해서 곱하기 함수를 정의하고 반환
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입 변수
	fn := getOperator("*")

	var result = fn(3, 4) // 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}
