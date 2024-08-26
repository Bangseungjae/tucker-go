package main

import "fmt"

func main() {
	ch := make(chan int) // 크기가 0인 채널

	ch <- 9                    // main() 함수가 여기서 멈춥니다.
	fmt.Println("Never print") // 실행되지 않습니다.
}
