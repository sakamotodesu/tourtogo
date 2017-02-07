package main

import "testing"

func BenchmarkNoMake(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

func BenchmarkUseMake(b *testing.B) {
	s := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}

}
