package gort

import (
	"sort"
	"testing"
)

func BenchmarkGortIntArray(b *testing.B) {
	data := make([]int, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}

func BenchmarkGortIntSlice(b *testing.B) {
	data := make([]int, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}

func BenchmarkDefaultSortIntSlice(b *testing.B) {
	data := make([]int, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}
