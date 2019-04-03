package main

import "testing"

//Function starts from Benchmark. WE use B structure b
//b.N - how many time we will run the benchmark

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fib(40)
		}
	})
}
