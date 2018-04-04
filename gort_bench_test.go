package gort

import (
	"sort"
	"testing"
)

func BenchmarkGortWithArray(b *testing.B) {
	b.ResetTimer()

	data := make([]int, 10)

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}

func BenchmarkForWithSlice(b *testing.B) {
	b.ResetTimer()

	data := make([]int, 10)

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}

func BenchmarkDefaultSortWithSlice(b *testing.B) {
	b.ResetTimer()

	data := make([]int, 10)

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}
