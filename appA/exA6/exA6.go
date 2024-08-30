package main

import (
	"embed"
	"net/http"
)

// static 폴더 아래 있는 모든 파일을 실행 파일 내에 포함시킵니다.
//
//go:embed static/*
var files embed.FS // 파일 추가

func main() {
	http.Handle("/", http.FileServer(http.FS(files))) // 파일 서버 실행
	http.ListenAndServe(":3000", nil)
}
