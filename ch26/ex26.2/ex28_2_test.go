package main

import "testing"

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N 만큼 반복
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
