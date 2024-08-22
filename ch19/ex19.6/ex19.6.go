package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i) // 캡쳐된 i갑 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v) // 캡쳐된 i갑 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

/*
*
1.22 버전 이전의 겨로가
ValueLoop
3
3
3
ValueLoop2
0
1
2

1.22 버전 이후의 결과
ValueLoop
0
1
2
ValueLoop2
0
1
2
*/
func main() {
	CaptureLoop()
	CaptureLoop2()
}
