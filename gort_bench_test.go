package gort

import (
	"sort"
	"testing"
)

const Len = 100

func BenchmarkGortIntArray(b *testing.B) {
	b.StopTimer()

	unsorted := make([]int, Len)
	for i := 0; i < Len; i++ {
		unsorted[i] = i ^ 0xcccc
	}
	data := make([]int, 10)

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)

		b.StartTimer()

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})

		b.StopTimer()
	}
}

func BenchmarkGortIntSlice(b *testing.B) {
	b.StopTimer()

	unsorted := make([]int, Len)
	for i := 0; i < Len; i++ {
		unsorted[i] = i ^ 0xcccc
	}

	data := make([]int, Len)

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)

		b.StartTimer()

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})

		b.StopTimer()
	}
}

func BenchmarkDefaultSortIntSlice(b *testing.B) {
	b.StopTimer()

	unsorted := make([]int, Len)
	for i := 0; i < Len; i++ {
		unsorted[i] = i ^ 0xcccc
	}

	data := make([]int, Len)

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)

		b.StartTimer()

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})

		b.StopTimer()
	}
}
