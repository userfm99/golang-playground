package fibo

import "testing"

func BenchmarkFibo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibo(20)
	}
}

func simpleTest(t *testing.T) {
	t.Errorf("nothing error")
}
