package main

import "testing"

var initialNumber uint = 100

func Benchmark(b *testing.B) {
	for initial := 0; initial < b.N; initial++ {
		Fibonacci(initialNumber)
	}
}
