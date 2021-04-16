package performances

import "testing"

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		Fib3(20)
	}
}

func Benchmark28(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib3(28)
	}
}

func BenchmarkFibWrong(b *testing.B) {
	Fib3(b.N)
}
