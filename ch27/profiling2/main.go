package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

// 피보나치 결과를 저장할 맵
var fibMap [65535]int

func Fib(n int) int {
	// 이미 계산했다면 바로 반환합니다.
	f := fibMap[n]
	if f > 0 {
		return f
	}

	if n == 0 {
		return 0
	} else if n == 1 {
		f = 1
	} else {
		f = Fib(n-1) + Fib(n-2)
	}
	fibMap[n] = f
	return f
}

func main() {
	//프로파일링 결과를 저장할 파일을 만듭니다.
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 프로파일링을 시작합니다.
	pprof.StartCPUProfile(f)
	// 프로그램 종료 전에 프로파일링을 종료합니다.
	defer pprof.StopCPUProfile()
	fmt.Println(Fib(50))

	// 10초를 대기합니다.
	time.Sleep(10 * time.Second)
}
