package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 5 데이터 빼옴

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // 1 채널 생성

	wg.Add(1)
	go square(&wg, ch) // 2 고루틴 생성
	ch <- 9            // 3 채널에 데이터 넣음
	wg.Wait()          // 4 작업이 완료되길 기다림
}
