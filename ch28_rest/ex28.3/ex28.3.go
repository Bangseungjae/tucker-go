package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // ServeMux인스턴스 생성
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!") // 인스턴스에 핸들러 등록
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListenAndServe(":3000", mux)
}
