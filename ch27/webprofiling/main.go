package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 웹 프로파일링을 실행합니다.
	"time"
)

func main() {
	http.HandleFunc("/log", logHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		// simulation of a time consuming process like wrong logs into db
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()

	select {
	case status := <-ch:
		w.WriteHeader(status)
	case <-time.After(200 * time.Millisecond):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
