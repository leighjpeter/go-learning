package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", n)
	}
}

func BenchmarkItoa(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(n)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	n := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(n, 10)
	}
}
